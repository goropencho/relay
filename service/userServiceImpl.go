package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/goropencho/relay/data/request"
	"github.com/goropencho/relay/data/response"
	"github.com/goropencho/relay/helper"
	"github.com/goropencho/relay/repository"
	"github.com/goropencho/relay/service/serviceinterface"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

// Me implements serviceinterface.UserService.
func (u *UserServiceImpl) Me(userId int) response.UserResponse {
	userData, err := u.UserRepository.FindById(userId)
	helper.ErrorPanic(err)

	userReponse := response.UserResponse{
		Id:    int(userData.ID),
		Name:  userData.Name,
		Email: userData.Email,
	}

	return userReponse
}

// UpdateProfile implements serviceinterface.UserService.
func (u *UserServiceImpl) UpdateProfile(users request.UpdateUserRequest) {
	userData, err := u.UserRepository.FindById(users.Id)
	helper.ErrorPanic(err)
	userData.Name = users.Name
	u.UserRepository.Update(userData)
}

func NewUserServiceImpl(userRepository repository.UserRepository, validate *validator.Validate) serviceinterface.UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
	}
}
