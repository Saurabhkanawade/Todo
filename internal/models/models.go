package models

import "github.com/gofrs/uuid"

const (
	TaskPriorityLow    TaskPriority = "Low"
	TaskPriorityMedium TaskPriority = "Medium"
	TaskPriorityHigh   TaskPriority = "High"
)

const (
	TaskStatusPending    TaskStatus = "Pending"
	TaskStatusInProgress TaskStatus = "In Progress"
	TaskStatusCompleted  TaskStatus = "Completed"
)

type TaskPriority string
type TaskStatus string

type Task struct {
	ID          uuid.UUID    `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Priority    TaskPriority `json:"priority"`
	Status      TaskStatus   `json:"status"`
	DueDate     string       `json:"due_date"`
}

type Todo struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type User struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
}

type Tag struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
