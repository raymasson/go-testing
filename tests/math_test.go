package tests

import "testing"
import "go-testing/functions"

func TestCanAddNumbers(t *testing.T) {
	result := functions.Add(1, 2)

	if result != 3 {
		t.Log("Failed to add one and two")
		t.Fail()
	}

	result = functions.Add(1, 2, 3, 4)

	if result != 10 {
		t.Error("Failed to add more than two numbers")
	}
}
