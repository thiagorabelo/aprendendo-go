package models

import (
	"shortfy/passwords"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username   string    `gorm:"username;size:128;not null;uniqueIndex"`
	Password   string    `gorm:"password;size:128;not null"`
	FullName   string    `gorm:"full_name;size:256;not null"`
	Email      string    `gorm:"email;size:256;not null"`
	DateJoined time.Time `gorm:"autoCreateTime;not null"`
}

func (user *User) SetPassword(plainPassword string) error {
	password, err := passwords.Hash(plainPassword)
	if err != nil {
		return err
	}
	user.Password = password

	return nil
}
