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
	Tasks     []Task    `json:"tasks"`
}

// NewUser - конструктор для пользователя
func NewUser(id int, username string, email string, password string, tasks []Task) *User {
	return &User{
		ID:        id,
		Username:  username,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Tasks:     tasks,
	}
}

func (u *User) AddTask(task Task) {
	u.Tasks = append(u.Tasks, task)
}

func (u *User) PrintTasks() {
	if len(u.Tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}
	fmt.Printf("User \"%s\" Tasks: \n", u.Username)
	for _, task := range u.Tasks {
		fmt.Printf("ID: %d, Title: %s, Description: %s\n", task.ID, task.Title, task.Description)
	}
	fmt.Println()
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

func (u *User) UpdatePassword(newPassword string) {
	u.Password = newPassword
	u.UpdatedAt = time.Now()
}

func (u *User) UpdateUser(username, email string) {
	u.Username = username
	u.Email = email
	u.UpdatedAt = time.Now()
}
