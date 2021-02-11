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

func TestMain(m *testing.M) {
	Setup()
	m.Run()
}

func Setup() {
	config := &config.Config{}
	config.Init()

	database.Init(config)
}

func TestResponse(t *testing.T) {
	router := &router.Router{}
	router.Init()

	// testing /user/create
	req_body := strings.NewReader(`{"name": "name"}`)
	req := httptest.NewRequest("POST", "/user/create", req_body)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	rec := httptest.NewRecorder()
	router.Engin.ServeHTTP(rec, req)

	var create_res_body controller.UserCreateResponse
	err := json.Unmarshal([]byte(rec.Body.String()), &create_res_body)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, nil, err)

	// tesiting /user/get
	token_string := create_res_body.Token
	req_body = strings.NewReader("")
	req = httptest.NewRequest("GET", "/user/get", req_body)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("x-token", token_string)
	rec = httptest.NewRecorder()
	router.Engin.ServeHTTP(rec, req)

	var get_res_body controller.UserGetResponse
	_ = json.Unmarshal([]byte(rec.Body.String()), &get_res_body)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "name", get_res_body.Name)

	// testing /user/update
	req_body = strings.NewReader(`{"name": "hogehoge"}`)
	req = httptest.NewRequest("PUT", "/user/update", req_body)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("x-token", token_string)
	rec = httptest.NewRecorder()
	router.Engin.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}
