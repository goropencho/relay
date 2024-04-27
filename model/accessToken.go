package model

import "gorm.io/gorm"

type AccessToken struct {
	*gorm.Model
	Token  string
	UserId string
	User   User `gorm:"foreignKey:UserId,constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
