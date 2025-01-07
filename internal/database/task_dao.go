package database

import (
	"context"
	"github.com/Saurabhkanawade/eagle-common-service/database"
	"github.com/Saurabhkanawade/todo_rest_service/internal/dbmodels"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type TaskDao interface {
	Create(ctx context.Context, task dbmodels.Task) (*dbmodels.Task, error)
	GetAll(ctx context.Context) (dbmodels.TaskSlice, error)
	Get(ctx context.Context, taskId uuid.UUID) (*dbmodels.Task, error)
	Update(ctx context.Context, taskId uuid.UUID, task dbmodels.Task) error
	Delete(ctx context.Context, taskId uuid.UUID) error
}

type taskDaoImpl struct {
	conn database.DbConnection
}

func NewTaskDao(conn database.DbConnection) TaskDao {
	return &taskDaoImpl{
		conn: conn,
	}
}

func (t taskDaoImpl) Create(ctx context.Context, task dbmodels.Task) (*dbmodels.Task, error) {
	logrus.Info("dao() - creating new task")

	err := task.Insert(ctx, t.conn.Conn, boil.Blacklist(dbmodels.TaskColumns.CreatedAt, dbmodels.TaskColumns.UpdatedAt))
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (t taskDaoImpl) GetAll(ctx context.Context) (dbmodels.TaskSlice, error) {
	logrus.Infof("dao() - fetching the tasks")

	tasks, err := dbmodels.Tasks().All(ctx, t.conn.Conn)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (t taskDaoImpl) Get(ctx context.Context, taskId uuid.UUID) (*dbmodels.Task, error) {
	logrus.Infof("dao() - fetching the task with taskId :%d", taskId)

	task, err := dbmodels.FindTask(ctx, t.conn.Conn, taskId.String())

	if err != nil {
		return nil, err
	}
	return task, nil
}

func (t taskDaoImpl) Update(ctx context.Context, taskId uuid.UUID, updateTask dbmodels.Task) error {
	logrus.Infof("dao() - updating the task with taskId :%d", taskId)

	_, err := updateTask.Update(ctx, t.conn.Conn, boil.Blacklist(dbmodels.TaskColumns.TodoListID, dbmodels.TaskColumns.CreatedAt))
	if err != nil {
		return err
	}

	return nil
}

func (t taskDaoImpl) Delete(ctx context.Context, taskId uuid.UUID) error {
	logrus.Debugf("dao() - removing the task with taskId :%d", taskId)

	task, err := t.Get(ctx, taskId)
	if err != nil {
		return err
	}

	_, err = task.Delete(ctx, t.conn.Conn)
	if err != nil {
		return err
	}
	return nil
}
