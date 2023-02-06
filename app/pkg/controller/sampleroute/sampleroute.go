package sampleroutes

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/labstack/echo"
)

func Routes(router *echo.Group) {
	router.GET("/sample", getSample)
	router.POST("/sample", createSample)
	router.PATCH("/sample/:id", updateSample)
	router.DELETE("/sample/:id", deleteSample)
}

type Sample struct {
	ID   int64  `json:"ID"`
	Name string `json:"name"`
}

type RequestBody struct {
	UserName string `json:"userName,omitempty"`
}

type StatusResponse struct {
	Message string `json:"message"`
	Sample  Sample `json:"sample"`
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
	requestBody := &RequestBody{}

	err := c.Bind(requestBody)
	if err != nil {
		return err
	}

	id := rand.Int63()
	newObject := Sample{
		ID:   id,
		Name: fmt.Sprintf("New Object %d", id),
	}

	body := &StatusResponse{
		Message: fmt.Sprintf("Sample Create ! User name: %s", requestBody.UserName),
		Sample:  newObject,
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
