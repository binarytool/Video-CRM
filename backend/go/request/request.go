package request

import (
	"log"
	"net/http"
	"time"
)

type Handlers struct {
	logger *log.Logger
}

func (h *Handlers) ProfileRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		h.logger.Println("Request processed in %s", time.Now().Sub(start))
	}
}

func (h *Handlers) SetupRequest(mux *http.ServeMux, path string, fn func(w http.ResponseWriter, r *http.Request)) {
	mux.HandleFunc(path, h.ProfileRequest(fn))
}

func NewHandlers(logger *log.Logger) *Handlers {
	return &Handlers{
		logger: logger,
	}
}

// Handler funcs

func (h *Handler s) Main(w http.ResponseWriter, r *http.Request) {
	h.logger.Println("Request processed")

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("here"))
}

func (h *Handlers) Info(w http.ResponseWriter, r *http.Request) {
	h.logger.Println("Request processed")

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Info"))
}
