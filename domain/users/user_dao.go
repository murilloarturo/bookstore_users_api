package users

import (
	"fmt"

	"github.com/murilloarturo/bookstore_users_api/datasources/mysql/users_db"
	"github.com/murilloarturo/bookstore_users_api/utils/date_utils"
	"github.com/murilloarturo/bookstore_users_api/utils/errors"
	"github.com/murilloarturo/bookstore_users_api/utils/mysql_utils"
)

const (
	insertUserQuery = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	getUserQuery    = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
	updateUserQuery    = "UPDATE users SET first_name=?, last_name=?, email=?;"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(getUserQuery)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)

	// other way to do this. important to close db connection
	// results, _ := stmt.Query(user.Id)
	// defer results.Close()

	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
		return mysql_utils.ParseError(getErr)
	}

	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(insertUserQuery)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()
	// other way to do this
	// result, err := users_db.Client.Exec(insertUserQuery, user.FirstName, user.LastName, user.Email, user.DateCreated)
	// is better to prepare the statement, so the system returns an error when the query is invalid
	// also has better performance
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}
	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(updateUserQuery)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}