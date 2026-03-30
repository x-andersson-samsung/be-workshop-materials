package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log/slog"
	"net/http"
	"time"
)

var defaultTemplateData = TemplateData{
	BackgroundImage: "https://placehold.co/1920x1080",
	Text:            "Hello World!",
}

//go:embed static
var templateFS embed.FS

type TemplateData struct {
	BackgroundImage string `json:"background_image"`
	Text            string `json:"text"`
}

func (d TemplateData) FromRequest(r *http.Request) (TemplateData, error) {
	r.ParseForm()

	data := TemplateData{
		BackgroundImage: r.Form.Get("background"),
		Text:            r.Form.Get("text"),
	}
	data.AssignDefaults()

	return data, nil
}

func (d *TemplateData) AssignDefaults() {
	if len(d.BackgroundImage) == 0 {
		d.BackgroundImage = defaultTemplateData.BackgroundImage
	}
	if len(d.Text) == 0 {
		d.Text = defaultTemplateData.Text
	}
}

type Handler struct {
	templates *template.Template
}

func MetricMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next(w, r)

		slog.Info("request finished",
			slog.String("path", r.URL.Path), // unsafe passing of parameters to logs. Path and Method should be sanitized first.
			slog.String("method", r.Method),
			slog.Duration("duration", time.Since(start)),
		)
	}
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	type routeData struct {
		path   string
		method string
	}
	switch (routeData{r.URL.Path, r.Method}) {
	case routeData{"/form", http.MethodGet}:
		MetricMiddleware(h.renderHandler)(w, r)
		return
	case routeData{"/json", http.MethodPost}:
		MetricMiddleware(h.renderJSONHandler)(w, r)
		return
	}

	http.NotFound(w, r)
}

func (h Handler) renderHandler(w http.ResponseWriter, r *http.Request) {
	templateData, err := TemplateData{}.FromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.templates.ExecuteTemplate(w, "main.tmpl", templateData)
}

func (h Handler) renderJSONHandler(w http.ResponseWriter, r *http.Request) {
	var body io.ReadCloser
	body = http.NoBody // Unnecessary assignment
	body = r.Body

	data, _ := io.ReadAll(body)

	templateData := TemplateData{}
	if err := json.Unmarshal(data, &templateData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	h.templates.ExecuteTemplate(w, "main.tmpl", defaultTemplateData)
}

func main() {
	templates, err := template.ParseFS(templateFS, "static/*.tmpl")
	if err != nil {
		panic(err)
	}

	handler := Handler{templates}

	port := ":8080"
	fmt.Printf("Listening on %s\n", port)
	if err = http.ListenAndServe(port, handler); err != nil {
		panic(err)
	}

	// gofumpt will complain about empty lines at the end of functions

}
