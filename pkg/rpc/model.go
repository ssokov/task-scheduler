package rpc

import (
	"time"
)

type CreateTaskRequest struct {
	UserID      int64      `json:"userId"`
	Title       string     `json:"Title"`
	Description *string    `json:"description"`
	Deadline    *time.Time `json:"deadline"`
}
type CreateTaskResponse struct {
	TaskID int64 `json:"taskId"`
}

type GetTaskByPeriodRequest struct {
	UserID      int64     `json:"userId"`
	PeriodStart time.Time `json:"periodStart"`
	PeriodEnd   time.Time `json:"periodEnd"`
}

type TaskResponse struct {
	ID          int64      `json:"id"`
	Title       string     `json:"title"`
	Description *string    `json:"description,omitempty"`
	IsDeadline  *bool      `json:"isDeadline,omitempty"`
	Deadline    *time.Time `json:"deadline,omitempty"`
	StatusID    int        `json:"statusId,omitempty"`
	Name        string     `json:"Name,omitempty"`
	DeadLine    *time.Time `json:"deadLine,omitempty"`

	CreatedAt time.Time `json:"createdAt"`
}

type GetTaskByPeriodResponse struct {
	Tasks []TaskResponse `json:"tasks"`
}

type UpdateTitleRequest struct {
	TaskID int64  `json:"taskId"`
	Title  string `json:"title"`
}

type UpdateDescriptionRequest struct {
	TaskID      int64  `json:"taskId"`
	Description string `json:"description"`
}

type UpdateDeadLineRequest struct {
	TaskID   int64     `json:"taskId"`
	DeadLine time.Time `json:"deadline"`
}

type DeleteTaskRequest struct {
	TaskID int64 `json:"taskId"`
}

type UpdateStatusRequest struct {
	TaskID   int64 `json:"taskId"`
	StatusID int   `json:"statusId"`
}

type StatusResponse struct {
	ID    int    `json:"statusId"`
	Name  string `json:"Name"`
	Alias string `json:"alias"`
}

type GetStatusesResponse struct {
	Statuses []StatusResponse `json:"statuses"`
}
