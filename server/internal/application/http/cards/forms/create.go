package forms

import "time"

type CreateCard struct {
	Title       string    `schema:"title,required"`
	Description string    `schema:"description,required"`
	DeadlineAt  time.Time `schema:"deadlineAt,required"`
	Files       *[]int64  `schema:"files"`
}
