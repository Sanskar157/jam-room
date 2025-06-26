package models

import (
	"jam-room/utils/token"
	"html"
	"jam-room/initializers"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
  gorm.Model        //(includes ID,CreatedAt,UpdatedAt,DeletedAt)
  Name         string         // A regular string field
  Email        string    `gorm:"uniqueIndex"`     // A regular string field
  Password     string         // A regular string field
}

func (u *User) SaveUser() (*User, error) {
	var err error
    err =  initializers.DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u,nil
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))

	return nil
}

func VerifyPassword(password,hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(email string, password string) (string, error) {	
	var err error

	u := User{}

	err = initializers.DB.Where("email = ?", email).Take(&u).Error
	if err != nil {
		return "", err
	}

	token, err := token.GenerateToken(u.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}