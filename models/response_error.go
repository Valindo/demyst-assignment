package models

import "fmt"

type ResponseError struct {
	Url          string
	ErrorMessage error
}

func (err ResponseError) String() string {
	return fmt.Sprintf("Failed for URL: %s Reason: %s", err.Url, err.ErrorMessage)
}
