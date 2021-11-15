package rest

import (
	"encoding/json"
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
func addRouterGesamtspielplan(engine *gin.Engine) {
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
		_, htmlTable_scrape, err := scrape.ScrapeGesamtspielplan(season, championship, group)
		if err != nil {
			log.WithFields(log.Fields{
				"season":       season,
				"championship": championship,
				"group":        group,
			},
			).Warning(err)
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		// parse website content to Matches
		matches, err := parser.ParseGesamtspielplanTable(htmlTable_scrape)
		if err != nil {
			err_msg := "parsing of matches failed"
			log.WithFields(log.Fields{
				"html_scrape": htmlTable_scrape,
				"matches":     matches,
				"error":       err,
			}).Warning(err_msg)
			c.String(http.StatusInternalServerError, err_msg)
			return
		}

		gsp := sport.Gesamtspielplan{
			Season:       season,
			Championship: championship,
			Group:        group,
			Matches:      matches,
		}

		// return matches as JSON
		c.Writer.Header().Set("Content-Type", "application/json")
		wr, err := json.Marshal(gsp)
		if err != nil {
			err_msg := "could not parse Gesamtspielplan to JSON"
			log.WithFields(log.Fields{
				"gesamtspielplan": gsp,
				"error":           err,
			}).Warning(err_msg)
			c.String(http.StatusInternalServerError, err_msg)
			return
		}
		c.Writer.Write(wr)
	})
}
