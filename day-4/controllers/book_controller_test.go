package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"alterrafour/controllers"
	"alterrafour/lib/inmemdb"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetAllBookController(t *testing.T) {
	testCases := []struct {
		name                string
		path                string
		expectStatus        int
		expectBodyStartWith string
	}{
		{
			name:         "test all books",
			path:         "/books",
			expectStatus: http.StatusOK,
			expectBodyStartWith: `{
				"data": [
					{
						"id": 1,
						"title": "Things Fall Apart",
						"author": "Chinua Achebe",
						"year": 1958
					},
					{
						"id": 2,
						"title": "Fairy tales",
						"author": "Hans Christian Andersen",
						"year": 1836
					},
					{
						"id": 3,
						"title": "The Divine Comedy",
						"author": "Dante Alighieri",
						"year": 1315
					}
				],
				"status": "success"
			}`,
		},
	}

	con := controllers.Controller{
		BookDB: *inmemdb.New(),
	}
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	for _, testCase := range testCases {
		ctx.SetPath(testCase.path)
		if assert.NoError(t, con.GetAllBookController(ctx)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()
			assert.JSONEq(t, body, testCase.expectBodyStartWith)
		}
	}
}
