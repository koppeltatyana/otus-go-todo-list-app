package model

import (
	"errors"
	"fmt"
	"time"
)

// TaskManager - структура для управления задачами
type TaskManager struct {
	Tasks  []Task
	nextID int
	user   User // ID пользователя, которому принадлежат задачи
}

// NewTaskManager - конструктор для менеджера задачами
func NewTaskManager(user User) *TaskManager {
	return &TaskManager{
		Tasks:  []Task{},
		nextID: 1,
		user:   user,
	}
}

// PrintUserTasks - вывод задач пользователя в консоль
func (tm *TaskManager) PrintUserTasks() {
	if len(tm.Tasks) == 0 {
		fmt.Printf("User \"%s\" does not have any tasks\n", tm.user.Username)
		return
	}
	fmt.Printf("User \"%s\" Tasks: \n", tm.user.Username)
	for _, task := range tm.Tasks {
		fmt.Printf(
			"ID: %d, Title: %s, Description: %s, Completed: %t\n",
			task.ID, task.Title, task.Description, task.Completed,
		)
	}
	fmt.Println()
}

// GetUserTasks - получить список всех задач пользователя
func (tm *TaskManager) GetUserTasks() []Task {
	return tm.Tasks
}

// GetTaskByTitle - получение задачи по заголовку
func (tm *TaskManager) GetTaskByTitle(title string) (Task, error) {
	for _, task := range tm.Tasks {
		if task.Title == title {
			return task, nil
		}
	}
	return Task{}, errors.New("task not found")
}

// GetTaskById - получение задачи по ID
func (tm *TaskManager) GetTaskById(id int) (Task, error) {
	for _, task := range tm.Tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return Task{}, errors.New("task not found")
}

// AddTask - добавление задачи
func (tm *TaskManager) AddTask(title, description string) (Task, error) {
	task, err := tm.GetTaskByTitle(title)
	if err != nil {
		task := NewTask(tm.nextID, title, description)
		tm.Tasks = append(tm.Tasks, *task)
		tm.nextID++
		return *task, nil
	}
	return task, errors.New("task already exists")
}

// RemoveTask - удаление задачи
func (tm *TaskManager) RemoveTask(id int) error {
	_, err := tm.GetTaskById(id)
	if err != nil {
		return err
	}
	for i, task := range tm.Tasks {
		if task.ID == id {
			tm.Tasks[i] = tm.Tasks[len(tm.Tasks)-1]
			tm.Tasks = tm.Tasks[:len(tm.Tasks)-1]
		}
	}
	return nil
}

func (tm *TaskManager) RemoveAllTasks() error {
	if len(tm.Tasks) == 0 {
		return errors.New("no tasks found")
	}
	tm.Tasks = []Task{}
	return nil
}

func (tm *TaskManager) CompleteTask(id int) error {
	for i, t := range tm.Tasks {
		if t.ID == id {
			tm.Tasks[i].Completed = true
			return nil
		}
	}
	return errors.New("task not found")
}

func (tm *TaskManager) UpdateTask(id int, title, description string) error {
	for i, t := range tm.Tasks {
		if t.ID == id {
			tm.Tasks[i].Completed = true
			tm.Tasks[i].Title = title
			tm.Tasks[i].Description = description
			tm.Tasks[i].UpdatedAt = time.Now()
			return nil
		}
	}
	return errors.New("task not found")
}
