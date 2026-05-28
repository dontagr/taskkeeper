package storage

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/dontagr/taskkeeper/internal/task"
)

// ErrNotFound is returned when task id does not exist.
var ErrNotFound = errors.New("not found")

// MemoryStore is an in-memory thread-safe task store (spec 0001).
type MemoryStore struct {
	mu    sync.RWMutex
	tasks map[string]task.Task
	next  int
}

// NewMemoryStore creates an empty store.
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{tasks: make(map[string]task.Task)}
}

// Create validates title, assigns id t_000001..., stores and returns the task.
func (s *MemoryStore) Create(title string) (task.Task, error) {
	normalized, err := task.ValidateTitle(title)
	if err != nil {
		return task.Task{}, err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.next++
	id := fmt.Sprintf("t_%06d", s.next)
	tk := task.Task{
		ID:        id,
		Title:     normalized,
		CreatedAt: time.Now().UTC(),
	}
	s.tasks[id] = tk
	return tk, nil
}

// Get returns a task by id or ErrNotFound.
func (s *MemoryStore) Get(id string) (task.Task, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	tk, ok := s.tasks[id]
	if !ok {
		return task.Task{}, ErrNotFound
	}
	return tk, nil
}
