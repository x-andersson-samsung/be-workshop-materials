package main

import (
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func NewHandler(tb testing.TB) *Handler {
	tb.Helper()

	templates, err := template.ParseFS(templateFS, "static/*.tmpl")
	if err != nil {
		tb.Fatalf("could not parse templates: %s", err.Error())
	}

	return &Handler{templates}
}

func BenchmarkRenderHandler(b *testing.B) {
	handler := NewHandler(b)
	for b.Loop() {
		rec := httptest.NewRecorder()

		handler.renderHandler(rec, nil)

		resp := rec.Result()

		require.Equal(b, resp.StatusCode, http.StatusOK)

		data, _ := io.ReadAll(resp.Body)
		require.NotEmpty(b, data)
	}
}
