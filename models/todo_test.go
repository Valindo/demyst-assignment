package models

import (
	"testing"
)

func TestTodoStringerMethod(t *testing.T) {
	todo := Todo{
		Completed: true,
		Title:     "test completed",
		UserId:    1,
		Id:        1,
	}
	expected := "[X] test completed"
	if todo.String() != expected {
		t.Error("Invalid stringer output")
	}
}
