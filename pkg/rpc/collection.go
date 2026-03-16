package rpc

//go:generate colgen -list -imports=apisrv/pkg/goldapi/task
//colgen:TaskResponse
//colgen:TaskResponse:Map(tasks.ServiceTask)
//go:generate colgen -list -imports=apisrv/pkg/goldapi/task
//colgen:StatusResponse
//colgen:StatusResponse:Map(tasks.ServiceStatus)
