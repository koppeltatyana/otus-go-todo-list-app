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

func (u *User) PrintUser() {
	fmt.Printf(
		"User ID: %d, Username: %s, Email: %s\n",
		u.ID,
		u.Username,
		u.Email,
	)
}

// UpdatePassword - обновление пароля пользователя
func (u *User) UpdatePassword(newPassword string) {
	u.Password = newPassword
	u.UpdatedAt = time.Now()
}

// UpdateUser - обновление данных пользователя
func (u *User) UpdateUser(username, email string) {
	u.Username = username
	u.Email = email
	u.UpdatedAt = time.Now()
}
