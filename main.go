package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/valindo/demyst-assignment/utils"
)

func main() {
	baseUrl := "https://jsonplaceholder.typicode.com/todos"
	iterations := 20
	urls := utils.GenerateURL(baseUrl, iterations)
	// var todos []utils.ResponseBody
	for _, url := range urls {
		var data utils.ResponseBody
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		body, readErr := io.ReadAll(res.Body)
		if readErr != nil {
			log.Fatal(err)
		}
		json.Unmarshal(body, &data)
		data.Details()
		// todos = append(todos, data)
		res.Body.Close()
	}
}
