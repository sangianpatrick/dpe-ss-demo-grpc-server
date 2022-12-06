package repository

import (
	"context"
	"database/sql"

	"github.com/sangianpatrick/dpe-ss-demo-grpc-server/entity"
	"github.com/sangianpatrick/dpe-ss-demo-grpc-server/exception"

	"github.com/sirupsen/logrus"
)

type AccountRepository interface {
	FindByEmail(ctx context.Context, email string) (account entity.Account, err error)
	Save(ctx context.Context, newAccount entity.Account) (id int64, err error)
}

type accountRepositoryImpl struct {
	logger *logrus.Logger
	db     *sql.DB
}

func NewAccountRepository(logger *logrus.Logger, db *sql.DB) AccountRepository {
	return &accountRepositoryImpl{
		logger: logger,
		db:     db,
	}
}

// FindByEmail implements AccountRepository
func (r *accountRepositoryImpl) FindByEmail(ctx context.Context, email string) (account entity.Account, err error) {
	query := "SELECT id, email, password, name, created_at, updated_at FROM account WHERE email = ?"

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		r.logger.WithFields(logrus.Fields{
			"query": query,
		}).Error(err.Error())

		return
	}

	defer stmt.Close()

	var updatedAt sql.NullTime

	if err = stmt.QueryRowContext(ctx, email).Scan(
		&account.Id,
		&account.Email,
		&account.Password,
		&account.Name,
		&account.CreatedAt,
		&updatedAt,
	); err != nil {
		r.logger.WithFields(logrus.Fields{
			"query": query,
		}).Error(err.Error())

		err = r.wrapError(err)

		return
	}

	if updatedAt.Valid {
		account.UpdatedAt = &updatedAt.Time
	}

	return
}

// Save implements AccountRepository
func (r *accountRepositoryImpl) Save(ctx context.Context, newAccount entity.Account) (id int64, err error) {
	query := "INSERT INTO account SET email = ?, password = ?, name = ?, created_at = ?"

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		r.logger.WithFields(logrus.Fields{
			"query": query,
		}).Error(err.Error())

		return
	}

	defer stmt.Close()

	result, err := stmt.ExecContext(ctx,
		newAccount.Email,
		newAccount.Password,
		newAccount.Name,
		newAccount.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return
		}

		r.logger.WithFields(logrus.Fields{
			"query": query,
		}).Error(err.Error())

		return
	}

	id, _ = result.LastInsertId()

	return
}

func (r *accountRepositoryImpl) wrapError(err error) (wrappedError error) {
	if err == sql.ErrNoRows {
		return exception.ErrNotFound
	}

	if err == context.DeadlineExceeded {
		return exception.ErrTimeout
	}

	if err == context.Canceled {
		return exception.ErrCancel
	}

	return exception.ErrInternalServer
}
