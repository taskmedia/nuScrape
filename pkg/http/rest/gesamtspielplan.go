package rest

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/taskmedia/nuScrape/pkg/parser"
	"github.com/taskmedia/nuScrape/pkg/scrape"
	"github.com/taskmedia/nuScrape/pkg/sport"
	"github.com/taskmedia/nuScrape/pkg/sport/championship"
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
		param_championship := c.Param("championship")
		param_group := c.Param("group")

		// validate parameters
		s, err := season.New(param_season)
		if err != nil {
			log.WithField("season", param_season).Warning(err)
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		ch, err := championship.ParseAbbreviation(param_championship)
		if err != nil {
			log.WithField("championship", param_championship).Warning(err.Error())
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		g, err := group.New(param_group)
		if err != nil {
			log.WithField("group", param_group).Warning(err)
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		// get data from scrapper
		htmlInfo_scrape, htmlTable_scrape, err := scrape.ScrapeGesamtspielplan(s, ch, g)
		if err != nil {
			log.WithFields(log.Fields{
				"season":       s,
				"championship": ch,
				"group":        g,
				"scrape_info":  htmlInfo_scrape,
				"scrape_table": htmlTable_scrape,
			},
			).Warning(err)
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		ac, cl, r, err := parser.ParseGesamtspielplanInfo(htmlInfo_scrape)
		if err != nil {
			err_msg := "parsing of gesamtspielplan info failed"
			log.WithFields(log.Fields{
				"html_scrape": htmlInfo_scrape.DOM.Text(),
				"ageCategory": ac,
				"class":       cl,
				"relay":       r,
				"err":         err,
			}).Warning(err_msg)
			c.String(http.StatusInternalServerError, err_msg)
			return
		}

		// parse website content to Matches
		matches, err := parser.ParseGesamtspielplanTable(htmlTable_scrape)
		if err != nil {
			err_msg := "parsing of matches failed"
			log.WithFields(log.Fields{
				"html_scrape": htmlTable_scrape.DOM.Text(),
				"matches":     matches,
				"error":       err,
			}).Warning(err_msg)
			c.String(http.StatusInternalServerError, err_msg)
			return
		}

		gsp := sport.Gesamtspielplan{
			Season:       s,
			Championship: ch,
			Group:        g,
			Matches:      matches,
		}

		// add Gesamtspielplan info to gsp
		gsp.AgeCategory = ac
		gsp.Class = cl
		gsp.Relay = r

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
