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

func NewServiceTask(dbTask db.Task) ServiceTask {
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
func NewServiceStatus(dbStatus db.TaskStatus) ServiceStatus {
	return ServiceStatus{
		ID:    dbStatus.StatusID,
		Title: dbStatus.Title,
		Alias: dbStatus.Alias,
	}
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
