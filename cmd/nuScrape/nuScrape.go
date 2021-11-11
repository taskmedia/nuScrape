package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/taskmedia/nuScrape/pkg/http/rest"
)

var version = "dev-build"

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("Starting nuScrape")
}

func main() {
	router := rest.SetupRouter()

	router.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, "version:"+version)
		return
	})

	router.Run("0.0.0.0:8080")
}
