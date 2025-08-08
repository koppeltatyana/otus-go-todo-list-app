package service

import (
	"fmt"
	"github.com/koppeltatyana/otus-go-todo-list-app/internal/model"
	"github.com/koppeltatyana/otus-go-todo-list-app/internal/repository"
	"github.com/koppeltatyana/otus-go-todo-list-app/internal/utils"
	"log"
	"math/rand"
	"time"
)

// GenerateEntities - генерация сущностей
func GenerateEntities(entityCount int) {
	var entities []repository.Entity
	taskCount, userCount := 0, 0
	log.Printf("Генерация %d сушностей", entityCount)
	for i := range entityCount {
		if rand.Intn(2) == 0 {
			nextID := repository.GetTasksNextID()
			entity := model.NewTask(
				nextID+i,
				fmt.Sprintf("Task %d", nextID+i),
				fmt.Sprintf("Task %d description", nextID+i),
			)
			entities = append(entities, entity)
			taskCount++
		} else {
			nextID := repository.GetUsersNextID()
			entity := model.NewUser(
				nextID+i,
				fmt.Sprintf("User %d", nextID+i),
				utils.GetRandomEmail(),
				fmt.Sprintf("user%dpass", nextID+i),
			)
			entities = append(entities, entity)
			userCount++
		}
	}
	log.Printf("Сгенерировано задач: %d, сгенерировано пользователей: %d", taskCount, userCount)

	// Передача в функцию-распределитель
	repository.DistributeEntities(entities)
}

// StartGenerating - генерация сущностей с заданным интервалом и таймаутом
func StartGenerating(timeout time.Duration, interval time.Duration) {
	log.Println("Начало генерации")
	startTime := time.Now()
	for {
		if time.Since(startTime) >= timeout {
			break
		}
		entityCount := utils.RandomInt(1, 5)
		GenerateEntities(entityCount)
		time.Sleep(interval)
	}
	log.Println("Конец генерации")
}
