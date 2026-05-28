package task

import (
	"errors"
	"strings"
	"time"
)

// ErrInvalidTitle is returned when title fails validation (spec: 1..200 after TrimSpace).
var ErrInvalidTitle = errors.New("invalid title")

// Task is a stored task (spec 0001).
type Task struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
}

// ValidateTitle trims and checks length 1..200.
func ValidateTitle(title string) (string, error) {
	t := strings.TrimSpace(title)
	if len(t) < 1 || len(t) > 200 {
		return "", ErrInvalidTitle
	}
	return t, nil
}
