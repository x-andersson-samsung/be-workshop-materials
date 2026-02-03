package main

import (
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const (
	surveyDir   = "/home/a.andersson/Projects/go/backend-workshops-materials/surveys"
	responseDir = "/home/a.andersson/Projects/go/backend-workshops-materials/surveys/responses"
)

func handleSurveyResponse(w http.ResponseWriter, r *http.Request) {
	surveyName := r.PathValue("surveyName")
	data, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Warn("Failed to read body", slog.String("surveyName", surveyName), slog.String("error", err.Error()))
		w.WriteHeader(http.StatusBadRequest)
	}
	defer r.Body.Close()

	err = os.WriteFile(filepath.Join(responseDir, fmt.Sprintf("survey_%s_%d.txt", surveyName, time.Now().UnixNano())), data, 0644)
	if err != nil {
		slog.Warn("Failed to write file", slog.String("surveyName", surveyName), slog.String("error", err.Error()))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func getSurvey(w http.ResponseWriter, r *http.Request) {
	surveyName := r.PathValue("surveyName")
	if surveyName != "" {
		f, err := os.Open(filepath.Join(surveyDir, surveyName+".html"))
		if err != nil {
			slog.Warn("Failed to open file", slog.String("surveyName", surveyName), slog.String("error", err.Error()))
			w.WriteHeader(http.StatusNotFound)
			return
		}

		_, err = io.Copy(w, f)
		if err != nil {
			slog.Warn("Failed to copy survey content", slog.String("surveyName", surveyName), slog.String("error", err.Error()))
			w.WriteHeader(http.StatusNotFound)
			return
		}

		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/{surveyName}", getSurvey)
	mux.HandleFunc("/survey/{surveyName}", getSurvey)
	mux.HandleFunc("/api/surveys/{surveyName}", handleSurveyResponse)

	if err := http.ListenAndServe(":10001", mux); err != nil {
		log.Fatal(err)
	}
}
