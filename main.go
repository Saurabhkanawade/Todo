package main

import (
	"github.com/Saurabhkanawade/eagle-common-service/config"
	"github.com/sirupsen/logrus"
)
//main
func main() {
	logrus.Infof("starting portal-services")

	config.LoadConfig(config.GetAppEnvLocation())

	//check required vars
	err := config.CheckRequiredVariables()
	if err != nil {
		logrus.Fatalf("%v", err)
	}

	startWebServer()
}
