package users

import (
	"fmt"

	"github.com/murilloarturo/bookstore_users_api/datasources/mysql/users_db"
	"github.com/murilloarturo/bookstore_users_api/utils/date"
	"github.com/murilloarturo/bookstore_users_api/utils/errors"
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

func (u *User) Save() *errors.RestErr {
	current := usersDB[u.Id]
	if current != nil {
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", u.Id))
	}
	u.DateCreated = date.GetNowString()

	usersDB[u.Id] = u

	return nil
}
