package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/taskmedia/nuScrape/pkg/http/rest"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("Starting nuScrape")
}

func main() {
	router := rest.SetupRouter()
	router.Run("localhost:8080")
}
