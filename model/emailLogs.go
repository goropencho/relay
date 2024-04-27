package model

import "gorm.io/gorm"

type EmailLogs struct {
	*gorm.Model
	EmailId string
	Status  string
	UserId  string
	User    User `gorm:"foreignKey:UserId,constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
