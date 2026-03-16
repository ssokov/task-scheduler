package db

import (
	"time"

	"github.com/go-pg/pg/v10"
)

func (ts *TaskSearch) WithDeadLineInPeriodOrNull(from time.Time, to time.Time) *TaskSearch {
	ts.With(
		`(t.? IS NULL OR t.? BETWEEN ? AND ?)`,
		pg.Ident(Columns.Task.DeadLine),
		pg.Ident(Columns.Task.DeadLine),
		from,
		to,
	)
	return ts
}
