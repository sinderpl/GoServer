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

type RequestBody struct {
}

type StatusResponse struct {
	Message string
}
ß
func getSample(c echo.Context) error {
	_ = &RequestBody{}

	body := &StatusResponse{
		Message: "Sample Get !",
	}

	return c.JSON(http.StatusOK, body)
}

func createSample(c echo.Context) error {

	body := &StatusResponse{
		Message: "Sample Create !",
	}

	return c.JSON(http.StatusOK, body)
}

func updateSample(c echo.Context) error {

	body := &StatusResponse{
		Message: "Sample Set !",
	}

	return c.JSON(http.StatusOK, body)
}

func deleteSample(c echo.Context) error {

	body := &StatusResponse{
		Message: "Sample Delete !",
	}

	return c.JSON(http.StatusOK, body)
}
