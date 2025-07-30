package model

import (
	"fmt"
	"time"
)

// User - структура для пользователя
type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewUser - конструктор для пользователя
func NewUser(id int, username string, email string, password string) *User {
	return &User{
		ID:        id,
		Username:  username,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// GetID - геттер для id
func (u *User) GetID() int {
	return u.ID
}

// GetUsername - геттер для username
func (u *User) GetUsername() string {
	return u.Username
}

// GetEmail - геттер для email
func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) PrintUser() {
	fmt.Printf("User ID: %d, Username: %s, Email: %s\n", u.ID, u.Username, u.Email)
}
