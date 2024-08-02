package error

import (
	"fmt"
	"goclub/webapp/internal/txt"
)

type (
	ErrorCode      int
	ErrorOperation string
	ErrorName      string

	AppError interface {
		error
		Code() ErrorCode
		Name() ErrorName
		Operation() ErrorOperation
		SourceErr() string
	}

	ApplicationError struct {
		code   ErrorCode
		opName ErrorOperation
		err    error
	}
)

const (
	ECNoError ErrorCode = iota
	ECUnknown
	ECStorageCreating
	ECRoutineRun
)

const (
	ENUnknown         = "unknown"
	ENStorageCreating = "storage creating"
	ENRoutineRun      = "routine run"
)

func (appErr *ApplicationError) Name() ErrorName {
	switch appErr.code {
	case ECStorageCreating:
		return ENStorageCreating
	case ECRoutineRun:
		return ENRoutineRun
	}
	return ENUnknown
}

func (appErr *ApplicationError) Error() string {
	return fmt.Sprintf("[%s(%d)] %s: %s", appErr.Name(), appErr.code, appErr.opName, appErr.err.Error())
}

func (appErr *ApplicationError) Code() ErrorCode {
	return appErr.code
}

func (appErr *ApplicationError) Operation() ErrorOperation {
	return appErr.opName
}

func (appErr *ApplicationError) SourceErr() string {
	return appErr.err.Error()
}

func FromMessage(_opName ErrorOperation, _code ErrorCode, _errMsg string, _msgArgs ...any) *ApplicationError {
	return &ApplicationError{
		opName: _opName,
		code:   _code,
		err:    fmt.Errorf(txt.T(_errMsg), _msgArgs...),
	}
}

func FromError(_opName ErrorOperation, _code ErrorCode, _err error) *ApplicationError {
	return &ApplicationError{
		opName: _opName,
		code:   _code,
		err:    _err,
	}
}
