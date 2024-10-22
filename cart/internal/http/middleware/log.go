package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

type LogMux struct {
	h http.Handler
}

func NewLogMux(h http.Handler) http.Handler {
	return &LogMux{h: h}
}

type statusRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (sr *statusRecorder) WriteHeader(code int) {
	sr.statusCode = code
	sr.ResponseWriter.WriteHeader(code)
}

func (m *LogMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	recorder := &statusRecorder{ResponseWriter: w, statusCode: http.StatusOK}

	t := time.Now()

	m.h.ServeHTTP(recorder, r)

	logFunc := slog.Info
	if recorder.statusCode != http.StatusOK {
		logFunc = slog.Error
	}
	logFunc(
		"got request",
		"url:", r.RequestURI,
		"duration_ms:", time.Since(t).Milliseconds(),
		"status_code", recorder.statusCode,
	)
}
