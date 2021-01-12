package test

import (
	"api/config"
	"api/database"
	"api/router"

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

	body := strings.NewReader(`{"name": "name"}`)

	req := httptest.NewRequest("POST", "/user/create", body)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	rec := httptest.NewRecorder()

	router.Engin.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}
