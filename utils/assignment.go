package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/valindo/demyst-assignment/models"
)

func generateURL(prefix string, iterations int) []string {
	var urls []string
	even := 2
	for counter := 0; counter < iterations; counter++ {
		url := fmt.Sprintf("%s/%d", prefix, even)
		even += 2
		urls = append(urls, url)
	}
	return urls
}

func FetchTodos(urls []string) ([]models.Todo, []models.ResponseError) {
	var todos []models.Todo
	var errors []models.ResponseError
	for _, url := range urls {
		var responseBody models.Todo
		response, err := http.Get(url)
		if err != nil {
			errors = append(errors, models.ResponseError{
				ErrorMessage: err,
				Url:          url,
			})
			continue
		}
		body, err := io.ReadAll(response.Body)
		defer response.Body.Close()
		if err != nil {
			errors = append(errors, models.ResponseError{
				ErrorMessage: err,
				Url:          url,
			})
			continue
		}
		err = json.Unmarshal(body, &responseBody)
		if err != nil {
			errors = append(errors, models.ResponseError{
				ErrorMessage: err,
				Url:          url,
			})
			continue
		}
		todos = append(todos, responseBody)

	}
	return todos, errors
}

func Display(baseUrl string, iterations int) {
	urls := generateURL(baseUrl, iterations)
	todos, err := FetchTodos(urls)
	for _, todo := range todos {
		fmt.Println(todo)
	}
	if err != nil {
		fmt.Printf("\nFollowing errors have occured\n")
		for _, e := range err {
			fmt.Println(e)
		}
	}

}
