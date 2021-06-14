package mysql_utils

import (
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/murilloarturo/bookstore_users_api/utils/errors"
)

const (
	NoRowsErrorCode = "no rows"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, isSqlErr := err.(*mysql.MySQLError)
	if !isSqlErr {
		if strings.Contains(err.Error(), NoRowsErrorCode) {
			return errors.NewNotFoundError("No record matching given id")
		}
		return errors.NewInternalServerError(fmt.Sprintf("error parsing database response"))
	}
	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("invalid data")
	default:
		return errors.NewInternalServerError("error processing request")
	}
}
