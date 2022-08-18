package api

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	logger Logger
	mux    *http.ServeMux
}

type Option func(handler *Handler)

type Logger interface {
	Printf(format string, values ...interface{})
}

func LogWith(logger Logger) Option {
	return func(handler *Handler) {
		handler.logger = logger
	}
}

func NewHandler(options ...Option) *Handler {
	handler := &Handler{}
	for _, opt := range options {
		opt(handler)
	}
	handler.mux = http.NewServeMux()
	handler.mux.HandleFunc("/", handler.index)
	handler.mux.HandleFunc("/other", handler.other)
	handler.mux.HandleFunc("/healthz", handler.healthz)
	return handler
}

func (h *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/healthz" && request.URL.Path != "/readyz" {
		h.log("%s %s", request.Method, request.URL.Path)
	}
	h.mux.ServeHTTP(writer, request)
}

func (h *Handler) index(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	_, err := writer.Write([]byte("Deployed with ❤️ Carvel!"))
	if err != nil {
		h.log("Error handling request %s %s", request.Method, request.URL.Path)
		return
	}
}

func (h *Handler) other(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(map[string]interface{}{
		"deployedBy": "Carvel",
	})
	if err != nil {
		h.log("Error handling request %s %s", request.Method, request.URL.Path)
		return
	}
}

func (h *Handler) healthz(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) log(message string, values ...interface{}) {
	if h.logger != nil {
		h.logger.Printf(message+"\n", values...)
	}
}
