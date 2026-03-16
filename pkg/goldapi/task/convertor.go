package task

import "apisrv/pkg/db"

func newDBTask(st *ServiceTask, isDeadLine bool) *db.Task {

	return &db.Task{
		Title:       st.Title,
		Description: st.Description,
		DeadLine:    st.DeadLine,
		IsDeadline:  &isDeadLine,
		UserID:      &st.UserId,
	}
}

func newServiceTaskFromDB(dbTask db.Task) ServiceTask {
	var userID int64
	if dbTask.UserID != nil {
		userID = *dbTask.UserID
	}

	return ServiceTask{
		ID:          dbTask.ID,
		UserId:      userID,
		Title:       dbTask.Title,
		Description: dbTask.Description,
		StatusID:    dbTask.StatusID,
		Name:        mapStatusIDToString(dbTask.StatusID),
		IsDeadline:  dbTask.IsDeadline,
		DeadLine:    dbTask.DeadLine,
		CreatedAt:   dbTask.CreatedAt,
	}
}

func newStatusesFromDB(dbStatuses []db.TaskStatus) Statuses {
	var statuses []Status
	for _, dbStatus := range dbStatuses {
		statuses = append(statuses, Status{
			ID:    dbStatus.StatusID,
			Title: dbStatus.Title,
			Alias: dbStatus.Alias,
		})
	}
	return Statuses{Statuses: statuses}
}

func mapStatusIDToString(statusID int) string {
	switch statusID {
	case 1:
		return "active"
	case 2:
		return "completed"
	case 3:
		return "deleted"
	default:
		return "unknown"
	}
}

func newServicesTaskFromDB(dbTask []db.Task) ServicesTasks {
	var tasks []ServiceTask
	for _, dbTask := range dbTask {
		tasks = append(tasks, newServiceTaskFromDB(dbTask))
	}
	return ServicesTasks{Tasks: tasks}
}
