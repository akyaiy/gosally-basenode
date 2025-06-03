package main

import (
	"github.com/akyaiy/gosally-basenode/internal/logger"
	_ "modernc.org/sqlite"
)

func main() {
	log := logger.InitBaseLog()
	log.App.Info("Starting the application...")
}
