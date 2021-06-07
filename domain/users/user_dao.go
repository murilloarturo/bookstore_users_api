package users

import (
	"fmt"
	"strings"

	"github.com/murilloarturo/bookstore_users_api/datasources/mysql/users_db"
	"github.com/murilloarturo/bookstore_users_api/utils/date"
	"github.com/murilloarturo/bookstore_users_api/utils/errors"
)

const (
	uniqueEmailCode = "email_UNIQUE"
	insertUserQuery = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
)

var (
	usersDB = make(map[int64]*User)
)

func (u *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

	result := usersDB[u.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", u.Id))
	}
	u.Id = result.Id
	u.FirstName = result.FirstName
	u.LastName = result.LastName
	u.Email = result.Email
	u.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(insertUserQuery)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date.GetNowString()
	// other way to do this
	// result, err := users_db.Client.Exec(insertUserQuery, user.FirstName, user.LastName, user.Email, user.DateCreated)
	// is better to prepare the statement, so the system returns an error when the query is invalid
	// also has better performance
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), uniqueEmailCode) {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user : %s", err.Error()))
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}
	user.Id = userId
	return nil
}
