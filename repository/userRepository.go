package repository

import (
	"errors"

	"github.com/goropencho/relay/data/request"
	"github.com/goropencho/relay/helper"
	"github.com/goropencho/relay/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user model.User)
	Update(user model.User)
	Delete(userId int)
	FindById(userId int) (user model.User, err error)
	FindAll() []model.User
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

// Delete implements UserRepository.
func (u *UserRepositoryImpl) Delete(userId int) {
	var user model.User
	result := u.DB.Where("ID = ?", userId).Delete(&user)
	helper.ErrorPanic(result.Error)
}

// FindAll implements UserRepository.
func (u *UserRepositoryImpl) FindAll() []model.User {
	var users []model.User
	u.DB.Find(&users)
	return users
}

// FindById implements UserRepository.
func (u *UserRepositoryImpl) FindById(userId int) (user model.User, err error) {
	result := u.DB.Where("ID = ?").First(&user)
	if result != nil {
		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}

// Save implements UserRepository.
func (u *UserRepositoryImpl) Save(user model.User) {
	result := u.DB.Save(&user)
	helper.ErrorPanic(result.Error)
}

// Update implements UserRepository.
func (u *UserRepositoryImpl) Update(user model.User) {
	var updateUser = request.UpdateUserRequest{
		Id:   int(user.ID),
		Name: user.Name,
	}
	result := u.DB.Model(&user).Updates(updateUser)
	helper.ErrorPanic(result.Error)
}

func NewUsersRepositoryImpl(DB *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB: DB}
}
