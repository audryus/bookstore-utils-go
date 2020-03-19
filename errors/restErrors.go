package errors

import (
	"errors"
	"fmt"
	"net/http"
)

//RestErr interface to represent a rest error
type RestErr interface {
	Message() string
	Status() int
	Error() string
	Causes() []interface{}
}

//RestErr Struct
type restErr struct {
	ErrMessage string        `json:"message"`
	ErrStatus  int           `json:"status"`
	ErrError   string        `json:"error"`
	ErrCauses  []interface{} `json:"causes"`
}

func (e restErr) Message() string {
	return e.ErrMessage
}
func (e restErr) Status() int {
	return e.ErrStatus
}
func (e restErr) Causes() []interface{} {
	return e.ErrCauses
}
func (e restErr) Error() string {
	return fmt.Sprintf("message: %s - status: %d - error: %s - causes [%p]",
		e.ErrMessage, e.ErrStatus, e.ErrError, e.ErrCauses)
}

//New error from default errors package
func New(detail string) error {
	return errors.New(detail)
}

//DefaultError Defautl error generator
func defaultError(message string, status int, errorDesc string, err error) RestErr {
	result := restErr{
		ErrMessage: message,
		ErrStatus:  status,
		ErrError:   errorDesc,
	}
	if err != nil {
		result.ErrCauses = append(result.ErrCauses, err.Error())
	}

	return result
}

//BadRequestError error for bad request
func BadRequestError(message string, err error) RestErr {
	return defaultError(message, http.StatusBadRequest, "bad_request", err)
}

//NotFoundError Not found resourcer error
func NotFoundError(message string) RestErr {
	return defaultError(message, http.StatusNotFound, "not_found", nil)
}

//InternalServerError 500 error
func InternalServerError(message string, err error) RestErr {
	return defaultError(message, http.StatusInternalServerError, "internal_server_error", err)
}

//NotImpemented 501 error
func NotImpemented() RestErr {
	return defaultError("Plase implement me.", http.StatusNotImplemented, "not_implemented", nil)
}

//UnautorizedError token error
func UnautorizedError() RestErr {
	return defaultError("Unable to retrieve user information from given access token.", http.StatusUnauthorized, "unauthorized", nil)
}
