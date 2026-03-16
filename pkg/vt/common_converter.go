package vt

import (
	"apisrv/pkg/db"
)

func NewTask(in *db.Task) *Task {
	if in == nil {
		return nil
	}

	task := &Task{
		ID:          in.ID,
		Title:       in.Title,
		Description: in.Description,
		CreatedAt:   in.CreatedAt,
		IsDeadline:  in.IsDeadline,
		DeadLine:    in.DeadLine,
		StatusID:    in.StatusID,

		Status: NewStatus(in.StatusID),
	}

	return task
}

func NewTaskSummary(in *db.Task) *TaskSummary {
	if in == nil {
		return nil
	}

	return &TaskSummary{
		ID:          in.ID,
		Title:       in.Title,
		Description: in.Description,
		CreatedAt:   in.CreatedAt,
		IsDeadline:  in.IsDeadline,
		DeadLine:    in.DeadLine,

		Status: NewStatus(in.StatusID),
	}
}
