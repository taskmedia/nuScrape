package rest

import (
	"fmt"
	"net/http"
	"testing"
)

// Test AddRouterGesamtspielplan HTTP endpoints
// Testing different responses of REST endpoints
func TestAddRouterGesamtspielplan(t *testing.T) {
	// 200 - OK
	testGroups := []map[string]string{
		// BHV
		{"season": "2021_22", "championship": "BHV", "group": "273310"}, // Bayernliga Männer Staffel Nord-West
		{"season": "2021_22", "championship": "BHV", "group": "279618"}, // Bayernliga Männer Staffel Süd-Ost
		{"season": "2021_22", "championship": "BHV", "group": "273572"}, // Landesliga Männer Staffel Nord
		{"season": "2021_22", "championship": "BHV", "group": "273413"}, // Landesliga Männer Staffel Süd-West
		{"season": "2021_22", "championship": "BHV", "group": "273616"}, // Bayernliga Frauen Staffel Süd
		{"season": "2021_22", "championship": "BHV", "group": "279620"}, // Landesliga Frauen Staffel Süd
		{"season": "2021_22", "championship": "BHV", "group": "273595"}, // Bayernliga männliche A-Jugend Nordwest
		{"season": "2021_22", "championship": "BHV", "group": "273463"}, // ÜBOL mA-Jgd. Nordost
		{"season": "2021_22", "championship": "BHV", "group": "273486"}, // ÜBOL wB-Jgd. Mitte
		{"season": "2021_22", "championship": "BHV", "group": "273395"}, // ÜBOL wC-Jgd. Südwest 1
		{"season": "2021_22", "championship": "BHV", "group": "273461"}, // ÜBL mC-Jgd. Mitte 1

		// Unterfranken
		{"season": "2021_22", "championship": "UF", "group": "280903"}, // Bezirksoberliga Männer Staffel B
		{"season": "2021_22", "championship": "UF", "group": "273549"}, // Bezirksklasse Staffel Nord Männer
		{"season": "2021_22", "championship": "UF", "group": "273611"}, // Bezirksliga Frauen
		{"season": "2021_22", "championship": "UF", "group": "280434"}, // Bezirksliga Staffel Nord

		// Schwaben
		{"season": "2021_22", "championship": "SW", "group": "273332"}, // Bezirksoberliga Männer - Staffel A
		{"season": "2021_22", "championship": "SW", "group": "273525"}, // Bezirksliga Frauen

		// Alpenvorland
		{"season": "2021_22", "championship": "AV", "group": "281102"}, // Bezirksoberliga West Männer
		{"season": "2021_22", "championship": "AV", "group": "273490"}, // Bezirksklasse Männer Staffel Süd-West
		{"season": "2021_22", "championship": "AV", "group": "281103"}, // Bezirksklasse Männer Staffel Nord-West
		{"season": "2021_22", "championship": "AV", "group": "281105"}, // Bezirksoberliga Frauen Staffel West
		{"season": "2021_22", "championship": "AV", "group": "273516"}, // Bezirksliga Frauen Staffel Süd-West
		{"season": "2021_22", "championship": "AV", "group": "281107"}, // Bezirksliga Frauen Staffel Nord-West
	}

	for _, g := range testGroups {
		url := fmt.Sprintf("/rest/v1/gesamtspielplan/%s/%s/%s", g["season"], g["championship"], g["group"])
		checkEndpointGetStatuscode(t, url, http.StatusOK)
	}

	// 400 - Bad Request
	checkEndpointGetStatuscode(t, "/rest/v1/gesamtspielplan/2021_TS/AV/281103", http.StatusBadRequest)
	checkEndpointGetStatuscode(t, "/rest/v1/gesamtspielplan/2021_23/AV/281103", http.StatusBadRequest)
	checkEndpointGetStatuscode(t, "/rest/v1/gesamtspielplan/2021_22/YZ/281103", http.StatusBadRequest)
	checkEndpointGetStatuscode(t, "/rest/v1/gesamtspielplan/2021_22/YZ/noint", http.StatusBadRequest)

	// 404 - Not Found
	checkEndpointGetStatuscode(t, "/notexisting", http.StatusNotFound)
}
