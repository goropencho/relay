package serviceinterface

import (
	"github.com/goropencho/relay/data/request"
	"github.com/goropencho/relay/data/response"
)

type UserService interface {
	Me(userId int) response.UserResponse
	UpdateProfile(users request.UpdateUserRequest)
}
