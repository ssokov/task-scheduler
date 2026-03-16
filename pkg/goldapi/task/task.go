package task

import (
	"apisrv/pkg/db"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/vmkteam/embedlog"
)

type TaskManager struct {
	dbc    db.DB
	tlRepo db.CommonRepo
	embedlog.Logger
}

func NewTaskManager(dbc db.DB, logger embedlog.Logger) *TaskManager {
	return &TaskManager{
		dbc:    dbc,
		tlRepo: db.NewCommonRepo(dbc),
		Logger: logger,
	}
}

var (
	ErrTaskNotFound    = errors.New("task not found")
	ErrDeadLineInPast  = errors.New("deadline is in the past")
	ErrInvalidPersiods = errors.New("invalid periods: periodStart must be before periodEnd")
)

func (tm *TaskManager) CreateTask(ctx context.Context, task *ServiceTask) (int64, error) {

	if task.DeadLine != nil && !task.DeadLine.After(time.Now()) {
		return -1, ErrDeadLineInPast
	}

	isDeadLine := false
	if task.DeadLine != nil {
		isDeadLine = true
	}

	dbTask := newDBTask(task, isDeadLine)

	dbTask, err := tm.tlRepo.AddTask(ctx, dbTask, db.WithoutColumns(db.Columns.Task.CreatedAt))
	if err != nil {
		tm.Log().Error("Failed to add task")
		return -1, fmt.Errorf("Failed to add task: %w", err)
	}

	return dbTask.ID, nil
}

// TODO : добавить пагинацию
func (tm *TaskManager) GetTaskByPeriod(ctx context.Context, userId int64, periodStart time.Time, periodEnd time.Time) (ServiceTaskList, error) {
	tm.Log().Info("GetTaskByPeriod ...")

	if periodStart.After(periodEnd) {
		return ServiceTaskList{}, ErrInvalidPersiods
	}

	dbTasks, err := tm.tlRepo.TasksByFilters(ctx, (&db.TaskSearch{
		UserID: &userId,
	}).WithDeadLineInPeriodOrNull(periodStart, periodEnd), db.Pager{Page: 1, PageSize: 100})
	if err != nil {
		return ServiceTaskList{}, fmt.Errorf("Failed to get tasks by period: %w", err)
	}

	return NewServiceTaskList(dbTasks), nil
}

func (tm *TaskManager) UpdateTitle(ctx context.Context, taskId int64, title string) error {

	ok, err := tm.tlRepo.UpdateTask(ctx, &db.Task{
		ID:    taskId,
		Title: title,
	}, db.WithColumns(db.Columns.Task.Title))

	if !ok {
		return ErrTaskNotFound
	} else if err != nil {
		return fmt.Errorf("Failed to update task title: %w", err)
	}
	return nil
}

func (tm *TaskManager) UpdateDescription(ctx context.Context, taskId int64, description string) error {

	ok, err := tm.tlRepo.UpdateTask(ctx, &db.Task{
		ID:          taskId,
		Description: &description,
	}, db.WithColumns(db.Columns.Task.Description))

	if !ok {
		return ErrTaskNotFound
	} else if err != nil {
		return fmt.Errorf("Failed to update task description: %w", err)
	}
	return nil
}

func (tm *TaskManager) UpdateDeadLine(ctx context.Context, taskId int64, deadline time.Time) error {

	if !deadline.After(time.Now()) {
		return ErrDeadLineInPast
	}

	isDeadline := true

	ok, err := tm.tlRepo.UpdateTask(ctx, &db.Task{
		ID:         taskId,
		DeadLine:   &deadline,
		IsDeadline: &isDeadline,
	}, db.WithColumns(db.Columns.Task.DeadLine, db.Columns.Task.IsDeadline))

	if !ok {
		return ErrTaskNotFound
	} else if err != nil {
		return fmt.Errorf("Failed to update task deadline: %w", err)
	}

	return nil
}

func (tm *TaskManager) DeleteTask(ctx context.Context, taskId int64) error {

	ok, err := tm.tlRepo.DeleteTask(ctx, taskId)
	if !ok {
		return ErrTaskNotFound
	} else if err != nil {
		return fmt.Errorf("Failed to delete task: %w", err)
	}

	return nil
}

func (tm *TaskManager) GetStatuses(ctx context.Context) (ServiceStatusList, error) {

	dbStatuses, err := tm.tlRepo.TaskStatusesByFilters(ctx, nil, db.PagerDefault)
	if err != nil {
		return ServiceStatusList{}, fmt.Errorf("Failed to get task statuses: %w", err)
	}

	return NewServiceStatusList(dbStatuses), nil
}

func (tm *TaskManager) UpdateStatus(ctx context.Context, taskId int64, statusId int) error {

	ok, err := tm.tlRepo.UpdateTask(ctx, &db.Task{
		ID:       taskId,
		StatusID: statusId,
	})

	if !ok {
		return ErrTaskNotFound
	} else if err != nil {
		return fmt.Errorf("Failed to update task status: %w", err)
	}

	return nil

}
