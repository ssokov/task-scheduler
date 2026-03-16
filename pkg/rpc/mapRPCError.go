package rpc

import (
	tasks "apisrv/pkg/goldapi/task"
	"errors"

	"github.com/vmkteam/zenrpc/v2"
)

func mapRPCError(err error) error {
	switch {
	case errors.Is(err, tasks.ErrDeadLineInPast):
		return &zenrpc.Error{
			Code:    zenrpc.InvalidParams,
			Message: "Invalid params",
			Data: map[string]any{
				"field":  "deadline",
				"reason": "must be in the future",
			},
		}
	case errors.Is(err, tasks.ErrInvalidPersiods):
		return &zenrpc.Error{
			Code:    zenrpc.InvalidParams,
			Message: "Invalid params",
			Data: map[string]any{
				"field":  "periods",
				"reason": "periodStart must be before periodEnd",
			},
		}
	case errors.Is(err, tasks.ErrTaskNotFound):
		return &zenrpc.Error{
			Code:    zenrpc.InvalidParams,
			Message: "Task not found",
			Data: map[string]any{
				"reason": "no task with the given ID exists",
			},
		}

	default:
		return zenrpc.NewError(zenrpc.InternalError, err)
	}
}
