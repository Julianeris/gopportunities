package main

import (
	"gopportunities/config"
	"gopportunities/router"
)

var (
	logger config.Logger
)

func main() {

	logger = *config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing config: %v", err)
		return
	}

	router.Initialize()

}
