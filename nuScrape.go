package main

import (
	"errors"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	log.Info("Starting nuScrape")

	router := gin.Default()
	router.GET("/gesamtspielplan/:year/:championship/:group", getGesamtspielplanChampionshipGroup)

	router.Run("localhost:8080")
}

func getGesamtspielplanChampionshipGroup(c *gin.Context) {
	year := c.Param("year")
	if !validateYear(year) {
		log.WithField("year", year).Warning("year not matching with regex")
		c.String(http.StatusBadRequest, "year pattern not valid - please use YYYYYY")
		return
	}

	championship := c.Param("championship")
	if !validateChampionshipAbb(championship) {
		log.WithField("championship", championship).Warning("championship not matching with list")
		c.String(http.StatusBadRequest, "championship pattern not valid")
		return
	}

	group, err := getGroup(c.Param("group"))
	if err != nil {
		log.WithField("group", group).Warning(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.String(http.StatusOK, "not yet implemented")
}

func validateYear(y string) bool {
	return regexp.MustCompile(`\d{6}`).MatchString(y)
}

// Check if the championship abbreviation is valid
func validateChampionshipAbb(c string) bool {
	switch c {
	case
		"BHV",
		"UF",
		"OF",
		"MF",
		"OS",
		"SW",
		"AB",
		"AV",
		"OB":
		return true
	}
	return false
}

// Get and validate group
func getGroup(groupString string) (int, error) {
	g, err := strconv.Atoi(groupString)
	if err != nil {
		return -1, errors.New("group not an integer")
	}

	if g <= 0 {
		return -1, errors.New("group not a positive integer")
	}

	return g, nil
}
