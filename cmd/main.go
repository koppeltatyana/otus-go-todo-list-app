package main

import (
	"fmt"
	"github.com/koppeltatyana/otus-go-todo-list-app/internal/model"
	"github.com/koppeltatyana/otus-go-todo-list-app/internal/utils"
	"math/rand/v2"
)

func main() {
	// создание менеджера пользователями
	userManager := model.NewUserManager()
	// Создание пользователей user1 и user2
	user1, _ := userManager.AddUser("user1", "user1", utils.GetRandomEmail())
	user2, _ := userManager.AddUser("user2", "pass", utils.GetRandomEmail())

	// создание менеджера задач для пользователя user1
	user1TaskManager := model.NewTaskManager(user1)
	for i := range utils.RandomInt(1, 10) {
		title := fmt.Sprintf("Task %d (user1)", i)
		description := fmt.Sprintf("Description %d (user1)", i)
		_, err := user1TaskManager.AddTask(title, description)
		if err != nil {
			fmt.Println("При добавлении задачи произошла ошибка:", err)
		}
	}
	user1TaskManager.PrintUserTasks()

	// создание менеджера задач для пользователя user2
	user2TaskManager := model.NewTaskManager(user2)
	for i := range utils.RandomInt(1, 5) {
		title := fmt.Sprintf("Task %d (user2)", i)
		description := fmt.Sprintf("Description %d (user2)", i)
		_, err := user2TaskManager.AddTask(title, description)
		if err != nil {
			fmt.Println("При добавлении задачи произошла ошибка:", err)
		}
	}
	user2TaskManager.PrintUserTasks()

	// завершение задачи (completed = true)
	randInd := rand.IntN(len(user2TaskManager.GetUserTasks()))
	_ = user2TaskManager.CompleteTask(user2TaskManager.GetUserTasks()[randInd].GetID())
	user2TaskManager.PrintUserTasks()

	// обновление задачи
	_ = user2TaskManager.UpdateTask(
		user2TaskManager.GetUserTasks()[randInd].GetID(), "New Task Title", "New Task Description",
	)
	user2TaskManager.PrintUserTasks()

	// удаление задачи
	_ = user2TaskManager.RemoveTask(user2TaskManager.GetUserTasks()[randInd].GetID())
	user2TaskManager.PrintUserTasks()

	// удаление всех задач
	err := user2TaskManager.RemoveAllTasks()
	if err != nil {
		fmt.Println("При удалении задач произошла ошибка:", err)
	}
	user2TaskManager.PrintUserTasks()
}
