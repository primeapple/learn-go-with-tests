package dictionaries

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}

	t.Run("find existing word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is a test"

		assertString(t, got, want)
	})

	t.Run("give error on non existing word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		assertError(t, err, ErrWordNotFound)
		assertString(t, err.Error(), ErrWordNotFound.Error())
	})
}

func TestAdd(t *testing.T) {
	word := "test"
	definition := "this is a test"
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "this is another test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	word := "test"
	definition := "this is a test"
    newDefinition := "this is another Test"

	t.Run("should update existing word", func(t *testing.T) {
        dictionary := Dictionary{word: definition}
        err := dictionary.Update(word, newDefinition)

        assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
    })

	t.Run("should error on new word", func(t *testing.T) {
        dictionary := Dictionary{}
        err := dictionary.Update(word, newDefinition)

		assertError(t, err, ErrWordNotExists)
    })
}

func TestDelete(t *testing.T) {
	word := "test"
	definition := "this is a test"

	t.Run("should delete existing word", func(t *testing.T) {
        dictionary := Dictionary{word: definition}
        dictionary.Delete(word)

        _, err := dictionary.Search(word)

		assertError(t, err, ErrWordNotFound)
    })
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("Got %q wanted %q", got.Error(), want.Error())
	}
}

func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("Got %q wanted %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("Should find added word")
	}

	assertString(t, got, definition)
}
