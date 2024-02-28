package dto

import (
	"errors"
	"net/mail"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// register
func (r *RegisterRequest) Validate() error {

	var errStr []string

	if r.Username == "" {
		errStr = append(errStr, "username is required")
	}

	if r.Email == "" {
		errStr = append(errStr, "email is required")
	}

	if _, err := mail.ParseAddress(r.Email); err != nil {
		errStr = append(errStr, "email not valid")
	}

	if r.Password == "" {
		errStr = append(errStr, "password is required")
	}

	if len(errStr) > 0 {
		return errors.New(strings.Join(errStr, ","))
	}

	return nil
}

func (l *RegisterRequest) GetHashPassword() (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(l.Password), 14)

	return string(bytes), err
}

func (l *RegisterRequest) PrepareDataDB(passwordHash string, isAdmin bool) UserDB {

	user := UserDB{
		Username:   l.Username,
		Email:      l.Email,
		Phone:      l.Phone,
		Password:   passwordHash,
		UserRoleID: 0,
	}

	if isAdmin {
		user.UserRoleID = 1
	}

	user.CreatedAt = time.Now()
	user.CreatedBy = l.Username
	user.ModifiedAt = time.Now()
	user.ModifiedBy = l.Username

	return user
}

// login
func (l *LoginRequest) UsernameIsEmail() bool {
	_, err := mail.ParseAddress(l.Username)

	return err == nil
}

func (l *LoginRequest) CheckPasswordHash(userPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(l.Password))

	return err == nil
}
