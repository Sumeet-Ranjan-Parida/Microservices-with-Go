package main

import "testing"

func test(t *testing.T) {
	if calculate(2) != 4 {
		t.Error("Expected 2 + 2 to equal 4")
	}
}
