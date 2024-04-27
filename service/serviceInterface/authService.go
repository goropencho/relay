package serviceinterface

import "github.com/goropencho/relay/data/request"

type AuthService interface {
	Login(login request.LoginRequest)
	Signup(signup request.SignupRequest)
}
