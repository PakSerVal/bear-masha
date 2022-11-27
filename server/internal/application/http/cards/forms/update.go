package forms

import "time"

type UpdateCard struct {
	Title       *string    `schema:"title"`
	Status      *string    `schema:"status"`
	Description *string    `schema:"description"`
	DeadlineAt  *time.Time `schema:"deadlineAt"`
}
