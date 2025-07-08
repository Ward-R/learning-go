package main

import (
	"log"
	"testing"
)

func TestAdd(t *testing.T) {
	expected := 12
	result := add(7, 5)
	if result != 12 {
		log.Fatalf("error -  expected: %d, result: %d.", expected, result)
	}
}

func TestSubtract(t *testing.T) {
	expected := 2
	result := subtract(7, 5)
	if result != 2 {
		log.Fatalf("error -  expected: %d, result: %d.", expected, result)
	}
}
