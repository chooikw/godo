package todoservice

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func initDb(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gormDB, err := gorm.Open("mysql", db)

	return gormDB, mock, err
}

func TestFindMany(t *testing.T) {
	gormDB, mock, err := initDb(t)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	db = gormDB

	// Test we have multiple todo items created by different users
	rows := sqlmock.NewRows([]string{"id", "title", "completed", "userId", "created_at", "updated_at"}).
		AddRow(1, "Sample Todo", true, "local:1", time.Now(), time.Now()).
		AddRow(2, "Sample Todo 2", true, "local:1", time.Now(), time.Now())
	mock.ExpectQuery("SELECT (.+) FROM `todos` WHERE (.*)").WithArgs("local:1").WillReturnRows(rows)

	todos := FindMany("local:1")
	expectedCount := 2
	if len(todos) != expectedCount {
		t.Errorf("Expected %d todos, but got %d", expectedCount, len(todos))
	}
}

func TestCreate(t *testing.T) {
	gormDB, mock, err := initDb(t)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	db = gormDB

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `todos`").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	todo := Todo{
		Id:        1,
		Title:     "Sample Todo",
		Completed: false,
		UserId:    "local:1",
	}

	if _, err := Create(todo); err != nil {
		t.Fatalf("error when creating todo %s", err)
		return
	}
}

func TestUpdate(t *testing.T) {
	gormDB, mock, err := initDb(t)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	db = gormDB

	rows := sqlmock.NewRows([]string{"id", "title", "completed", "userId", "created_at", "updated_at"}).
		AddRow(1, "Sample Todo", false, "local:1", time.Now(), time.Now())

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `todos` (.*)").WithArgs(true, time.Now(), 1, "local:1").WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	mock.ExpectQuery("SELECT (.+) FROM `todos` WHERE (.*)").WithArgs(1).WillReturnRows(rows)

	if _, err := Update(1, UpdateInput{Completed: true}, "local:1"); err != nil {
		t.Fatalf("error when updating todo %s", err)
		return
	}
}
