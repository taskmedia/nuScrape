package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/taskmedia/nuScrape/pkg/http/rest"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("Starting nuScrape")
}

func main() {
	router := gin.Default()
	rest.AddRouterGesamtspielplan(router)
	router.Run("localhost:8080")
}
