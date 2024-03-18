package models

import "fmt"

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func (item Todo) String() string {
	completed := " "
	if item.Completed {
		completed = "X"
	}
	return fmt.Sprintf("[%s] %s", completed, item.Title)
}
