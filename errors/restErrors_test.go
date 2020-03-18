package errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	errorMessage     = "this is the message"
	errorDescription = "this is the error"
)

func TestDefaultError(t *testing.T) {
	err := defaultError(errorMessage, 666, errorDescription, errors.New("new error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, 666, err.Status)
	assert.EqualValues(t, errorMessage, err.Message)
	assert.EqualValues(t, errorDescription, err.Error)

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes()))
	assert.EqualValues(t, "new error", err.Causes()[0])
}

//BadRequestError error for bad request
func TestBadRequestError(t *testing.T) {
	err := BadRequestError(errorMessage, errors.New("something is missing"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status)
	assert.EqualValues(t, errorMessage, err.Message)
	assert.EqualValues(t, "bad_request", err.Error)

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes()))
	assert.EqualValues(t, "something is missing", err.Causes()[0])

	err = InternalServerError(errorMessage, nil)
	assert.Nil(t, err.Causes)

}

//NotFound Not found resourcer error
func TestNotFound(t *testing.T) {
	err := NotFoundError(errorMessage)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, errorMessage, err.Message)
	assert.EqualValues(t, "not_found", err.Error)
}

//InternalServerError 500 error
func TestInternalServerError(t *testing.T) {
	err := InternalServerError(errorMessage, errors.New("database error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, errorMessage, err.Message)
	assert.EqualValues(t, "internal_server_error", err.Error)

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes()))
	assert.EqualValues(t, "database error", err.Causes()[0])

	err = InternalServerError(errorMessage, nil)
	assert.Nil(t, err.Causes)
}

func TestNotImplemented(t *testing.T) {
	err := NotImpemented()
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotImplemented, err.Status)
	assert.EqualValues(t, "Plase implement me.", err.Message)
	assert.EqualValues(t, "not_implemented", err.Error)
}

func TestUnauthorized(t *testing.T) {
	err := UnautorizedError()
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.Status)
	assert.EqualValues(t, "Unable to retrieve user information from given access token.", err.Message)
	assert.EqualValues(t, "unauthorized", err.Error)
}
