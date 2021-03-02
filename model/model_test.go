package model

import "testing"

func TestModel(t *testing.T) {
	assertMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("Hellow World model", func(t *testing.T) {
		got := HelloWorld()
		want := "Hello, world!"

		assertMessage(t, got, want)
	})

}
