package main

import (
	"github.com/Saurabhkanawade/eagle-common-service/config"
	"github.com/sirupsen/logrus"
)

// Package todo API.
//
// # Endpoints for todo
//
// To:
//
// Schemes: http, https
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta
func main() {
	logrus.Infof("starting portal-services")

	config.LoadConfig(config.GetAppEnvLocation())

	err := config.CheckRequiredVariables()
	if err != nil {
		logrus.Fatalf("%v", err)
	}

	startWebServer()
}
