package models

import (
	"errors"
	"strings"
	"time"
)

// User representa usuário na rede social
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name"`
	Nick      string    `json:"nick"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func (user *User) Prepare() error {

	if err := user.Validate(); err != nil {
		return err
	}

	user.TrimSpaces()

	return nil
}

func (user *User) Validate() error {

	if user.Name == "" {
		return errors.New("Username cannot be empty")
	}

	if user.Email == "" {
		return errors.New("email cannot be empty")
	}

	if user.Password == "" {
		return errors.New("Password cannot be empty")
	}

	if user.Nick == "" {
		return errors.New("Nick cannot be empty")
	}

	return nil
}

func (user *User) TrimSpaces() {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.Password = strings.TrimSpace(user.Password)
	user.Nick = strings.TrimSpace(user.Nick)
}
