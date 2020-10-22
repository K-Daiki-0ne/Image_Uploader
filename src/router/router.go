package router

import (
	"0ne/src/controller"

	"github.com/gin-gonic/gin"
)

// Router : application router
func Router(app *gin.Engine) {
	app.POST("/file", controller.Upload)
}
