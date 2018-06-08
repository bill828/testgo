package testgo

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Log("Hello, world.")
	c := sum(1, 1)
	if c != 2 {
		t.Fatalf("2 is expected for sum(1, 1), but got %d.", c)
	}
}

