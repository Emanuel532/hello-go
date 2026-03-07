package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Say hello to people", func(t *testing.T) {
		got := Hello("Mike", "")
		want := "Hello, Mike"
		assertCorrectMessage(t, got, want)
	})

	t.Run("greet the world if no name provided", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello to spanish people", func(t *testing.T) {
		got := Hello("Pedro", "Spanish")
		want := "Hola, Pedro"
		assertCorrectMessage(t, got, want)
	})

	t.Run("greet french people", func(t *testing.T) {
		got := Hello("Florian", "French")
		want := "Bonjour, Florian"
		assertCorrectMessage(t, got, want)
	})

	t.Run("greet romanians", func(t *testing.T) {
		got := Hello("Alexandru", "Romanian")
		want := "Salut, Alexandru"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
