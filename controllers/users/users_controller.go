package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mirkoremix/bookstore_users-api/domain/users"
	"github.com/mirkoremix/bookstore_users-api/services"
	"github.com/mirkoremix/bookstore_users-api/utils/errors"
)

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		// handle json error
		restErr := errors.NewBadRequestError("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		// handle user creation error
		c.JSON(saveErr.Status, saveErr)
		return
	}
	fmt.Println(user)
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("Invalid user id")
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		// handle user get error
		c.JSON(getErr.Status, getErr)
		return
	}
	fmt.Println(user)
	c.JSON(http.StatusOK, user)

}

func SearchUser(c *gin.Context) {}
