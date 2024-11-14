package dictionaries

type Dictionary map[string]string
type DictionaryError string

const (
	ErrWordNotFound  = DictionaryError("could not find the given word")
	ErrWordExists = DictionaryError("word exists already")
	ErrWordNotExists = DictionaryError("word doesn't exist already")
)

func (e DictionaryError) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	found, ok := d[word]

	if !ok {
		return "", ErrWordNotFound
	}
	return found, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrWordNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrWordNotFound:
        return ErrWordNotExists
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) {
    delete(d, word)
}
