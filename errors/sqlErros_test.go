package errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestParseErrorNotFound(t *testing.T) {
	err := ParseError(errors.New(noRowsError))

	assert.NotNil(t, err)
	assert.EqualValues(t, "No record", err.Message)
}

func TestParseErrorInteralError(t *testing.T) {
	err := ParseError(errors.New("any not mysql error"))

	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "error parsing database response", err.Message)

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes))
	assert.EqualValues(t, "any not mysql error", err.Causes[0])

}

func TestParseErrorBadRequest(t *testing.T) {
	mysqlErro := &mysql.MySQLError{
		Number:  1062,
		Message: "My test message",
	}
	err := ParseError(mysqlErro)

	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status)
	assert.EqualValues(t, "invalid email", err.Message)

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes))
	assert.EqualValues(t, "Error 1062: My test message", err.Causes[0])

}

func TestParseErrorInternalErrorDataBase(t *testing.T) {
	mysqlErro := &mysql.MySQLError{
		Number:  1064,
		Message: "My test message",
	}
	err := ParseError(mysqlErro)

	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "error processing request", err.Message)

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes))
	assert.EqualValues(t, "Database out", err.Causes[0])

}
