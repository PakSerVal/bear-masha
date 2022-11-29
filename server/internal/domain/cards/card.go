package cards

import (
	"time"

	"github.com/PakSerVal/bear-masha/internal/domain/file"
)

const (
	StatusTodo = "todo"
	StatusDone = "done"
)

type Card struct {
	ID          int64       `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Status      string      `json:"status"`
	CreatedAt   time.Time   `db:"created_at" json:"created_at"`
	DeadlineAt  time.Time   `db:"deadline_at" json:"deadline_at"`
	Files       []file.File `json:"files"`
}
