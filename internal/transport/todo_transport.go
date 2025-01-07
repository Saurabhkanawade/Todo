package transport

import (
	"context"
	"encoding/json"
	"github.com/Saurabhkanawade/todo_rest_service/internal/endpoints"
	httpServer "github.com/go-kit/kit/transport/http"
	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func CreateTaskHttpHandler(endpoint endpoints.TaskEndpoints, router *mux.Router) {
	router.Handle("/task",
		httpServer.NewServer(
			endpoint.CreateTask,
			decodeCreateTask,
			encodeResponse,
		),
	).Methods(http.MethodPost)

}

func decodeCreateTask(ctx context.Context, request *http.Request) (request2 interface{}, err error) {
	var task endpoints.CreateTaskRequest

	body, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func GetTasksHttpHandlers(endpoint endpoints.TaskEndpoints, router *mux.Router) {
	router.Handle("/tasks",
		httpServer.NewServer(
			endpoint.GetAllTask,
			decodeGetAllTask,
			encodeResponse,
		),
	).Methods(http.MethodGet)
}

func decodeGetAllTask(ctx context.Context, request2 *http.Request) (request interface{}, err error) {
	return endpoints.GetAllTaskRequest{}, nil
}

func GetTaskHttpHandlers(endpoint endpoints.TaskEndpoints, router *mux.Router) {
	router.Handle("/task/{taskId}",
		httpServer.NewServer(
			endpoint.GetTask,
			decodeGetTask,
			encodeResponse,
		),
	).Methods(http.MethodGet)
}

func decodeGetTask(ctx context.Context, request2 *http.Request) (request interface{}, error error) {

	vars := mux.Vars(request2)
	taskId := vars["taskId"]

	taskUuid, err := uuid.FromString(taskId)
	if err != nil {
		return nil, err
	}

	return endpoints.GetTaskByIDParamsRequest{
		TaskId: taskUuid,
	}, nil
}

func UpdateTaskHttpHandlers(endpoint endpoints.TaskEndpoints, router *mux.Router) {
	router.Handle("/task/{taskId}",
		httpServer.NewServer(
			endpoint.UpdateTask,
			decodeUpdateTask,
			encodeResponse,
		),
	).Methods(http.MethodPut)
}

func decodeUpdateTask(ctx context.Context, request2 *http.Request) (interface{}, error) {
	logrus.Infof("decode () - decoding the update task")
	var task endpoints.UpdateTaskRequestBody

	vars := mux.Vars(request2)
	taskId := vars["taskId"]

	taskUuid, err := uuid.FromString(taskId)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(request2.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &task)
	if err != nil {
		return nil, err
	}

	task.Task.ID = taskUuid
	return task, nil
}

func DeleteTaskHttpHandlers(endpoint endpoints.TaskEndpoints, router *mux.Router) {
	router.Handle("/task/{taskId}",
		httpServer.NewServer(
			endpoint.DeleteTask,
			decodeDeleteTask,
			encodeResponse,
		),
	).Methods(http.MethodDelete)
}

func decodeDeleteTask(ctx context.Context, request2 *http.Request) (interface{}, error) {
	vars := mux.Vars(request2)
	taskId := vars["taskId"]
	taskUuid, err := uuid.FromString(taskId)
	if err != nil {
		return nil, err
	}

	return endpoints.DeleteTaskByIdRequest{
		TaskId: taskUuid,
	}, nil
}

func encodeResponse(ctx context.Context, writer http.ResponseWriter, response interface{}) error {
	if err, ok := response.(error); ok && err != nil {
		logrus.Warnf("encode() - error %v :", ok)
	}
	writer.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(writer).Encode(response)
}
