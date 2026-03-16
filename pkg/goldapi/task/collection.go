package task

//go:generate colgen -list -imports=apisrv/pkg/db
//colgen:ServiceTask
//colgen:ServiceTask:Map(db.Task)

//colgen:ServiceStatus
//colgen:ServiceStatus:Map(db.TaskStatus)