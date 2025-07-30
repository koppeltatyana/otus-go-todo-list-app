package model

import (
	"errors"
	"fmt"
	"time"
)

// UserManager - структура для управления пользователями
type UserManager struct {
	Users  []User
	nextID int
}

// NewUserManager - конструктор для менеджера пользователями
func NewUserManager() *UserManager {
	return &UserManager{
		Users:  []User{},
		nextID: 1,
	}
}

// PrintUsers - вывод списка пользователей в консоль
func (um *UserManager) PrintUsers() {
	if len(um.Users) == 0 {
		fmt.Println("Список пользователей пуст")
		return
	}
	fmt.Printf("Users: \n")
	for _, user := range um.Users {
		fmt.Printf("ID: %d, Username: %s, Email: %s\n", user.ID, user.Username, user.Email)
	}
	fmt.Println()
}

// AddUser - добавление пользователя
func (um *UserManager) AddUser(username string, password string, email string) (User, error) {
	user, err := um.GetUserByUsername(username)
	if err != nil {
		user := NewUser(um.nextID, username, password, email)
		um.Users = append(um.Users, *user)
		um.nextID++
		return *user, nil
	}
	return user, errors.New("user already exists")
}

// GetUserById - получить пользователя по id
func (um *UserManager) GetUserById(id int) (User, error) {
	for _, user := range um.Users {
		if user.ID == id {
			return user, nil
		}
	}
	return User{}, errors.New("user does not exist")
}

// GetUserByUsername - получить пользователя по username
func (um *UserManager) GetUserByUsername(username string) (User, error) {
	for _, user := range um.Users {
		if user.Username == username {
			return user, nil
		}
	}
	return User{}, errors.New("user not found")
}

// RemoveUser - удалить пользователя
func (um *UserManager) RemoveUser(id int, userTaskManager *TaskManager) error {
	_, err := um.GetUserById(id)
	if err != nil {
		return errors.New("user does not exist")
	}
	for i, user := range um.Users {
		if user.ID == id {
			for _, task := range userTaskManager.Tasks {
				_ = userTaskManager.RemoveTask(task.ID)
			}
			um.Users = append(um.Users[:i], um.Users[i+1:]...)
		}
	}
	return nil
}

// GetUserByEmail - получить пользователя по email
func (um *UserManager) GetUserByEmail(email string) (User, error) {
	for _, user := range um.Users {
		if user.Email == email {
			return user, nil
		}
	}
	return User{}, errors.New("user does not exist")
}

// UpdateUser - обновление пользователя
func (um *UserManager) UpdateUser(id int, username, email string) error {
	for i, user := range um.Users {
		if user.ID == id {
			um.Users[i].Username = username
			um.Users[i].Email = email
			um.Users[i].UpdatedAt = time.Now()
			return nil
		}
	}
	return errors.New("user does not exist")
}

func (um *UserManager) Authenticate(username, password string) bool {
	for _, user := range um.Users {
		if user.Username == username && user.Password == password {
			return true
		}
	}
	return false
}
