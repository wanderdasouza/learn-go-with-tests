package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper() // Indica para a bateria de testes que esta é uma função auxiliar

	if got != want {
		t.Errorf("got %q want %q", got, want) 
			//Errorf printa a string dada e falha o teste. 
			// O 'f' no final permite interpolação na string 
	}
}