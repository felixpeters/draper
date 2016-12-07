package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/stretchr/testify/assert"
)

func checkError(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
}

func Test_ResponseCode(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8080", nil)
	checkError(err, t)

	res := httptest.NewRecorder()
	HelloWorld(res, req)

	exp := 200
	act := res.Code
	assert.Equal(t, exp, act, "Response code should be 200")
}

func Test_HelloWorld(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8080", nil)
	checkError(err, t)

	res := httptest.NewRecorder()
	HelloWorld(res, req)

	exp := "Hello World"
	act := res.Body.String()
	assert.Equal(t, exp, act, "Response should be Hello World")
}

func Test_JsonHelloWorld(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8080/json", nil)
	checkError(err, t)

	res := httptest.NewRecorder()
	JsonHelloWorld(res, req)

	assert.Equal(t, 200, res.Code, "Response code should be 200")
	assert.JSONEq(t, `{ "success": true, "message": "Hello World" }`, res.Body.String(), "Response body should be JSON")

	body, err := ioutil.ReadAll(res.Body)
	checkError(err, t)
	js, err := simplejson.NewJson(body)
	checkError(err, t)

	assert.True(t, js.Get("success").MustBool(), "Success value should be true")
	assert.Equal(t, "Hello World", js.Get("message").MustString(), "Message should be Hello World")
}
