package users

import (
	"fmt"

	"github.com/mirkoremix/bookstore_users-api/utils/errors"
)

var (
	usersDb = make(map[int64]*User)
)

// da nisi prosledio kao pointer (*User) to bi znacilo da saljes kopiju objekta user.
func (user *User) Get() *errors.RestErr {
	result := usersDb[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

func (user *User) Save() *errors.RestErr {
	current := usersDb[user.Id]
	if current != nil {
		if user.Email == current.Email {
			return errors.NewBadRequestError(fmt.Sprintf("user with email %s already exist", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user with id %d already exist", user.Id))
	}
	usersDb[user.Id] = user
	return nil
}
