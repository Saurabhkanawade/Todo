package main

import (
	"context"
	"fmt"
	"github.com/Saurabhkanawade/eagle-common-service/config"
	"github.com/Saurabhkanawade/eagle-common-service/database"
	dao "github.com/Saurabhkanawade/todo_rest_service/internal/database"
	"github.com/Saurabhkanawade/todo_rest_service/internal/endpoints"
	"github.com/Saurabhkanawade/todo_rest_service/internal/middleware"
	"github.com/Saurabhkanawade/todo_rest_service/internal/services"
	"github.com/Saurabhkanawade/todo_rest_service/internal/transport"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func startWebServer() {

	dbConfig := database.DbConfig{
		Host:   config.GetPostgresHost(),
		Port:   config.GetPostgresPort(),
		User:   config.GetPostgresUser(),
		Pass:   config.GetPostgresPass(),
		DbName: config.GetPostgresDb(),
	}
	dbConnection, err := database.InitDatabase(dbConfig)
	if err != nil {
		logrus.Fatalf("error establishing connection :%s", err.Error())
	} else {
		logrus.Debugf("successfully connected to the db :%v", dbConnection)
	}

	//throwing error will handle later

	//err = dbConnection.PingDB()
	//if err != nil {
	//	logrus.Fatalf("error pinging database :%s", err.Error())
	//}

	//setup the router
	router := mux.NewRouter()
	ctx := context.Background()

	// set up v1 router
	v1Router := router.PathPrefix("/v1").Subrouter().StrictSlash(true)

	//attach middleware
	v1Router.Use(middleware.LoggingMiddleware)

	// server swagger page
	fs := http.FileServer(http.Dir("./swagger-ui"))

	router.PathPrefix("/swagger-ui").
		Handler(http.StripPrefix("/swagger-ui", fs))

	//instantiate the endpoint
	taskDao := dao.NewTaskDao(dbConnection)

	//services
	taskService := services.NewTaskService(taskDao)

	//endpoints
	taskEndpoint := endpoints.MakeTaskEndpoints(taskService)

	//transport
	transport.CreateTaskHttpHandler(taskEndpoint, v1Router)
	transport.GetTaskHttpHandlers(taskEndpoint, v1Router)
	transport.GetTasksHttpHandlers(taskEndpoint, v1Router)
	transport.UpdateTaskHttpHandlers(taskEndpoint, v1Router)
	transport.DeleteTaskHttpHandlers(taskEndpoint, v1Router)

	startServer(ctx, v1Router)
}

// starting the server function
func startServer(ctx context.Context, router *mux.Router) {

	serverPort := fmt.Sprintf(":%s", config.GetServerPort())
	logrus.Infof("Starting server on %s.........", serverPort)

	readTimeout := time.Duration(config.GetReadTimeout())
	writeTimeout := time.Duration(config.GetWriteTimeout())

	server := &http.Server{
		Addr:         serverPort,
		ReadTimeout:  readTimeout * time.Second,
		WriteTimeout: writeTimeout * time.Second,
		Handler:      router,
	}

	if err := server.ListenAndServe(); err != nil {
		logrus.Fatalf("failed to start the server on %s %v", config.GetServerPort(), err)
	}
}
