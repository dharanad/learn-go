package helloworld

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to people in spanish", func(t *testing.T) {
		got := Hello("Dharan", Spanish)
		want := "Hola, Dharan"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say world when no people in french", func(t *testing.T) {
		got := Hello("", French)
		want := "Bonjour, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say world when no people in english", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got string, want string) {
	// t.Helper() is needed to tell the test suite that this method is a helper.
	//By doing this when it fails the line number reported will be in our function
	//call rather than inside our test helper.
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
