package errors

import (
	"errors"
	"net/http"
)

//RestErr Struct
type RestErr struct {
	Message string        `json:"message"`
	Status  int           `json:"status"`
	Error   string        `json:"error"`
	Causes  []interface{} `json:"causes"`
}

//New error from default errors package
func New(detail string) error {
	return errors.New(detail)
}

//DefaultError Defautl error generator
func defaultError(message string, status int, errorDesc string, err error) *RestErr {
	result := &RestErr{
		Message: message,
		Status:  status,
		Error:   errorDesc,
	}
	if err != nil {
		result.Causes = append(result.Causes, err.Error())
	}

	return result
}

//BadRequestError error for bad request
func BadRequestError(message string, err error) *RestErr {
	return defaultError(message, http.StatusBadRequest, "bad_request", err)
}

//NotFoundError Not found resourcer error
func NotFoundError(message string) *RestErr {
	return defaultError(message, http.StatusNotFound, "not_found", nil)
}

//InternalServerError 500 error
func InternalServerError(message string, err error) *RestErr {
	return defaultError(message, http.StatusInternalServerError, "internal_server_error", err)
}

//NotImpemented 501 error
func NotImpemented() *RestErr {
	return defaultError("Plase implement me.", http.StatusNotImplemented, "not_implemented", nil)
}

//UnautorizedError token error
func UnautorizedError() *RestErr {
	return defaultError("Unable to retrieve user information from given access token.", http.StatusUnauthorized, "unauthorized", nil)
}
