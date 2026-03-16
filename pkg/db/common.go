package db

import (
	"context"
	"errors"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type CommonRepo struct {
	db      orm.DB
	filters map[string][]Filter
	sort    map[string][]SortField
	join    map[string][]string
}

// NewCommonRepo returns new repository
func NewCommonRepo(db orm.DB) CommonRepo {
	return CommonRepo{
		db: db,
		filters: map[string][]Filter{
			Tables.User.Name:       {StatusFilter},
			Tables.Task.Name:       {StatusFilter},
			Tables.TaskStatus.Name: {StatusFilter},
		},
		sort: map[string][]SortField{
			Tables.User.Name:       {{Column: Columns.User.CreatedAt, Direction: SortDesc}},
			Tables.AppUser.Name:    {{Column: Columns.AppUser.ID, Direction: SortDesc}},
			Tables.Task.Name:       {{Column: Columns.Task.Title, Direction: SortAsc}},
			Tables.TaskStatus.Name: {{Column: Columns.TaskStatus.Title, Direction: SortAsc}},
		},
		join: map[string][]string{
			Tables.User.Name:       {TableColumns},
			Tables.AppUser.Name:    {TableColumns},
			Tables.TaskFile.Name:   {TableColumns, Columns.TaskFile.Task, Columns.TaskFile.File},
			Tables.Task.Name:       {TableColumns, Columns.Task.User},
			Tables.TaskStatus.Name: {TableColumns},
		},
	}
}

// WithTransaction is a function that wraps CommonRepo with pg.Tx transaction.
func (cr CommonRepo) WithTransaction(tx *pg.Tx) CommonRepo {
	cr.db = tx
	return cr
}

// WithEnabledOnly is a function that adds "statusId"=1 as base filter.
func (cr CommonRepo) WithEnabledOnly() CommonRepo {
	f := make(map[string][]Filter, len(cr.filters))
	for i := range cr.filters {
		f[i] = make([]Filter, len(cr.filters[i]))
		copy(f[i], cr.filters[i])
		f[i] = append(f[i], StatusEnabledFilter)
	}
	cr.filters = f

	return cr
}

/*** User ***/

// FullUser returns full joins with all columns
func (cr CommonRepo) FullUser() OpFunc {
	return WithColumns(cr.join[Tables.User.Name]...)
}

// DefaultUserSort returns default sort.
func (cr CommonRepo) DefaultUserSort() OpFunc {
	return WithSort(cr.sort[Tables.User.Name]...)
}

// UserByID is a function that returns User by ID(s) or nil.
func (cr CommonRepo) UserByID(ctx context.Context, id int, ops ...OpFunc) (*User, error) {
	return cr.OneUser(ctx, &UserSearch{ID: &id}, ops...)
}

// OneUser is a function that returns one User by filters. It could return pg.ErrMultiRows.
func (cr CommonRepo) OneUser(ctx context.Context, search *UserSearch, ops ...OpFunc) (*User, error) {
	obj := &User{}
	err := buildQuery(ctx, cr.db, obj, search, cr.filters[Tables.User.Name], PagerTwo, ops...).Select()

	if errors.Is(err, pg.ErrMultiRows) {
		return nil, err
	} else if errors.Is(err, pg.ErrNoRows) {
		return nil, nil
	}

	return obj, err
}

// UsersByFilters returns User list.
func (cr CommonRepo) UsersByFilters(ctx context.Context, search *UserSearch, pager Pager, ops ...OpFunc) (users []User, err error) {
	err = buildQuery(ctx, cr.db, &users, search, cr.filters[Tables.User.Name], pager, ops...).Select()
	return
}

// CountUsers returns count
func (cr CommonRepo) CountUsers(ctx context.Context, search *UserSearch, ops ...OpFunc) (int, error) {
	return buildQuery(ctx, cr.db, &User{}, search, cr.filters[Tables.User.Name], PagerOne, ops...).Count()
}

// AddUser adds User to DB.
func (cr CommonRepo) AddUser(ctx context.Context, user *User, ops ...OpFunc) (*User, error) {
	q := cr.db.ModelContext(ctx, user)
	if len(ops) == 0 {
		q = q.ExcludeColumn(Columns.User.CreatedAt)
	}
	applyOps(q, ops...)
	_, err := q.Insert()

	return user, err
}

// UpdateUser updates User in DB.
func (cr CommonRepo) UpdateUser(ctx context.Context, user *User, ops ...OpFunc) (bool, error) {
	q := cr.db.ModelContext(ctx, user).WherePK()
	if len(ops) == 0 {
		q = q.ExcludeColumn(Columns.User.CreatedAt)
	}
	applyOps(q, ops...)
	res, err := q.Update()
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, err
}

// DeleteUser set statusId to deleted in DB.
func (cr CommonRepo) DeleteUser(ctx context.Context, id int) (deleted bool, err error) {
	user := &User{ID: id, StatusID: StatusDeleted}

	return cr.UpdateUser(ctx, user, WithColumns(Columns.User.StatusID))
}

/*** AppUser ***/

// FullAppUser returns full joins with all columns
func (cr CommonRepo) FullAppUser() OpFunc {
	return WithColumns(cr.join[Tables.AppUser.Name]...)
}

// DefaultAppUserSort returns default sort.
func (cr CommonRepo) DefaultAppUserSort() OpFunc {
	return WithSort(cr.sort[Tables.AppUser.Name]...)
}

// AppUserByID is a function that returns AppUser by ID(s) or nil.
func (cr CommonRepo) AppUserByID(ctx context.Context, id int64, ops ...OpFunc) (*AppUser, error) {
	return cr.OneAppUser(ctx, &AppUserSearch{ID: &id}, ops...)
}

// OneAppUser is a function that returns one AppUser by filters. It could return pg.ErrMultiRows.
func (cr CommonRepo) OneAppUser(ctx context.Context, search *AppUserSearch, ops ...OpFunc) (*AppUser, error) {
	obj := &AppUser{}
	err := buildQuery(ctx, cr.db, obj, search, cr.filters[Tables.AppUser.Name], PagerTwo, ops...).Select()

	if errors.Is(err, pg.ErrMultiRows) {
		return nil, err
	} else if errors.Is(err, pg.ErrNoRows) {
		return nil, nil
	}

	return obj, err
}

// AppUsersByFilters returns AppUser list.
func (cr CommonRepo) AppUsersByFilters(ctx context.Context, search *AppUserSearch, pager Pager, ops ...OpFunc) (appUsers []AppUser, err error) {
	err = buildQuery(ctx, cr.db, &appUsers, search, cr.filters[Tables.AppUser.Name], pager, ops...).Select()
	return
}

// CountAppUsers returns count
func (cr CommonRepo) CountAppUsers(ctx context.Context, search *AppUserSearch, ops ...OpFunc) (int, error) {
	return buildQuery(ctx, cr.db, &AppUser{}, search, cr.filters[Tables.AppUser.Name], PagerOne, ops...).Count()
}

// AddAppUser adds AppUser to DB.
func (cr CommonRepo) AddAppUser(ctx context.Context, appUser *AppUser, ops ...OpFunc) (*AppUser, error) {
	q := cr.db.ModelContext(ctx, appUser)
	applyOps(q, ops...)
	_, err := q.Insert()

	return appUser, err
}

// UpdateAppUser updates AppUser in DB.
func (cr CommonRepo) UpdateAppUser(ctx context.Context, appUser *AppUser, ops ...OpFunc) (bool, error) {
	q := cr.db.ModelContext(ctx, appUser).WherePK()
	if len(ops) == 0 {
		q = q.ExcludeColumn(Columns.AppUser.ID)
	}
	applyOps(q, ops...)
	res, err := q.Update()
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, err
}

// DeleteAppUser deletes AppUser from DB.
func (cr CommonRepo) DeleteAppUser(ctx context.Context, id int64) (deleted bool, err error) {
	appUser := &AppUser{ID: id}

	res, err := cr.db.ModelContext(ctx, appUser).WherePK().Delete()
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, err
}

/*** TaskFile ***/

// FullTaskFile returns full joins with all columns
func (cr CommonRepo) FullTaskFile() OpFunc {
	return WithColumns(cr.join[Tables.TaskFile.Name]...)
}

// DefaultTaskFileSort returns default sort.
func (cr CommonRepo) DefaultTaskFileSort() OpFunc {
	return WithSort(cr.sort[Tables.TaskFile.Name]...)
}

// OneTaskFile is a function that returns one TaskFile by filters. It could return pg.ErrMultiRows.
func (cr CommonRepo) OneTaskFile(ctx context.Context, search *TaskFileSearch, ops ...OpFunc) (*TaskFile, error) {
	obj := &TaskFile{}
	err := buildQuery(ctx, cr.db, obj, search, cr.filters[Tables.TaskFile.Name], PagerTwo, ops...).Select()

	if errors.Is(err, pg.ErrMultiRows) {
		return nil, err
	} else if errors.Is(err, pg.ErrNoRows) {
		return nil, nil
	}

	return obj, err
}

// TaskFilesByFilters returns TaskFile list.
func (cr CommonRepo) TaskFilesByFilters(ctx context.Context, search *TaskFileSearch, pager Pager, ops ...OpFunc) (taskFiles []TaskFile, err error) {
	err = buildQuery(ctx, cr.db, &taskFiles, search, cr.filters[Tables.TaskFile.Name], pager, ops...).Select()
	return
}

// CountTaskFiles returns count
func (cr CommonRepo) CountTaskFiles(ctx context.Context, search *TaskFileSearch, ops ...OpFunc) (int, error) {
	return buildQuery(ctx, cr.db, &TaskFile{}, search, cr.filters[Tables.TaskFile.Name], PagerOne, ops...).Count()
}

// AddTaskFile adds TaskFile to DB.
func (cr CommonRepo) AddTaskFile(ctx context.Context, taskFile *TaskFile, ops ...OpFunc) (*TaskFile, error) {
	q := cr.db.ModelContext(ctx, taskFile)
	applyOps(q, ops...)
	_, err := q.Insert()

	return taskFile, err
}

// UpdateTaskFile updates TaskFile in DB.
func (cr CommonRepo) UpdateTaskFile(ctx context.Context, taskFile *TaskFile, ops ...OpFunc) (bool, error) {
	q := cr.db.ModelContext(ctx, taskFile).WherePK()
	applyOps(q, ops...)
	res, err := q.Update()
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, err
}

/*** Task ***/

// FullTask returns full joins with all columns
func (cr CommonRepo) FullTask() OpFunc {
	return WithColumns(cr.join[Tables.Task.Name]...)
}

// DefaultTaskSort returns default sort.
func (cr CommonRepo) DefaultTaskSort() OpFunc {
	return WithSort(cr.sort[Tables.Task.Name]...)
}

// TaskByID is a function that returns Task by ID(s) or nil.
func (cr CommonRepo) TaskByID(ctx context.Context, id int64, ops ...OpFunc) (*Task, error) {
	return cr.OneTask(ctx, &TaskSearch{ID: &id}, ops...)
}

// OneTask is a function that returns one Task by filters. It could return pg.ErrMultiRows.
func (cr CommonRepo) OneTask(ctx context.Context, search *TaskSearch, ops ...OpFunc) (*Task, error) {
	obj := &Task{}
	err := buildQuery(ctx, cr.db, obj, search, cr.filters[Tables.Task.Name], PagerTwo, ops...).Select()

	if errors.Is(err, pg.ErrMultiRows) {
		return nil, err
	} else if errors.Is(err, pg.ErrNoRows) {
		return nil, nil
	}

	return obj, err
}

// TasksByFilters returns Task list.
func (cr CommonRepo) TasksByFilters(ctx context.Context, search *TaskSearch, pager Pager, ops ...OpFunc) (tasks []Task, err error) {
	err = buildQuery(ctx, cr.db, &tasks, search, cr.filters[Tables.Task.Name], pager, ops...).Select()
	return
}

// CountTasks returns count
func (cr CommonRepo) CountTasks(ctx context.Context, search *TaskSearch, ops ...OpFunc) (int, error) {
	return buildQuery(ctx, cr.db, &Task{}, search, cr.filters[Tables.Task.Name], PagerOne, ops...).Count()
}

// AddTask adds Task to DB.
func (cr CommonRepo) AddTask(ctx context.Context, task *Task, ops ...OpFunc) (*Task, error) {
	q := cr.db.ModelContext(ctx, task)
	applyOps(q, ops...)
	_, err := q.Insert()

	return task, err
}

// UpdateTask updates Task in DB.
func (cr CommonRepo) UpdateTask(ctx context.Context, task *Task, ops ...OpFunc) (bool, error) {
	q := cr.db.ModelContext(ctx, task).WherePK()
	if len(ops) == 0 {
		q = q.ExcludeColumn(Columns.Task.ID)
	}
	applyOps(q, ops...)
	res, err := q.Update()
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, err
}

// DeleteTask set statusId to deleted in DB.
func (cr CommonRepo) DeleteTask(ctx context.Context, id int64) (deleted bool, err error) {
	task := &Task{ID: id, StatusID: StatusDeleted}

	return cr.UpdateTask(ctx, task, WithColumns(Columns.Task.StatusID))
}

/*** TaskStatus ***/

// FullTaskStatus returns full joins with all columns
func (cr CommonRepo) FullTaskStatus() OpFunc {
	return WithColumns(cr.join[Tables.TaskStatus.Name]...)
}

// DefaultTaskStatusSort returns default sort.
func (cr CommonRepo) DefaultTaskStatusSort() OpFunc {
	return WithSort(cr.sort[Tables.TaskStatus.Name]...)
}

// OneTaskStatus is a function that returns one TaskStatus by filters. It could return pg.ErrMultiRows.
func (cr CommonRepo) OneTaskStatus(ctx context.Context, search *TaskStatusSearch, ops ...OpFunc) (*TaskStatus, error) {
	obj := &TaskStatus{}
	err := buildQuery(ctx, cr.db, obj, search, cr.filters[Tables.TaskStatus.Name], PagerTwo, ops...).Select()

	if errors.Is(err, pg.ErrMultiRows) {
		return nil, err
	} else if errors.Is(err, pg.ErrNoRows) {
		return nil, nil
	}

	return obj, err
}

// TaskStatusesByFilters returns TaskStatus list.
func (cr CommonRepo) TaskStatusesByFilters(ctx context.Context, search *TaskStatusSearch, pager Pager, ops ...OpFunc) (taskStatuses []TaskStatus, err error) {
	err = buildQuery(ctx, cr.db, &taskStatuses, search, cr.filters[Tables.TaskStatus.Name], pager, ops...).Select()
	return
}

// CountTaskStatuses returns count
func (cr CommonRepo) CountTaskStatuses(ctx context.Context, search *TaskStatusSearch, ops ...OpFunc) (int, error) {
	return buildQuery(ctx, cr.db, &TaskStatus{}, search, cr.filters[Tables.TaskStatus.Name], PagerOne, ops...).Count()
}

// AddTaskStatus adds TaskStatus to DB.
func (cr CommonRepo) AddTaskStatus(ctx context.Context, taskStatus *TaskStatus, ops ...OpFunc) (*TaskStatus, error) {
	q := cr.db.ModelContext(ctx, taskStatus)
	applyOps(q, ops...)
	_, err := q.Insert()

	return taskStatus, err
}

// UpdateTaskStatus updates TaskStatus in DB.
func (cr CommonRepo) UpdateTaskStatus(ctx context.Context, taskStatus *TaskStatus, ops ...OpFunc) (bool, error) {
	q := cr.db.ModelContext(ctx, taskStatus).WherePK()
	applyOps(q, ops...)
	res, err := q.Update()
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, err
}
