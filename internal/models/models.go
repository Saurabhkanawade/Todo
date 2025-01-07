package models

import (
	"github.com/Saurabhkanawade/todo_rest_service/internal/dbmodels"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"github.com/volatiletech/null/v8"
	"time"
)

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
	TodoListID  uuid.UUID    `json:"todo_list_id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Priority    TaskPriority `json:"priority"`
	Status      TaskStatus   `json:"status"`
	DueDate     time.Time    `json:"due_date"`
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

func MarshalTaskDaoToModel(dbTask *dbmodels.Task) Task {

	taskUuid, err := uuid.FromString(dbTask.ID)
	if err != nil {
		logrus.Errorf("models() - error while parsing the uuid :%v", err)
	}

	todoListID, err := uuid.FromString(dbTask.TodoListID)
	if err != nil {
		logrus.Errorf("models() - error while parsing the uuid :%v", err)
	}

	task := Task{
		ID:          taskUuid,
		TodoListID:  todoListID,
		Title:       dbTask.Title,
		Description: dbTask.Description.String,
		Priority:    TaskPriority(dbTask.Ppriority.String),
		Status:      TaskStatus(dbTask.Status.String),
		DueDate:     dbTask.DueDate.Time,
	}
	return task
}

func MarshalTaskModelToDao(task Task) dbmodels.Task {
	dbTask := dbmodels.Task{
		ID:          task.ID.String(),
		TodoListID:  task.TodoListID.String(),
		Title:       task.Title,
		Description: null.StringFrom(task.Description),
		DueDate:     null.TimeFrom(task.DueDate),
		Ppriority:   null.StringFrom(string(task.Priority)),
		Status:      null.StringFrom(string(task.Status)),
		CreatedAt:   null.TimeFrom(time.Now()),
		UpdatedAt:   null.TimeFrom(time.Now()),
	}
	return dbTask
}

type Test interface {
}
