package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	if app != nil {
		fmt.Println("Server listening ...OK")
		app.Run(":8080")
	} else {
		log.Panic("Server listening ...NO")
	}
}
