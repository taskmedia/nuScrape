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
		{"season": "2021_22", "championship": "AV", "group": "281103"},
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
