package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/goropencho/relay/data/request"
	"github.com/goropencho/relay/repository"
	"github.com/goropencho/relay/service/serviceinterface"
)

type AuthServiceImpl struct {
	UserRepository repository.UserRepository
	Validator      *validator.Validate
}

// Login implements serviceinterface.AuthService.
func (*AuthServiceImpl) Login(login request.LoginRequest) {
	panic("unimplemented")
}

// Signup implements serviceinterface.AuthService.
func (*AuthServiceImpl) Signup(signup request.SignupRequest) {
	panic("unimplemented")
}

func NewAuthServiceImpl(userRepository repository.UserRepository, validator *validator.Validate) serviceinterface.AuthService {
	return &AuthServiceImpl{
		UserRepository: userRepository,
		Validator:      validator,
	}
}
