package rest

import (
	"net/http"
	"testing"
)

// Test AddRouterGesamtspielplan HTTP endpoints
// Testing different responses of REST endpoints
func TestAddRouterGesamtspielplan(t *testing.T) {
	// 200 - OK
	checkEndpointGetStatuscode(t, "/rest/v1/gesamtspielplan/2021_22/AV/281103", http.StatusOK)

	// 400 - Bad Request
	checkEndpointGetStatuscode(t, "/rest/v1/gesamtspielplan/2021_TS/AV/281103", http.StatusBadRequest)
	checkEndpointGetStatuscode(t, "/rest/v1/gesamtspielplan/2021_23/AV/281103", http.StatusBadRequest)
	checkEndpointGetStatuscode(t, "/rest/v1/gesamtspielplan/2021_22/YZ/281103", http.StatusBadRequest)
	checkEndpointGetStatuscode(t, "/rest/v1/gesamtspielplan/2021_22/YZ/noint", http.StatusBadRequest)

	// 404 - Not Found
	checkEndpointGetStatuscode(t, "/notexisting", http.StatusNotFound)
}
