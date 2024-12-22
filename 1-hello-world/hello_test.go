package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people in English", func(t *testing.T) {
		got := Hello("Andrii", "English")
		want := "Hello, Andrii"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("say 'Hello, World' when an empty string is supplied in English", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Carlos", "Spanish")
		want := "Hola, Carlos"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("default to 'World' in Spanish when name is empty", func(t *testing.T) {
		got := Hello("", "Spanish")
		want := "Hola, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Marie", "French")
		want := "Bonjour, Marie"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("default to 'World' in French when name is empty", func(t *testing.T) {
		got := Hello("", "French")
		want := "Bonjour, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}