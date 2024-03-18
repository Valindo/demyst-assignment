package models

import (
	"errors"
	"testing"
)

func TestResponseErrorStringerMethod(t *testing.T) {
	err := ResponseError{
		Url:          "http://test.com",
		ErrorMessage: errors.New("testing"),
	}
	expected := "Failed for URL: http://test.com Reason: testing"
	if err.String() != expected {
		t.Error("Invalid stringer output")
	}
}
