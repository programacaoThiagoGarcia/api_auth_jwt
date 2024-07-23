package models

import (
	"errors"
	"html"
	"strings"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;unique" json:"password"`
}

// Register USER

func (u *User) SaveUser() (*User, error) {
	err := DB.Create(&u).Error

	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BeforeSave() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	var user User
	result := DB.Where("Username = ?", u.Username).Find(&user)
	if result.RowsAffected > 0 {
		return errors.New("User not available")
	}

	u.Password = string(hashedPassword)
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	return nil
}

//Query USER

func (u *User) CheckLogin() error {
	var user User
	result := DB.Where("Username = ?", u.Username).Find(&user)
	if result.Error != nil {
		return result.Error
	}

	err := verifyPassword(u.Password, user.Password)
	if err != nil {
		return err
	}
	return nil

}

func verifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
