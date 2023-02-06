package sampleroutes

import (
	"fmt"
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
	UserName string `json:"userName"`
}

type StatusResponse struct {
	Message string `json:"message"`
}

func getSample(c echo.Context) error {
	requestBody := &RequestBody{}

	err := c.Bind(requestBody)
	if err != nil {
		return err
	}

	body := &StatusResponse{
		Message: fmt.Sprintf("Sample Get ! User name : %s", requestBody.UserName),
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
