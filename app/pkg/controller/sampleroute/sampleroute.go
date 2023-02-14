package sampleroutes

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func Routes(router *echo.Group) {
	router.GET("/sample", getSample)
	router.POST("/sample", createSample)
	router.PATCH("/sample/:id", updateSample)
	router.DELETE("/sample/:id", deleteSample)
}

func getSample(c echo.Context) error {
	requestBody := &RequestBody{}

	err := c.Bind(requestBody)
	if err != nil {
		return c.JSON(http.StatusNotFound, ErrorResponse{
			Title:   "invalid request body",
			Message: "body could not be parsed",
		})
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
		return c.JSON(http.StatusNotFound, ErrorResponse{
			Title:   "invalid request body",
			Message: "body could not be parsed",
		})
	}

	id := rand.Int63()
	newObject := Sample{
		ID:   id,
		Name: requestBody.SampleName,
	}

	body := &StatusResponse{
		Message: fmt.Sprintf("Sample Create ! User name: %s", requestBody.UserName),
		Sample:  newObject,
	}

	return c.JSON(http.StatusOK, body)
}

func updateSample(c echo.Context) error {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusNotFound, ErrorResponse{
			Title:   "invalid parameter",
			Message: "id must be numerical",
		})
	}

	requestBody := &RequestBody{}

	err = c.Bind(requestBody)
	if err != nil {
		return c.JSON(http.StatusNotFound, ErrorResponse{
			Title:   "invalid request body",
			Message: "body could not be parsed",
		})
	}

	sample := Sample{
		ID:   id,
		Name: requestBody.SampleName,
	}

	body := &StatusResponse{
		Message: "Sample Set !",
		Sample:  sample,
	}

	return c.JSON(http.StatusOK, body)
}

func deleteSample(c echo.Context) error {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusNotFound, ErrorResponse{
			Title:   "invalid parameter",
			Message: "id must be numerical",
		})
	}

	body := &StatusResponse{
		Message: fmt.Sprintf("Sample Deleted ID: %d", id),
	}

	return c.JSON(http.StatusOK, body)
}
