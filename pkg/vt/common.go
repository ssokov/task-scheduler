package vt

import (
	"context"

	"apisrv/pkg/db"

	"github.com/vmkteam/embedlog"
	"github.com/vmkteam/zenrpc/v2"
)

type TaskService struct {
	zenrpc.Service
	embedlog.Logger
	commonRepo db.CommonRepo
}

func NewTaskService(dbo db.DB, logger embedlog.Logger) *TaskService {
	return &TaskService{
		Logger:     logger,
		commonRepo: db.NewCommonRepo(dbo),
	}
}

func (s TaskService) dbSort(ops *ViewOps) db.OpFunc {
	v := s.commonRepo.DefaultTaskSort()
	if ops == nil {
		return v
	}

	switch ops.SortColumn {
	case db.Columns.Task.ID, db.Columns.Task.Title, db.Columns.Task.Description, db.Columns.Task.CreatedAt, db.Columns.Task.IsDeadline, db.Columns.Task.DeadLine, db.Columns.Task.StatusID:
		v = db.WithSort(db.NewSortField(ops.SortColumn, ops.SortDesc))
	}

	return v
}

// Count returns count Tasks according to conditions in search params.
//
//zenrpc:search TaskSearch
//zenrpc:return int
//zenrpc:500 Internal Error
func (s TaskService) Count(ctx context.Context, search *TaskSearch) (int, error) {
	count, err := s.commonRepo.CountTasks(ctx, search.ToDB())
	if err != nil {
		return 0, InternalError(err)
	}
	return count, nil
}

// Get returns а list of Tasks according to conditions in search params.
//
//zenrpc:search TaskSearch
//zenrpc:viewOps ViewOps
//zenrpc:return []TaskSummary
//zenrpc:500 Internal Error
func (s TaskService) Get(ctx context.Context, search *TaskSearch, viewOps *ViewOps) ([]TaskSummary, error) {
	list, err := s.commonRepo.TasksByFilters(ctx, search.ToDB(), viewOps.Pager(), s.dbSort(viewOps), s.commonRepo.FullTask())
	if err != nil {
		return nil, InternalError(err)
	}
	tasks := make([]TaskSummary, 0, len(list))
	for i := 0; i < len(list); i++ {
		if task := NewTaskSummary(&list[i]); task != nil {
			tasks = append(tasks, *task)
		}
	}
	return tasks, nil
}

// GetByID returns a Task by its ID.
//
//zenrpc:id int64
//zenrpc:return Task
//zenrpc:500 Internal Error
//zenrpc:404 Not Found
func (s TaskService) GetByID(ctx context.Context, id int64) (*Task, error) {
	db, err := s.byID(ctx, id)
	if err != nil {
		return nil, err
	}
	return NewTask(db), nil
}

func (s TaskService) byID(ctx context.Context, id int64) (*db.Task, error) {
	db, err := s.commonRepo.TaskByID(ctx, id, s.commonRepo.FullTask())
	if err != nil {
		return nil, InternalError(err)
	} else if db == nil {
		return nil, ErrNotFound
	}
	return db, nil
}

// Add adds a Task from the query.
//
//zenrpc:task Task
//zenrpc:return Task
//zenrpc:500 Internal Error
//zenrpc:400 Validation Error
func (s TaskService) Add(ctx context.Context, task Task) (*Task, error) {
	if ve := s.isValid(ctx, task, false); ve.HasErrors() {
		return nil, ve.Error()
	}

	db, err := s.commonRepo.AddTask(ctx, task.ToDB())
	if err != nil {
		return nil, InternalError(err)
	}
	return NewTask(db), nil
}

// Update updates the Task data identified by id from the query.
//
//zenrpc:tasks Task
//zenrpc:return Task
//zenrpc:500 Internal Error
//zenrpc:400 Validation Error
//zenrpc:404 Not Found
func (s TaskService) Update(ctx context.Context, task Task) (bool, error) {
	if _, err := s.byID(ctx, task.ID); err != nil {
		return false, err
	}

	if ve := s.isValid(ctx, task, true); ve.HasErrors() {
		return false, ve.Error()
	}

	ok, err := s.commonRepo.UpdateTask(ctx, task.ToDB())
	if err != nil {
		return false, InternalError(err)
	}
	return ok, nil
}

// Delete deletes the Task by its ID.
//
//zenrpc:id int64
//zenrpc:return isDeleted
//zenrpc:500 Internal Error
//zenrpc:400 Validation Error
//zenrpc:404 Not Found
func (s TaskService) Delete(ctx context.Context, id int64) (bool, error) {
	if _, err := s.byID(ctx, id); err != nil {
		return false, err
	}

	ok, err := s.commonRepo.DeleteTask(ctx, id)
	if err != nil {
		return false, InternalError(err)
	}
	return ok, err
}

// Validate verifies that Task data is valid.
//
//zenrpc:task Task
//zenrpc:return []FieldError
//zenrpc:500 Internal Error
func (s TaskService) Validate(ctx context.Context, task Task) ([]FieldError, error) {
	isUpdate := task.ID != 0
	if isUpdate {
		_, err := s.byID(ctx, task.ID)
		if err != nil {
			return nil, err
		}
	}

	ve := s.isValid(ctx, task, isUpdate)
	if ve.HasInternalError() {
		return nil, ve.Error()
	}

	return ve.Fields(), nil
}

func (s TaskService) isValid(ctx context.Context, task Task, isUpdate bool) Validator {
	var v Validator

	if v.CheckBasic(ctx, task); v.HasInternalError() {
		return v
	}

	// custom validation starts here
	return v
}
