package main

import (
	"0ne/src/router"

	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	if app != nil {
		router.Router(app)
		fmt.Println("Server listening ...OK")
		app.Run(":8080")
	} else {
		log.Panic("Server listening ...NO")
	}
}
