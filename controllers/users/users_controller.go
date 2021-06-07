package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/murilloarturo/bookstore_users_api/domain/users"
	"github.com/murilloarturo/bookstore_users_api/services"
	"github.com/murilloarturo/bookstore_users_api/utils/errors"
)

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		reqErr := errors.NewBadRequestError("invalid user id")
		c.JSON(reqErr.Status, reqErr)
		return
	}
	result, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusFound, result)
}

func CreateUser(c *gin.Context) {
	var user users.User
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	//TODO: handle error
	// 	return
	// }
	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	//	this line replace all above
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement Search!")
}
