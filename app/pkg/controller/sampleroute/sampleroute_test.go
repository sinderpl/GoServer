package sampleroutes

import (
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetSample(t *testing.T) {
	t.Run("Test simple get", func(t *testing.T) {
		//request, _ := http.NewRequest(http.MethodGet, "/sample", nil)
		//response := httptest.NewRecorder()

		//routes.InitializeRouter()

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/sample/")
		c.SetParamNames("userName")
		c.SetParamValues("John")

		// Assertions
		//if assert.NoError(t, h.getUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		//assert.Equal(t, userJSON, rec.Body.String())
		//}
	})
}

func TestPostSample(t *testing.T) {

}

func TestPatchSample(t *testing.T) {

}

func TestDeleteSample(t *testing.T) {

}
