package main

import (
	"github.com/Saurabhkanawade/eagle-common-service/config"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Infof("starting portal-services")

	//check required vars
	err := config.CheckRequiredVariables()
	if err != nil {
		logrus.Fatalf("%v", err)
	}

	startWebServer()
}
