package main

import (
	"fmt"
	"github.com/labstack/echo"
	"gorm-many2many/database"
	"gorm-many2many/handlers"
)

func main() {
	database.Init()
	e := echo.New()

	// health check
	e.GET("/ping", handlers.Ping)
	// get place by id
	// example: /places?uid=1
	e.GET("/places", handlers.GetPlace)

	echoStartError := e.Start(":1339")
	if echoStartError != nil {
		fmt.Printf("Unable to start http listener")
	}
}