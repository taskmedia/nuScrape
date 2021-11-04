package rest

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/taskmedia/nuScrape/pkg/parser"
	"github.com/taskmedia/nuScrape/pkg/scrape"
	"github.com/taskmedia/nuScrape/pkg/sport"
	"github.com/taskmedia/nuScrape/pkg/sport/group"
	"github.com/taskmedia/nuScrape/pkg/sport/season"
)

// AddRouterGesamtspielplan will add a GET request to receive a gesamtspielplan for the added engine.
// It validates the parameters and uses a website scrapper to generate the gesamtspielplan.
// The REST endpoint will requre three parameters: season, championship and group.
func AddRouterGesamtspielplan(engine *gin.Engine) {
	engine.GET("/rest/v1/gesamtspielplan/:season/:championship/:group", func(c *gin.Context) {
		// get rest parameters
		param_season := c.Param("season")
		championship := c.Param("championship")
		param_group := c.Param("group")

		// validate parameters
		season, err := season.New(param_season)
		if err != nil {
			log.WithField("season", param_season).Warning(err)
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		if !sport.ValidateChampionshipAbb(championship) {
			msg := "The given championship abbreviation is invalid"
			log.WithField("championship", championship).Warning(msg)
			c.String(http.StatusBadRequest, msg)
			return
		}

		group, err := group.New(param_group)
		if err != nil {
			log.WithField("group", param_group).Warning(err)
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		// get data from scrapper
		html_scrape, err := scrape.ScrapeGesamtspielplan(season, championship, group)
		if err != nil {
			log.WithFields(
				log.Fields{
					"season":       season,
					"championship": championship,
					"group":        group,
				},
			).Warning(err)
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		// parse website content to Matches
		matches, err := parser.ParseGesamtspielplan(html_scrape)
		if err != nil {
			log.Warning("parsing of gesamtspielplan failed")
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		// test
		fmt.Println(matches)

		// return data
		c.String(http.StatusOK, "not yet implemented")
	})
}
