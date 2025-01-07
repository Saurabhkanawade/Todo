package services

import (
	"context"
	"github.com/Saurabhkanawade/todo_rest_service/internal/database"
	"github.com/Saurabhkanawade/todo_rest_service/internal/models"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

type TaskService interface {
	Create(ctx context.Context, task models.Task) (*models.Task, error)
	GetAll(ctx context.Context) ([]models.Task, error)
	Get(ctx context.Context, taskId uuid.UUID) (*models.Task, error)
	Update(ctx context.Context, taskId uuid.UUID, task models.Task) error
	Delete(ctx context.Context, taskId uuid.UUID) error
}

type taskServiceImpl struct {
	taskDao database.TaskDao
}

func NewTaskService(taskDao database.TaskDao) TaskService {
	return &taskServiceImpl{
		taskDao: taskDao,
	}
}

func (t taskServiceImpl) Create(ctx context.Context, task models.Task) (*models.Task, error) {
	logrus.Infof("service() - creating the tasks")

	dbTask := models.MarshalTaskModelToDao(task)
	v1, err := uuid.NewV1()
	if err != nil {
		return nil, err
	}
	dbTask.ID = v1.String()

	dbTaskRes, err := t.taskDao.Create(ctx, dbTask)
	if err != nil {
		return nil, err
	}

	modelTask := models.MarshalTaskDaoToModel(dbTaskRes)
	return &modelTask, nil
}

func (t taskServiceImpl) GetAll(ctx context.Context) ([]models.Task, error) {
	logrus.Infof("service() - fetching the tasks")

	tasks := make([]models.Task, 0)

	tasksDao, err := t.taskDao.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	for _, task := range tasksDao {
		modelTask := models.MarshalTaskDaoToModel(task)
		tasks = append(tasks, modelTask)
	}

	return tasks, nil
}

func (t taskServiceImpl) Get(ctx context.Context, taskId uuid.UUID) (*models.Task, error) {
	logrus.Infof("service() - fetching the task with taskId : %d", taskId)

	dbTask, err := t.taskDao.Get(ctx, taskId)
	if err != nil {
		return nil, err
	}

	modelTask := models.MarshalTaskDaoToModel(dbTask)

	return &modelTask, nil
}

func (t taskServiceImpl) Update(ctx context.Context, taskId uuid.UUID, task models.Task) error {
	logrus.Infof("service() - updating the task with taskId : %d", taskId)

	oldTask, err := t.Get(ctx, taskId)
	if err != nil {
		return err
	}

	if oldTask != nil {
		dbTask := models.MarshalTaskModelToDao(task)
		err = t.taskDao.Update(ctx, taskId, dbTask)
		if err != nil {
			return err
		}
	}
	return nil
}

func (t taskServiceImpl) Delete(ctx context.Context, taskId uuid.UUID) error {
	logrus.Infof("service() - deleting the task with taskId : %d", taskId)

	err := t.taskDao.Delete(ctx, taskId)
	if err != nil {
		return err
	}
	return nil
}
