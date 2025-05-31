package main

import (
	"github.com/akyaiy/gosally-basenode/internal/logger"
)

func main() {
	log := logger.InitBaseLog()
	log.App.Info("Starting the application...")
}
