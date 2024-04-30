package repository

import (
	"errors"

	"github.com/goropencho/relay/helper"
	"github.com/goropencho/relay/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AccessTokenRepository interface {
	FindByToken(token string) (AccessToken model.AccessToken, err error)
	Create(accessToken model.AccessToken)
}

type AccessTokenRepositoryImpl struct {
	DB *gorm.DB
}

func (u *AccessTokenRepositoryImpl) FindById(token string) (accessToken model.AccessToken, err error) {
	result := u.DB.Where("Token = ?").First(&accessToken)
	if result != nil {
		return accessToken, nil
	} else {
		return accessToken, errors.New("user not found")
	}
}

func (u *AccessTokenRepositoryImpl) Create(accessToken model.AccessToken) (token string, err error) {
	var accesstoken model.AccessToken
	result := u.DB.Model(&accesstoken).Clauses(clause.Returning{}).Save(&accessToken)
	helper.ErrorPanic(result.Error)
	return accesstoken.Token, nil
}
