package sampleroutes

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func Routes(router *echo.Group) {
	router.GET("/forms", getSample)
	router.POST("/forms", createSample)
	router.PATCH("/forms/:id", updateSample)
	router.DELETE("/forms/:id", deleteSample)
}

func getSample(c echo.Context) error {
	fmt.Println("Hello, World ! get sample")
	return c.JSON(http.StatusOK, "Hello, World Get !")
}

func createSample(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello, World Create!")
}

func updateSample(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello, World Set !")
}

func deleteSample(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello, World Delete !")
}
