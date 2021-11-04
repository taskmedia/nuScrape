package rest

import "github.com/gin-gonic/gin"

// SetupRouter creates the Gin.Engine with its routers
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// add routes to Engine
	addRouterGesamtspielplan(r)

	return r
}
