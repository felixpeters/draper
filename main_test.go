package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
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
	if exp != act {
		t.Fatalf("Expected response code to be %d, but got %d", exp, act)
	}
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
	if exp != act {
		t.Fatalf("Expected %s, but got %s", exp, act)
	}
}
