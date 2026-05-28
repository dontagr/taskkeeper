package httpapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"taskkeeper/internal/storage"
)

func TestPOST_tasks_created(t *testing.T) {
	srv := NewServer(storage.NewMemoryStore())

	req := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(`{"title":"Buy milk"}`))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	srv.Handler().ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Fatalf("status=%d body=%s", rr.Code, rr.Body.String())
	}

	var body map[string]any
	if err := json.NewDecoder(rr.Body).Decode(&body); err != nil {
		t.Fatal(err)
	}
	if body["id"] != "t_000001" {
		t.Fatalf("id=%v", body["id"])
	}
	if body["title"] != "Buy milk" {
		t.Fatalf("title=%v", body["title"])
	}
	if body["createdAt"] == nil || body["createdAt"] == "" {
		t.Fatal("missing createdAt")
	}
}

func TestPOST_tasks_emptyTitle(t *testing.T) {
	srv := NewServer(storage.NewMemoryStore())

	req := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(`{"title":"   "}`))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	srv.Handler().ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("status=%d", rr.Code)
	}

	var body apiError
	if err := json.NewDecoder(rr.Body).Decode(&body); err != nil {
		t.Fatal(err)
	}
	if body.Error.Code != "invalid_argument" {
		t.Fatalf("code=%q", body.Error.Code)
	}
}

func TestGET_tasks_found(t *testing.T) {
	store := storage.NewMemoryStore()
	srv := NewServer(store)

	createReq := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(`{"title":"Buy milk"}`))
	createReq.Header.Set("Content-Type", "application/json")
	createRR := httptest.NewRecorder()
	srv.Handler().ServeHTTP(createRR, createReq)
	if createRR.Code != http.StatusCreated {
		t.Fatalf("create status=%d", createRR.Code)
	}

	getReq := httptest.NewRequest(http.MethodGet, "/tasks/t_000001", nil)
	getRR := httptest.NewRecorder()
	srv.Handler().ServeHTTP(getRR, getReq)

	if getRR.Code != http.StatusOK {
		t.Fatalf("status=%d body=%s", getRR.Code, getRR.Body.String())
	}

	var body map[string]any
	if err := json.NewDecoder(getRR.Body).Decode(&body); err != nil {
		t.Fatal(err)
	}
	if body["id"] != "t_000001" || body["title"] != "Buy milk" {
		t.Fatalf("body=%v", body)
	}
}

func TestGET_tasks_notFound(t *testing.T) {
	srv := NewServer(storage.NewMemoryStore())

	req := httptest.NewRequest(http.MethodGet, "/tasks/t_999999", nil)
	rr := httptest.NewRecorder()
	srv.Handler().ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Fatalf("status=%d", rr.Code)
	}

	var body apiError
	if err := json.NewDecoder(rr.Body).Decode(&body); err != nil {
		t.Fatal(err)
	}
	if body.Error.Code != "not_found" {
		t.Fatalf("code=%q", body.Error.Code)
	}
}
