package rpc

import (
	"apisrv/pkg/db"
	tasks "apisrv/pkg/goldapi/task"
	"context"
	"fmt"

	"github.com/vmkteam/embedlog"
	"github.com/vmkteam/zenrpc/v2"
)

type TaskService struct {
	zenrpc.Service
	embedlog.Logger

	taskManager *tasks.TaskManager
}

func NewTaskService(dbc db.DB, logger embedlog.Logger) *TaskService {
	return &TaskService{
		taskManager: tasks.NewTaskManager(dbc, logger),
		Logger:      logger,
	}
}

func (ts *TaskService) CreateTask(ctx context.Context, req CreateTaskRequest) (CreateTaskResponse, error) {
	ts.Log().Info("CreateTask ...")

	serviceTask := newServiceTask(req)

	taskId, err := ts.taskManager.CreateTask(ctx, &serviceTask)
	if err != nil {
		ts.Log().Error("Failed to create task", "err", err)
		return CreateTaskResponse{}, mapRPCError(err)
	}

	ts.Log().Info("Task created successfully")
	return newCreateTaskResponse(taskId), nil

}

func (ts *TaskService) GetTaskByPeriod(ctx context.Context, req GetTaskByPeriodRequest) (TaskResponseList, error) {
	ts.Log().Info("GetTaskByPeriod ...")

	if req.UserID <= 0 {
		return TaskResponseList{}, zenrpc.NewError(zenrpc.InvalidParams, fmt.Errorf("User ID must be greater than 0, got %d", req.UserID))
	}

	res, err := ts.taskManager.GetTaskByPeriod(ctx, req.UserID, req.PeriodStart, req.PeriodEnd)
	if err != nil {
		ts.Log().Error("Failed to get tasks by period", "err", err)
		return TaskResponseList{}, mapRPCError(err)
	}

	return NewTaskResponseList(res), nil
}

func (ts *TaskService) UpdateTitle(ctx context.Context, req UpdateTitleRequest) error {
	ts.Log().Info("UpdateTitle ...")

	err := ts.taskManager.UpdateTitle(ctx, req.TaskID, req.Title)
	if err != nil {
		ts.Log().Error("Failed to update task title", "err", err)
		return mapRPCError(err)
	}

	return nil
}

func (ts *TaskService) UpdateDescription(ctx context.Context, req UpdateDescriptionRequest) error {
	ts.Log().Info("UpdateDescription ...")

	err := ts.taskManager.UpdateDescription(ctx, req.TaskID, req.Description)
	if err != nil {
		ts.Log().Error("Failed to update task description", "err", err)
		return mapRPCError(err)
	}

	return nil
}

func (ts *TaskService) UpdateDeadLine(ctx context.Context, req UpdateDeadLineRequest) error {
	ts.Log().Info("UpdateDeadLine ...")

	err := ts.taskManager.UpdateDeadLine(ctx, req.TaskID, req.DeadLine)
	if err != nil {
		ts.Log().Error("Failed to update task deadline", "err", err)
		return mapRPCError(err)
	}

	return nil
}

func (ts *TaskService) DeleteTask(ctx context.Context, req DeleteTaskRequest) error {
	ts.Log().Info("DeleteTask ...")

	err := ts.taskManager.DeleteTask(ctx, req.TaskID)
	if err != nil {
		ts.Log().Error("Failed to delete task", "err", err)
		return mapRPCError(err)
	}

	return nil

}

func (ts *TaskService) UpdateStatus(ctx context.Context, req UpdateStatusRequest) error {
	ts.Log().Info("UpdateStatus ...")

	err := ts.taskManager.UpdateStatus(ctx, req.TaskID, req.StatusID)
	if err != nil {
		ts.Log().Error("Failed to update task status", "err", err)
		return mapRPCError(err)
	}

	return nil
}

func (ts *TaskService) GetStatuses(ctx context.Context) (StatusResponseList, error) {
	ts.Log().Info("GetStatuses ...")

	res, err := ts.taskManager.GetStatuses(ctx)
	if err != nil {
		ts.Log().Error("Failed to get task statuses", "err", err)
		return StatusResponseList{}, mapRPCError(err)
	}

	return NewStatusResponseList(res), nil
}
