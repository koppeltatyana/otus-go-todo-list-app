package model

import (
	"fmt"
	"time"
)

// Task - структура для задачи
type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewTask(id int, title, description string) *Task {
	return &Task{
		ID:          id,
		Title:       title,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (t *Task) GetID() int {
	return t.ID
}

func (t *Task) PrintTask() {
	fmt.Printf("ID: %d, Title: %s, Description: %s, Completed: %t\n", t.ID, t.Title, t.Description, t.Completed)
}

func (t *Task) GetTitle() string {
	return t.Title
}

func (t *Task) GetDescription() string {
	return t.Description
}

func (t *Task) GetCompleted() bool {
	return t.Completed
}

func (t *Task) UpdateTask(title, description string, completed bool) {
	t.Title = title
	t.Description = description
	t.Completed = completed
	t.UpdatedAt = time.Now()
}
