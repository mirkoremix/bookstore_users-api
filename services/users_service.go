package services

import (
	"github.com/mirkoremix/bookstore_users-api/domain/users"
	"github.com/mirkoremix/bookstore_users-api/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	if userId <= 0 {
		return nil, errors.NewBadRequestError("Invalid id value")
	}

	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil

}

func CreateUser(user users.User) (*users.User, *errors.RestErr) { // ako imas error uvek je poslednji u return-u
	if err := user.Validate(); err != nil {
		// return validation error
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func FindUser() {}
