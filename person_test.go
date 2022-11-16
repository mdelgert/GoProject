//https://ieftimov.com/posts/testing-in-go-go-test/#:~:text=To%20run%20your%20tests%20in,prints%20the%20complete%20test%20output.

package main

import (
	"testing"
)

func TestNewPersonPositiveAge(t *testing.T) {
	_, err := NewPerson(1)
	if err != nil {
		t.Errorf("Expected person, received %v", err)
	}
}

func TestNewPersonNegativeAge(t *testing.T) {
	p, err := NewPerson(-1)
	if err == nil {
		t.Errorf("Expected error, received %v", p)
	}
}
