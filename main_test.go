package main

import (
	"net/http/httptest"
	"testing"
)

func TestGetUser(t *testing.T) {
	request := httptest.NewRequest("GET", "/users/6160beda11e038afddbfc335", nil)
	response := httptest.NewRecorder()
	UserById(response, request)

	resp := response.Result()
	if 200 != resp.StatusCode {
		t.Fatal("Test Failed")
	}
}

func TestGetPost(t *testing.T) {
	request := httptest.NewRequest("GET", "/posts/6160bf9c11e038afddbfc336", nil)
	response := httptest.NewRecorder()
	UserById(response, request)

	resp := response.Result()
	if 200 != resp.StatusCode {
		t.Fatal("Test Failed")
	}
}

func TestGetPostbyUserId(t *testing.T) {
	request := httptest.NewRequest("GET", "/posts/users/6160bf9c11e038afddbfc336", nil)
	response := httptest.NewRecorder()
	UserById(response, request)

	resp := response.Result()
	if 200 != resp.StatusCode {
		t.Fatal("Test Failed")
	}
}

func TestCreateUser(t *testing.T) {
	request := httptest.NewRequest("GET", "/posts/users/6160bf9c11e038afddbfc336", nil)
	response := httptest.NewRecorder()
	UserById(response, request)

	resp := response.Result()
	if 200 != resp.StatusCode {
		t.Fatal("Test Failed")
	}
}

func TestCreatePost(t *testing.T) {
	request := httptest.NewRequest("GET", "/posts/users/6160bf9c11e038afddbfc336", nil)
	response := httptest.NewRecorder()
	UserById(response, request)

	resp := response.Result()
	if 200 != resp.StatusCode {
		t.Fatal("Test Failed")
	}
}
