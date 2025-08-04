package main

import (
	"github.com/koppeltatyana/otus-go-todo-list-app/internal/model"
	"github.com/koppeltatyana/otus-go-todo-list-app/internal/utils"
)

func main() {
	// Создание пользователей user1 и user2
	user1 := model.NewUser(1, "user1", "user1", utils.GetRandomEmail(), []model.Task{})
	user2 := model.NewUser(2, "user2", "user2", utils.GetRandomEmail(), []model.Task{})
	task1 := model.NewTask(1, "Title 1", "Description 1")
	task2 := model.NewTask(2, "Title 2", "Description 2")
	user1.AddTask(*task1)
	user2.AddTask(*task2)
	user1.PrintTasks()
	user2.PrintTasks()
}
