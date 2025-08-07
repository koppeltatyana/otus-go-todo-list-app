package main

import (
	"fmt"
	"github.com/koppeltatyana/otus-go-todo-list-app/internal/repository"
	"github.com/koppeltatyana/otus-go-todo-list-app/internal/service"
	"time"
)

func main() {
	service.StartGenerating(5*time.Second, 1*time.Second)
	fmt.Println("Список пользователей: ")
	for _, u := range repository.Users {
		fmt.Print("\t")
		u.PrintUser()
	}

	fmt.Println("Список задач: ")
	for _, t := range repository.Tasks {
		fmt.Print("\t")
		t.PrintTask()
	}
}
