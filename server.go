package main

import (
	"context"
	"github.com/Saurabhkanawade/eagle-common-service/config"
	"github.com/Saurabhkanawade/eagle-common-service/database"
	"github.com/Saurabhkanawade/eagle-common-service/server"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"time"
)

func startWebServer() {

	dbConfig := database.DbConfig{
		Host:   config.GetPostgresPass(),
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
	router := mux.NewRouter().StrictSlash(true)
	ctx := context.Background()

	// set up v1 router
	v1Router := router.PathPrefix("/v1").Subrouter()

	// server swagger page

	//transport
	//endpoints
	//services
	//dao

	err = server.StartServer(ctx, router,
		server.SetPort(config.GetServerPort()),
		server.SetReadTimeout(time.Duration(config.GetReadTimeout())),
		server.SetWriteTimeout(time.Duration(config.GetWriteTimeout())))

	if err != nil {
		logrus.Errorf("error while starting the server %s", err.Error())
	}
}
