package errors

import (
	"errors"
	"strings"

	"github.com/go-sql-driver/mysql"
)

const (
	noRowsError = "no rows in result set"
)

//ParseError erros
func ParseError(err error) *RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), noRowsError) {
			return NotFoundError("No record")
		}
		return InternalServerError("error parsing database response", err)
	}
	switch sqlErr.Number {
	case 1062:
		return BadRequestError("invalid email", sqlErr)
	}
	return InternalServerError("error processing request", errors.New("Database out"))
}
