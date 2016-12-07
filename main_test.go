package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ResponseCode(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8080", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	HelloWorld(res, req)

	exp := 200
	act := res.Code
	assert.Equal(t, exp, act, "Response code should be 200")
}

func Test_HelloWorld(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8080", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	HelloWorld(res, req)

	exp := "Hello World"
	act := res.Body.String()
	assert.Equal(t, exp, act, "Response should be Hello World")
}

func Test_JsonHelloWorld(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8080/json", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	JsonHelloWorld(res, req)

	assert.Equal(t, 200, res.Code, "Response code should be 200")
	assert.JSONEq(t, `{ "success": true, "message": "Hello World" }`, res.Body.String(), "Response body should be JSON")

	var sr SimpleResponse
	json.NewDecoder(res.Body).Decode(&sr)
	assert.True(t, sr.Success, "Success value should be true")
	assert.Equal(t, "Hello World", sr.Message, "Message should be Hello World")
}
