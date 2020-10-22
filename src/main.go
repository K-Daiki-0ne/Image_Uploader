package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	if app != nil {
		fmt.Println("Server start ...OK")
		app.Run(":8080")
	}
}
