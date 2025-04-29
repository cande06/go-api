package tests

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"go-api/api"
	"go-api/internal/user"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIntegrationCreateAndGet(t *testing.T) {
	gin.SetMode(gin.TestMode)
	app := gin.Default()
	api.InitRoutes(app)

	var createdUser user.User
	var createdSaleID string

	t.Run("create user", func(t *testing.T) {
		reqUser, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBufferString(`{
			"name":"Ayrton",
			"address":"Pringles",
			"nickname":"Chiche"
		}`))
		reqUser.Header.Set("Content-Type", "application/json")

		resUser := fakeRequest(app, reqUser)

		require.Equal(t, http.StatusCreated, resUser.Code)
		require.NoError(t, json.Unmarshal(resUser.Body.Bytes(), &createdUser))
		require.NotEmpty(t, createdUser.ID)
	})

	t.Run("patch user", func(t *testing.T) {
		updatePayload := map[string]interface{}{
			"name":     "Juan",
			"address":  "San Martin 134",
			"nickname": "Juan Doe",
		}
		payloadBytes, _ := json.Marshal(updatePayload)

		reqUser, _ := http.NewRequest(http.MethodPatch, "/users/"+createdUser.ID, bytes.NewBuffer(payloadBytes))
		reqUser.Header.Set("Content-Type", "application/json")

		resUser := fakeRequest(app, reqUser)

		require.Equal(t, http.StatusOK, resUser.Code)

		var patchUser user.User
		require.NoError(t, json.Unmarshal(resUser.Body.Bytes(), &patchUser))
		require.Equal(t, "Juan", patchUser.Name)
	})

	t.Run("create sale", func(t *testing.T) {
		salePayload := map[string]interface{}{
			"user_id": createdUser.ID,
			"amount":  999.9,
		}
		payloadBytes, _ := json.Marshal(salePayload)

		reqSale, _ := http.NewRequest(http.MethodPost, "/sales", bytes.NewBuffer(payloadBytes))
		reqSale.Header.Set("Content-Type", "application/json")

		resSale := fakeRequest(app, reqSale)
		require.Equal(t, http.StatusCreated, resSale.Code)

		var createdSale map[string]interface{}
		require.NoError(t, json.Unmarshal(resSale.Body.Bytes(), &createdSale))
		createdSaleID = createdSale["id"].(string)
		require.NotEmpty(t, createdSaleID)
	})

	t.Run("get sale", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/sales?user_id="+createdUser.ID, nil)
		res := fakeRequest(app, req)

		require.Equal(t, http.StatusOK, res.Code)

		var fetchedSale map[string]interface{}
		require.NoError(t, json.Unmarshal(res.Body.Bytes(), &fetchedSale))

		var response map[string]interface{}
		require.NoError(t, json.Unmarshal(res.Body.Bytes(), &response))

		results := response["results"].([]interface{})
		require.NotEmpty(t, results)

		firstSale := results[0].(map[string]interface{})
		require.Equal(t, createdUser.ID, firstSale["user_id"])
	})

	// t.Run("delete user", func(t *testing.T) {
	// 	req, _ := http.NewRequest(http.MethodDelete, "/users/"+createdUser.ID, nil)
	// 	res := fakeRequest(app, req)
	// 	require.Equal(t, http.StatusOK, res.Code)
	// })

	t.Run("ping", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/ping", nil)
		res := fakeRequest(app, req)
		require.Equal(t, http.StatusOK, res.Code)
		require.Contains(t, res.Body.String(), "pong")
	})
}

func fakeRequest(e *gin.Engine, r *http.Request) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	e.ServeHTTP(w, r)

	return w
}
