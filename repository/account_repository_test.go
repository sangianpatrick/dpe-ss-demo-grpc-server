package repository_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/sangianpatrick/dpe-ss-demo-grpc-server/exception"
	"github.com/sangianpatrick/dpe-ss-demo-grpc-server/repository"
	"github.com/sirupsen/logrus"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestAccountRepository_FindByEmail_Success(t *testing.T) {
	timeFormat := "2006-01-02T15:04:05.999Z07:00"
	location, _ := time.LoadLocation("Asia/Jakarta")
	logger := logrus.New()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	accountEmail := "account@mail.com"

	expectedSQL := "SELECT id, email, password, name, created_at, updated_at FROM account WHERE email = ?"
	columns := []string{
		"id", "email", "password", "name", "created_at", "updated_at",
	}
	createdAt, _ := time.ParseInLocation(timeFormat, "2022-01-01 01:02:03.123", location)
	updatedAt, _ := time.ParseInLocation(timeFormat, "2022-02-02 01:02:03.123", location)
	mock.ExpectPrepare(expectedSQL).ExpectQuery().WithArgs(accountEmail).WillReturnRows(
		sqlmock.NewRows(columns).AddRow(
			int64(1), accountEmail, "password", "account name", createdAt, updatedAt,
		),
	)

	repo := repository.NewAccountRepository(logger, db)
	account, err := repo.FindByEmail(context.TODO(), accountEmail)

	assert.NoError(t, err, "should not be an error when finding account by email")
	assert.Equal(t, int64(1), account.Id, "should be a type int64 with value `1`")
	assert.Equalf(t, accountEmail, account.Email, "account email should be %d", accountEmail)
	assert.Equalf(t, createdAt.UnixMilli(), account.CreatedAt.UnixMilli(), "account createdAt should be at %d in millisecond", createdAt.UnixMilli())
	assert.NotNil(t, account.UpdatedAt, "account updatedAt should not be nil")
	assert.Equalf(t, updatedAt.UnixMilli(), account.UpdatedAt.UnixMilli(), "account updatedAt should be at %d in millisecond", updatedAt.UnixMilli())

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestAccountRepository_FindByEmail_Error_NotFound(t *testing.T) {
	logger := logrus.New()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	accountEmail := "invalid_account@mail.com"

	expectedSQL := "SELECT id, email, password, name, created_at, updated_at FROM account WHERE email = ?"
	mock.ExpectPrepare(expectedSQL).ExpectQuery().WithArgs(accountEmail).WillReturnError(sql.ErrNoRows)

	repo := repository.NewAccountRepository(logger, db)
	_, err = repo.FindByEmail(context.TODO(), accountEmail)

	assert.Error(t, err, "should be an error")
	assert.EqualError(t, err, exception.ErrNotFound.Error())

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestAccountRepository_FindByEmail_Error_ContextDeadline(t *testing.T) {
	logger := logrus.New()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	accountEmail := "invalid_account@mail.com"

	expectedSQL := "SELECT id, email, password, name, created_at, updated_at FROM account WHERE email = ?"
	mock.ExpectPrepare(expectedSQL).ExpectQuery().WithArgs(accountEmail).WillReturnError(context.DeadlineExceeded)

	repo := repository.NewAccountRepository(logger, db)
	_, err = repo.FindByEmail(context.TODO(), accountEmail)

	assert.Error(t, err, "should be an error")
	assert.EqualError(t, err, exception.ErrTimeout.Error())

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestAccountRepository_FindByEmail_Error_ContextCancel(t *testing.T) {
	logger := logrus.New()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	accountEmail := "invalid_account@mail.com"

	expectedSQL := "SELECT id, email, password, name, created_at, updated_at FROM account WHERE email = ?"
	mock.ExpectPrepare(expectedSQL).ExpectQuery().WithArgs(accountEmail).WillReturnError(context.Canceled)

	repo := repository.NewAccountRepository(logger, db)
	_, err = repo.FindByEmail(context.TODO(), accountEmail)

	assert.Error(t, err, "should be an error")
	assert.EqualError(t, err, exception.ErrCancel.Error())

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
