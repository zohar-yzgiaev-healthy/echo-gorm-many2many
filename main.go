package main

import (
	"fmt"
	"github.com/labstack/echo"
	"gorm-many2many/handlers"
)

func main() {
	e := echo.New()
	e.GET("/ping", handlers.Ping)

	echoStartError := e.Start(":1339")
	if echoStartError != nil {
		fmt.Printf("Unable to start http listener")
	}
}