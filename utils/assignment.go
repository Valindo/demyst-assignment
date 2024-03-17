package utils

import "fmt"

type ResponseBody struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func (item *ResponseBody) Details() {
	completed := " "
	if item.Completed {
		completed = "X"
	}
	fmt.Printf("[%s] %s\n", completed, item.Title)
}

func GenerateURL(prefix string, iterations int) []string {
	var urls []string
	even := 2
	for counter := 0; counter < iterations; counter++ {
		url := fmt.Sprintf("%s/%d", prefix, even)
		even += 2
		urls = append(urls, url)
	}
	return urls
}
