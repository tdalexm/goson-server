package domain

import (
	"fmt"
)

const (
	ErrCodeIServer    = "internal_server_error"
	ErrCodeNotFound   = "not_found"
	ErrFieldNotString = "field_not_string"
	ErrWrongParams    = "wrong_params"
	ErrSearchByID     = "invalid_search_by_id"
)

type AppError struct {
	Code string `json:"code"`
	Msg  string `json:"message"`
}

func NewAppError(code, msg string) AppError {
	return AppError{
		Code: code,
		Msg:  msg,
	}
}

func (e AppError) Error() string {
	return fmt.Sprintf("%s:, %s", e.Code, e.Msg)
}
