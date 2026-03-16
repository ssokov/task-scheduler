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

func newGetTaskByPeriodResponse(tasks tasks.ServicesTasks) GetTaskByPeriodResponse {
	var taskResponses []TaskResponse
	for _, t := range tasks.Tasks {
		taskResponses = append(taskResponses, TaskResponse{
			ID:          t.ID,
			Title:       t.Title,
			Description: t.Description,
			IsDeadline:  t.IsDeadline,
			Deadline:    t.DeadLine,
			StatusID:    t.StatusID,
			Name:        t.Name,
			DeadLine:    t.DeadLine,
			CreatedAt:   t.CreatedAt,
		})
	}
	return GetTaskByPeriodResponse{Tasks: taskResponses}
}

func newStatusResponse(status tasks.Status) StatusResponse {
	return StatusResponse{
		ID:    status.ID,
		Name:  status.Title,
		Alias: status.Alias,
	}
}

func newGetStatusesResponse(statuses tasks.Statuses) GetStatusesResponse {
	var statusResponses []StatusResponse
	for _, s := range statuses.Statuses {
		statusResponses = append(statusResponses, newStatusResponse(s))
	}
	return GetStatusesResponse{Statuses: statusResponses}
}
