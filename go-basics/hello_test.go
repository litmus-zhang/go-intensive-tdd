package main

import "testing"

func TestHelloo(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"

		ascertCorrectMessage(t, got, want)
	})
	t.Run("Say 'Hello, World', when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		ascertCorrectMessage(t, got, want)
	})
}

func ascertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
