package sampleroutes

import (
	"net/http"

	"github.com/labstack/echo"
)

func Routes(router *echo.Group) {
	router.GET("/sample", getSample)
	router.POST("/sample", createSample)
	router.PATCH("/sample/:id", updateSample)
	router.DELETE("/sample/:id", deleteSample)
}

func getSample(c echo.Context) error {
	return c.JSON(http.StatusOK, "Sample Get !")
}

func createSample(c echo.Context) error {
	return c.JSON(http.StatusOK, "Sample Create!")
}

func updateSample(c echo.Context) error {
	return c.JSON(http.StatusOK, "Sample Set !")
}

func deleteSample(c echo.Context) error {
	return c.JSON(http.StatusOK, "Sample Delete !")
}
