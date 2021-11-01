package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
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
	year, err := getYear(c.Param("year"))
	if err != nil {
		log.WithField("year", year).Warning(err)
		c.String(http.StatusBadRequest, err.Error())
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

	url := generateGesamtspielplanUrl(year, championship, group)

	log.WithField("url", url).Info("Scraping URL")

	c.String(http.StatusOK, "not yet implemented")
}

func generateGesamtspielplanUrl(year int, championship string, group int) url.URL {
	u := url.URL{}
	u.Scheme = "https"
	u.Host = "bhv-handball.liga.nu"
	u.Path = "/cgi-bin/WebObjects/nuLigaHBDE.woa/wa/groupPage"

	query := url.Values{}
	query.Add("displayTyp", "gesamt")
	query.Add("displayDetail", "meetings")
	query.Add("championship", championship+" "+getYearFormatted(year))
	query.Add("group", strconv.Itoa(group))
	u.RawQuery = query.Encode()

	return u
}

// Get and validate year
// Checks if the year pattern is matching 6 digits and is correct
// YYYYYY - e.g. 202122 for 2021/22
func getYear(yearString string) (int, error) {
	if !regexp.MustCompile(`\d{6}`).MatchString(yearString) {
		return -1, errors.New("year pattern not valid (expected 6 digits)")
	}

	y, err := strconv.Atoi(yearString)
	if err != nil {
		return -1, errors.New("year not an integer")
	}

	// check if years match to each other
	// valid: e.g. 202122, 199900
	// not valid: e.g. 202123
	hr, _ := strconv.Atoi(yearString[2:4])
	rr, _ := strconv.Atoi(yearString[4:])

	if !(((hr + 1) == rr) || (hr == 99 && rr == 0)) {
		return -1, errors.New("year patter not valid (years must follow each other)")
	}

	return y, nil
}

// Return the year to nuLiga URL
func getYearFormatted(y int) string {
	yearString := strconv.Itoa(y)
	return fmt.Sprintf("%s/%s", yearString[:4], yearString[4:])
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
