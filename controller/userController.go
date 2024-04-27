package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/goropencho/relay/data/request"
	"github.com/goropencho/relay/data/response"
	"github.com/goropencho/relay/helper"
	serviceinterface "github.com/goropencho/relay/service/serviceInterface"
)

type UserController struct {
	userService serviceinterface.UserService
}

func NewUserController(service serviceinterface.UserService) *UserController {
	return &UserController{
		userService: service,
	}
}

func (controller *UserController) Me(ctx *gin.Context) {
	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)

	userResponse := controller.userService.Me(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   userResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

func (controller *UserController) UpdateProfile(ctx *gin.Context) {
	updateUserRequest := request.UpdateUserRequest{}
	err := ctx.ShouldBind(updateUserRequest)
	helper.ErrorPanic(err)

	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)

	updateUserRequest.Id = id

	controller.userService.UpdateProfile(updateUserRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
