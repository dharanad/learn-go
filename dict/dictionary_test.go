package dict

import "testing"

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
func assertNotNilErr(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Errorf("expecting an error")
	}
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertNilErr(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Errorf("expecting nil error but got %q", got)
	}
}

func TestSearch(t *testing.T) {
	t.Run("known word", func(t *testing.T) {
		dict := Dictionary{"test": "Hello, World"}
		got, _ := dict.Search("test")
		want := "Hello, World"
		assertString(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		dict := Dictionary{"test": "Hello, World"}
		_, err := dict.Search("test1")
		assertNotNilErr(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("word exists", func(t *testing.T) {
		dict := make(Dictionary)
		err := dict.Add("test", "Hello, World")
		assertNilErr(t, err)
		got, _ := dict.Search("test")
		want := "Hello, World"
		assertString(t, got, want)
	})
	t.Run("word not exists", func(t *testing.T) {
		dict := make(Dictionary)
		_ = dict.Add("test", "Hello, World")
		err := dict.Add("test", "Hello, World")
		assertNotNilErr(t, err, ErrKeyExists)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update existing key", func(t *testing.T) {
		d := Dictionary{"A": "B"}
		err := d.Update("A", "C")
		assertNilErr(t, err)
		got, _ := d.Search("A")
		want := "C"
		assertString(t, got, want)
	})
	t.Run("update non existing key", func(t *testing.T) {
		d := make(Dictionary)
		err := d.Update("A", "C")
		assertNotNilErr(t, err, ErrKeyNotExists)
	})
}
