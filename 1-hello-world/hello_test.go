package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Andrii")
	want := "Hello, Andrii"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
