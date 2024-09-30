package maps

import "testing"

func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q , want %q", got, want)
	}
}

func assertError(t testing.TB, err error, want error) {
	t.Helper()
	if err != want {
		t.Errorf("got error %q want %q", err, want)
	}
}

func TestDictionary(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertString(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")
		if got == nil {
			t.Fatal("expected to get an error.")
		}
		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "this is just a test")

	want := "this is just a test"
	got, err := dictionary.Search("test")
	if err != nil {
		t.Fatal("expected error but got nothing")
	}
	assertString(t, got, want)
}
