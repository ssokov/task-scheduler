package rpc

import tasks "apisrv/pkg/goldapi/task"

func newServiceTask(createTaskRequest CreateTaskRequest) tasks.ServiceTask {
	return tasks.ServiceTask{
		UserId:      createTaskRequest.UserID,
		Title:       createTaskRequest.Title,
		Description: createTaskRequest.Description,
		DeadLine:    createTaskRequest.Deadline,
	}
}

func newCreateTaskResponse(taskId int64) CreateTaskResponse {
	return CreateTaskResponse{
		TaskID: taskId,
	}
}

func NewTaskResponse(dbTask tasks.ServiceTask) TaskResponse {
	return TaskResponse{
		ID:          dbTask.ID,
		Title:       dbTask.Title,
		Description: dbTask.Description,
		IsDeadline:  dbTask.IsDeadline,
		Deadline:    dbTask.DeadLine,
		StatusID:    dbTask.StatusID,
		Name:        dbTask.Name,
		DeadLine:    dbTask.DeadLine,
		CreatedAt:   dbTask.CreatedAt,
	}
}

func NewStatusResponse(dbStatus tasks.ServiceStatus) StatusResponse {
	return StatusResponse{
		ID:    dbStatus.ID,
		Name:  dbStatus.Title,
		Alias: dbStatus.Alias,
	}
}
