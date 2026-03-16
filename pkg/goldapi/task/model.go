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

type ServiceStatus struct {
	ID    int
	Title string
	Alias string
}

type Statuses struct {
	Statuses []ServiceStatus
}
type ServicesTasks struct {
	Tasks []ServiceTask
}
