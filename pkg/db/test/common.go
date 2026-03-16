//nolint:dupl,funlen
package test

// import (
// 	"testing"
// 	"time"

// 	"apisrv/pkg/db"

// 	"github.com/brianvoe/gofakeit/v7"
// 	"github.com/go-pg/pg/v10/orm"
// )

// type UserOpFunc func(t *testing.T, dbo orm.DB, in *db.User) Cleaner

// func User(t *testing.T, dbo orm.DB, in *db.User, ops ...UserOpFunc) (*db.User, Cleaner) {
// 	repo := db.NewCommonRepo(dbo)
// 	var cleaners []Cleaner

// 	// Fill the incoming entity
// 	if in == nil {
// 		in = &db.User{}
// 	}

// 	// Check if PKs are provided
// 	if in.ID != 0 {
// 		// Fetch the entity by PK
// 		user, err := repo.UserByID(t.Context(), in.ID, repo.FullUser())
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		// We must find the entity by PK
// 		if user == nil {
// 			t.Fatalf("the entity User is not found by provided PKs ID=%v", in.ID)
// 		}

// 		// Return if found without real cleanup
// 		return user, emptyClean
// 	}

// 	for _, op := range ops {
// 		if cl := op(t, dbo, in); cl != nil {
// 			cleaners = append(cleaners, cl)
// 		}
// 	}

// 	// Create the main entity
// 	user, err := repo.AddUser(t.Context(), in)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	return user, func() {
// 		if _, err := dbo.ModelContext(t.Context(), &db.User{ID: user.ID}).WherePK().Delete(); err != nil {
// 			t.Fatal(err)
// 		}
// 		// Clean up related entities from the last to the first
// 		for i := len(cleaners) - 1; i >= 0; i-- {
// 			cleaners[i]()
// 		}
// 	}
// }

// func WithFakeUser(t *testing.T, dbo orm.DB, in *db.User) Cleaner {
// 	if in.CreatedAt.IsZero() {
// 		in.CreatedAt = time.Now()
// 	}

// 	if in.Login == "" {
// 		in.Login = cutS(gofakeit.Word(), 64)
// 	}

// 	if in.Password == "" {
// 		in.Password = cutS(gofakeit.Password(true, true, true, false, false, 12), 64)
// 	}

// 	if in.AuthKey == "" {
// 		in.AuthKey = cutS(gofakeit.Sentence(3), 32)
// 	}

// 	if in.StatusID == 0 {
// 		in.StatusID = 1
// 	}

// 	return emptyClean
// }

// type AppUserOpFunc func(t *testing.T, dbo orm.DB, in *db.AppUser) Cleaner

// func AppUser(t *testing.T, dbo orm.DB, in *db.AppUser, ops ...AppUserOpFunc) (*db.AppUser, Cleaner) {
// 	repo := db.NewCommonRepo(dbo)
// 	var cleaners []Cleaner

// 	// Fill the incoming entity
// 	if in == nil {
// 		in = &db.AppUser{}
// 	}

// 	// Check if PKs are provided
// 	if in.ID != 0 {
// 		// Fetch the entity by PK
// 		appUser, err := repo.AppUserByID(t.Context(), in.ID, repo.FullAppUser())
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		// We must find the entity by PK
// 		if appUser == nil {
// 			t.Fatalf("the entity AppUser is not found by provided PKs ID=%v", in.ID)
// 		}

// 		// Return if found without real cleanup
// 		return appUser, emptyClean
// 	}

// 	for _, op := range ops {
// 		if cl := op(t, dbo, in); cl != nil {
// 			cleaners = append(cleaners, cl)
// 		}
// 	}

// 	// Create the main entity
// 	appUser, err := repo.AddAppUser(t.Context(), in)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	return appUser, func() {
// 		if _, err := dbo.ModelContext(t.Context(), &db.AppUser{ID: appUser.ID}).WherePK().Delete(); err != nil {
// 			t.Fatal(err)
// 		}
// 		// Clean up related entities from the last to the first
// 		for i := len(cleaners) - 1; i >= 0; i-- {
// 			cleaners[i]()
// 		}
// 	}
// }

// func WithFakeAppUser(t *testing.T, dbo orm.DB, in *db.AppUser) Cleaner {
// 	if in.Login == "" {
// 		in.Login = cutS(gofakeit.Word(), 255)
// 	}

// 	if in.Password == "" {
// 		in.Password = cutS(gofakeit.Password(true, true, true, false, false, 12), 255)
// 	}

// 	return emptyClean
// }

// type TaskFileOpFunc func(t *testing.T, dbo orm.DB, in *db.TaskFile) Cleaner

// func TaskFile(t *testing.T, dbo orm.DB, in *db.TaskFile, ops ...TaskFileOpFunc) (*db.TaskFile, Cleaner) {
// 	repo := db.NewCommonRepo(dbo)
// 	var cleaners []Cleaner

// 	// Fill the incoming entity
// 	if in == nil {
// 		in = &db.TaskFile{}
// 	}

// 	for _, op := range ops {
// 		if cl := op(t, dbo, in); cl != nil {
// 			cleaners = append(cleaners, cl)
// 		}
// 	}

// 	// Create the main entity
// 	taskFile, err := repo.AddTaskFile(t.Context(), in)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	return taskFile, func() {
// 		// Clean up related entities from the last to the first
// 		for i := len(cleaners) - 1; i >= 0; i-- {
// 			cleaners[i]()
// 		}
// 	}
// }

// type TaskOpFunc func(t *testing.T, dbo orm.DB, in *db.Task) Cleaner

// func Task(t *testing.T, dbo orm.DB, in *db.Task, ops ...TaskOpFunc) (*db.Task, Cleaner) {
// 	repo := db.NewCommonRepo(dbo)
// 	var cleaners []Cleaner

// 	// Fill the incoming entity
// 	if in == nil {
// 		in = &db.Task{}
// 	}

// 	// Check if PKs are provided
// 	if in.ID != 0 {
// 		// Fetch the entity by PK
// 		task, err := repo.TaskByID(t.Context(), in.ID, repo.FullTask())
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		// We must find the entity by PK
// 		if task == nil {
// 			t.Fatalf("the entity Task is not found by provided PKs ID=%v", in.ID)
// 		}

// 		// Return if found without real cleanup
// 		return task, emptyClean
// 	}

// 	for _, op := range ops {
// 		if cl := op(t, dbo, in); cl != nil {
// 			cleaners = append(cleaners, cl)
// 		}
// 	}

// 	// Create the main entity
// 	task, err := repo.AddTask(t.Context(), in)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	return task, func() {
// 		if _, err := dbo.ModelContext(t.Context(), &db.Task{ID: task.ID}).WherePK().Delete(); err != nil {
// 			t.Fatal(err)
// 		}
// 		// Clean up related entities from the last to the first
// 		for i := len(cleaners) - 1; i >= 0; i-- {
// 			cleaners[i]()
// 		}
// 	}
// }

// func WithFakeTask(t *testing.T, dbo orm.DB, in *db.Task) Cleaner {
// 	if in.CreatedAt.IsZero() {
// 		in.CreatedAt = time.Now()
// 	}

// 	return emptyClean
// }
