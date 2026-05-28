package httpapi

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/dontagr/taskkeeper/internal/storage"
	"github.com/dontagr/taskkeeper/internal/task"
)

// Store is the persistence layer used by HTTP handlers.
type Store interface {
	Create(title string) (task.Task, error)
	Get(id string) (task.Task, error)
}

// Server exposes Task keeper HTTP API (spec 0001).
type Server struct {
	store Store
	mux   *http.ServeMux
}

// NewServer wires routes for the given store.
func NewServer(store Store) *Server {
	s := &Server{store: store, mux: http.NewServeMux()}
	s.mux.HandleFunc("POST /tasks", s.handleCreate)
	s.mux.HandleFunc("GET /tasks/{id}", s.handleGet)
	return s
}

// Handler returns the root HTTP handler.
func (s *Server) Handler() http.Handler {
	return s.mux
}

type createRequest struct {
	Title string `json:"title"`
}

type apiError struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func writeAPIError(w http.ResponseWriter, status int, code, message string) {
	var body apiError
	body.Error.Code = code
	body.Error.Message = message
	writeJSON(w, status, body)
}

func (s *Server) handleCreate(w http.ResponseWriter, r *http.Request) {
	var req createRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeAPIError(w, http.StatusBadRequest, "invalid_argument", "invalid JSON body")
		return
	}

	tk, err := s.store.Create(req.Title)
	if err != nil {
		if errors.Is(err, task.ErrInvalidTitle) {
			writeAPIError(w, http.StatusBadRequest, "invalid_argument", "title is required")
			return
		}
		writeAPIError(w, http.StatusInternalServerError, "internal", "unexpected error")
		return
	}

	writeJSON(w, http.StatusCreated, tk)
}

func (s *Server) handleGet(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	tk, err := s.store.Get(id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			writeAPIError(w, http.StatusNotFound, "not_found", "task not found")
			return
		}
		writeAPIError(w, http.StatusInternalServerError, "internal", "unexpected error")
		return
	}

	writeJSON(w, http.StatusOK, tk)
}
