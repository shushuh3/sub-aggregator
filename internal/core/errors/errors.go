package errors

import (
	"errors"
	"fmt"
	"net/http"
)

type AppError struct {
	Code       string
	Message    string
	HTTPStatus int
	Err        error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

var (
	ErrNotFound = &AppError{
		Code:       "NOT_FOUND",
		Message:    "resource not found",
		HTTPStatus: http.StatusNotFound,
	}

	ErrInvalidInput = &AppError{
		Code:       "INVALID_INPUT",
		Message:    "invalid input data",
		HTTPStatus: http.StatusBadRequest,
	}

	ErrInternalServer = &AppError{
		Code:       "INTERNAL_ERROR",
		Message:    "internal server error",
		HTTPStatus: http.StatusInternalServerError,
	}

	ErrDatabase = &AppError{
		Code:       "DATABASE_ERROR",
		Message:    "database operation failed",
		HTTPStatus: http.StatusInternalServerError,
	}

	ErrConflict = &AppError{
		Code:       "CONFLICT",
		Message:    "resource conflict",
		HTTPStatus: http.StatusConflict,
	}
)

func NewInvalidInput(message string, err error) *AppError {
	return &AppError{
		Code:       ErrInvalidInput.Code,
		Message:    message,
		HTTPStatus: ErrInvalidInput.HTTPStatus,
		Err:        err,
	}
}

func NewNotFound(message string, err error) *AppError {
	return &AppError{
		Code:       ErrNotFound.Code,
		Message:    message,
		HTTPStatus: ErrNotFound.HTTPStatus,
		Err:        err,
	}
}

func NewDatabase(message string, err error) *AppError {
	return &AppError{
		Code:       ErrDatabase.Code,
		Message:    message,
		HTTPStatus: ErrDatabase.HTTPStatus,
		Err:        err,
	}
}

func NewInternal(message string, err error) *AppError {
	return &AppError{
		Code:       ErrInternalServer.Code,
		Message:    message,
		HTTPStatus: ErrInternalServer.HTTPStatus,
		Err:        err,
	}
}

func NewConflict(message string, err error) *AppError {
	return &AppError{
		Code:       ErrConflict.Code,
		Message:    message,
		HTTPStatus: ErrConflict.HTTPStatus,
		Err:        err,
	}
}

func AsAppError(err error) (*AppError, bool) {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr, true
	}
	return nil, false
}
