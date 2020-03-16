package errors

import "net/http"

//RestErr Struct
type RestErr struct {
	Message string        `json:"message"`
	Status  int           `json:"status"`
	Error   string        `json:"error"`
	Causes  []interface{} `json:"causes"`
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
