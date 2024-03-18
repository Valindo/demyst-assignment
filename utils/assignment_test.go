package utils

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGenerateUrl(t *testing.T) {
	baseUrl := "https://example.com"
	iterations := 4
	expectedResult := []string{
		"https://example.com/2",
		"https://example.com/4",
		"https://example.com/6",
		"https://example.com/8",
	}
	result := generateURL(baseUrl, iterations)
	for index, url := range result {
		if expectedResult[index] != url {
			t.Errorf("Generated URL does not match expected values")
		}
	}
}

func TestFetchTodosSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/todos/2" {
			t.Errorf("Expected request from /todos/2, but got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"userId": 1,
			"id": 1,
			"title": "delectus aut autem",
			"completed": false
		  }`))
	}))

	defer server.Close()
	baseUrl := fmt.Sprintf("%s/todos", server.URL)
	urls := generateURL(baseUrl, 1)
	todos, err := FetchTodos(urls)
	if len(todos) != 1 {
		t.Errorf("Expected a single todo but got none")
	}
	if err != nil {
		t.Errorf("Expected no errors, but got one")
	}
}

func TestFetchTodosBadResponse(t *testing.T) {
	urls := generateURL("htp:badurl", 1)
	errorMessage := `Failed for URL: htp:badurl/2 Reason: Get "htp:badurl/2": unsupported protocol scheme "htp"`
	todos, err := FetchTodos(urls)
	if len(todos) != 0 {
		t.Errorf("Expected no todos")
	}
	if err == nil {
		t.Errorf("Expected errors, but got none")
	}
	if err[0].String() != errorMessage {
		t.Errorf("Invalid Error message")
	}
}

func TestFetchTodosMalformedJson(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		// Sending back an empty response
		w.Write([]byte(``))
	}))

	defer server.Close()
	baseUrl := fmt.Sprintf("%s/todos", server.URL)
	errorMessage := fmt.Sprintf(`Failed for URL: %s/todos/2 Reason: unexpected end of JSON input`, server.URL)
	urls := generateURL(baseUrl, 1)
	todos, err := FetchTodos(urls)
	if len(todos) != 0 {
		t.Errorf("No todos expected")
	}
	if err == nil {
		t.Errorf("Expected errors, but got none")
	}
	if err[0].String() != errorMessage {
		t.Errorf("Invalid error message")
	}
}
