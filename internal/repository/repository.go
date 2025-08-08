package repository

import (
	"github.com/koppeltatyana/otus-go-todo-list-app/internal/model"
	"github.com/koppeltatyana/otus-go-todo-list-app/internal/utils"
	"log"
)

// Entity - интерфейс сущности
type Entity interface {
	GetID() int
}

var (
	Tasks []model.Task
	Users []model.User
)

// GetTasksNextID - получить следующий id в списке задач
func GetTasksNextID() int {
	ids := make([]int, 0)
	for _, t := range Tasks {
		ids = append(ids, t.GetID())
	}
	return utils.GetMaxVal(ids) + 1
}

// GetUsersNextID - получить следующий id в списке пользователей
func GetUsersNextID() int {
	ids := make([]int, 0)
	for _, t := range Users {
		ids = append(ids, t.GetID())
	}
	return utils.GetMaxVal(ids) + 1
}

// DistributeEntities - функция-распределитель сущностей
func DistributeEntities(entities []Entity) {
	for _, entity := range entities {
		switch e := entity.(type) {
		case *model.Task:
			Tasks = append(Tasks, *e)
		case *model.User:
			Users = append(Users, *e)
		}
	}
	log.Printf("Распределено сущностей: %d", len(entities))
}
