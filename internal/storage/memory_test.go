package storage

import (
	"errors"
	"sync"
	"testing"
)

func TestMemoryStore_CreateGet(t *testing.T) {
	s := NewMemoryStore()

	t1, err := s.Create("Buy milk")
	if err != nil {
		t.Fatal(err)
	}
	if t1.ID != "t_000001" {
		t.Fatalf("id=%q want t_000001", t1.ID)
	}
	if t1.Title != "Buy milk" {
		t.Fatalf("title=%q", t1.Title)
	}
	if t1.CreatedAt.IsZero() {
		t.Fatal("createdAt is zero")
	}

	t2, err := s.Create("Walk dog")
	if err != nil {
		t.Fatal(err)
	}
	if t2.ID != "t_000002" {
		t.Fatalf("id=%q want t_000002", t2.ID)
	}

	got, err := s.Get("t_000001")
	if err != nil {
		t.Fatal(err)
	}
	if got.Title != "Buy milk" {
		t.Fatalf("got title %q", got.Title)
	}
}

func TestMemoryStore_GetNotFound(t *testing.T) {
	s := NewMemoryStore()
	_, err := s.Get("t_999999")
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("err=%v want ErrNotFound", err)
	}
}

func TestMemoryStore_CreateInvalidTitle(t *testing.T) {
	s := NewMemoryStore()
	_, err := s.Create("   ")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestMemoryStore_ConcurrentCreate(t *testing.T) {
	s := NewMemoryStore()
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, _ = s.Create("x")
		}()
	}
	wg.Wait()
}
