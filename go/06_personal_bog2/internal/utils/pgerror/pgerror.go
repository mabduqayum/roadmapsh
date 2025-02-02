package pgerror

import (
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

var (
	ErrUniqueViolation     = errors.New("unique constraint violation")
	ErrForeignKeyViolation = errors.New("foreign key constraint violation")
	ErrNoRowsFound         = errors.New("no rows found")
	ErrConstraint          = errors.New("constraint violation")
)

type WrappedError struct {
	err     error
	message string
}

// for db constraint errors send client error instead of server errors
func (we *WrappedError) Error() string {
	if we.message != "" {
		return fmt.Sprint(we.message)
	}
	return we.err.Error()
}

func (we *WrappedError) Unwrap() error {
	return we.err
}

func (we *WrappedError) WithMessage(msg string) *WrappedError {
	we.message = msg
	return we
}

func Wrap(err error) *WrappedError {
	if err == nil {
		return nil
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505":
			return &WrappedError{err: fmt.Errorf("%w: %v", ErrUniqueViolation, pgErr.Detail)}
		case "23503":
			return &WrappedError{err: fmt.Errorf("%w: %v", ErrForeignKeyViolation, pgErr.Detail)}
		case "23514":
			return &WrappedError{err: fmt.Errorf("%w: %v", ErrNoRowsFound, pgErr.Message)}
		}
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return &WrappedError{err: fmt.Errorf("%w: %v", ErrNoRowsFound, err)}
	}

	return &WrappedError{err: err}
}
