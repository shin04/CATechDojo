package test

import (
	"api/config"
	"api/controller"
	"api/database"
	"api/router"
	"encoding/json"

	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type reqBody map[string]string

func TestMain(m *testing.M) {
	Setup()
	m.Run()
}

func Setup() {
	config := &config.Config{}
	config.Init()

	database.Init(config)
}

func createRequest(body reqBody, method string, endpoint string, token string) *http.Request {
	bodyByte, _ := json.Marshal(body)
	reqBody := strings.NewReader(string(bodyByte))
	req := httptest.NewRequest(method, endpoint, reqBody)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("x-token", token)

	return req
}

func TestResponse(t *testing.T) {
	router := &router.Router{}
	router.Init()

	var req *http.Request
	var rec *httptest.ResponseRecorder

	// testing /user/create
	createReqBody := reqBody{"name": "name"}
	req = createRequest(createReqBody, "POST", "/user/create", "")
	rec = httptest.NewRecorder()
	router.Engin.ServeHTTP(rec, req)

	var createResBody controller.UserCreateResponse
	err := json.Unmarshal([]byte(rec.Body.String()), &createResBody)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, nil, err)

	// tesiting /user/get
	tokenString := createResBody.Token
	req = createRequest(reqBody{}, "GET", "/user/get", tokenString)
	rec = httptest.NewRecorder()
	router.Engin.ServeHTTP(rec, req)

	var getResBody controller.UserGetResponse
	_ = json.Unmarshal([]byte(rec.Body.String()), &getResBody)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "name", getResBody.Name)

	// testing /user/update
	updateReqBody := reqBody{"name": "rename"}
	req = createRequest(updateReqBody, "PUT", "/user/update", tokenString)
	rec = httptest.NewRecorder()
	router.Engin.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}
