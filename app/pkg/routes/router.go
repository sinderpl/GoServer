package routes

import (
	"os"

	// sampleroutes "app/pkg/controller/sampleroute"
	sampleroutes "github.com/sinderpl/GoServer/app/pkg/controller/sampleroute"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func InitializeRouter() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route registering
	sampleroutes.Routes(e.Group(("")))

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
