package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	log.Info("Starting nuScrape")

	router := gin.Default()
	router.GET("/gesamtspielplan/:championship/:group", getGesamtspielplanChampionshipGroup)

	router.Run("localhost:8080")
}

func getGesamtspielplanChampionshipGroup(c *gin.Context) {
	c.String(http.StatusOK, "not yet implemented")
}
