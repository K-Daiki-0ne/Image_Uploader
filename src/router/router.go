package router

import (
	"github.com/gin-gonic/gin"
)

// Router : application router
func Router(app *gin.Engine) {
	app.GET("/file")
}
