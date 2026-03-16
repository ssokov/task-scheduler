package task

import "time"

type ServiceTask struct {
	ID          int64
	UserId      int64
	Title       string
	Description *string
	IsDeadline  *bool
	DeadLine    *time.Time
	StatusID    int
	Name        string
	CreatedAt   time.Time
}

type Status struct {
	ID    int
	Title string
	Alias string
}

type Statuses struct {
	Statuses []Status
}
type ServicesTasks struct {
	Tasks []ServiceTask
}
