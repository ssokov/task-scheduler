//nolint:dupl
package vt

import (
	"time"

	"apisrv/pkg/db"
)

type Task struct {
	ID          int64      `json:"id"`
	Title       string     `json:"title" validate:"required,max=255"`
	Description *string    `json:"description" validate:"omitempty,max=255"`
	CreatedAt   time.Time  `json:"createdAt" validate:"required"`
	IsDeadline  *bool      `json:"isDeadline"`
	DeadLine    *time.Time `json:"deadLine"`
	StatusID    int        `json:"statusId" validate:"required,status"`

	Status *Status `json:"status"`
}

func (t *Task) ToDB() *db.Task {
	if t == nil {
		return nil
	}

	task := &db.Task{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		CreatedAt:   t.CreatedAt,
		IsDeadline:  t.IsDeadline,
		DeadLine:    t.DeadLine,
		StatusID:    t.StatusID,
	}

	return task
}

type TaskSearch struct {
	ID          *int64     `json:"id"`
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
	CreatedAt   *time.Time `json:"createdAt"`
	IsDeadline  *bool      `json:"isDeadline"`
	DeadLine    *time.Time `json:"deadLine"`
	StatusID    *int       `json:"statusId"`
	IDs         []int64    `json:"ids"`
}

func (ts *TaskSearch) ToDB() *db.TaskSearch {
	if ts == nil {
		return nil
	}

	return &db.TaskSearch{
		ID:               ts.ID,
		TitleILike:       ts.Title,
		DescriptionILike: ts.Description,
		CreatedAt:        ts.CreatedAt,
		IsDeadline:       ts.IsDeadline,
		DeadLine:         ts.DeadLine,
		StatusID:         ts.StatusID,
		IDs:              ts.IDs,
	}
}

type TaskSummary struct {
	ID          int64      `json:"id"`
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	CreatedAt   time.Time  `json:"createdAt"`
	IsDeadline  *bool      `json:"isDeadline"`
	DeadLine    *time.Time `json:"deadLine"`

	Status *Status `json:"status"`
}
