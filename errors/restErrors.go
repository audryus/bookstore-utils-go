package errors

import (
	"errors"
	"fmt"
	"net/http"
)

type RestErr interface {
	Message() string
	Status() int
	Error() string
	Causes() []interface{}
}

//RestErr Struct
type restErr struct {
	message string        `json:"message"`
	status  int           `json:"status"`
	error   string        `json:"error"`
	causes  []interface{} `json:"causes"`
}

func (e restErr) Message() string {
	return e.message
}
func (e restErr) Status() int {
	return e.status
}
func (e restErr) Causes() []interface{} {
	return e.causes
}
func (e restErr) Error() string {
	return fmt.Sprintf("message: %s - status: %s - error: %s - causes [%y]",
		e.message, e.status, e.error, e.causes)
}

//New error from default errors package
func New(detail string) error {
	return errors.New(detail)
}

//DefaultError Defautl error generator
func defaultError(message string, status int, errorDesc string, err error) RestErr {
	result := restErr{
		message: message,
		status:  status,
		error:   errorDesc,
	}
	if err != nil {
		result.causes = append(result.causes, err.Error())
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
