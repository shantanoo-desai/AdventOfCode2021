package main

import "testing"

func TestSteering(t *testing.T) {

	// Start with some defined forward motion
	steering := &Steering{}

	steering.Navigate("forward", 5)
	steering.Navigate("down", 5)
	steering.Navigate("forward", 8)
	steering.Navigate("up", 3)
	steering.Navigate("down", 8)

	steering.Navigate("forward", 2)

	got := steering.Product()
	expected := 900

	if got != expected {
		t.Errorf("expected %q, got: %v", expected, got)
	}
}
