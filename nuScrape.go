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
	router.GET("/gesamtspielplan/:championship/:group", getGesamtspielplanChampionshipGroup)

	router.Run("localhost:8080")
}

func getGesamtspielplanChampionshipGroup(c *gin.Context) {
	championship := c.Param("championship")
	if !validateChampionshipAbb(championship) {
		log.WithField("championship", championship).Warning("championship not matching with regex")
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

// Check if the championship abbreviation is valid
func validateChampionshipAbb(c string) bool {
	CHAMPIONSHIP_REGEX := `(BHV|UF|OF|MF|OS|SW|AB|AV|OB)\\s(\\d{4}\\/\\d{2})`

	re := regexp.MustCompile(CHAMPIONSHIP_REGEX)

	return re.MatchString(c)
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
