package main

const (
	ErrWordNotFound     = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("word already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

// normal errors cannot be constants

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrWordNotFound
	}
	return definition, nil

}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrWordNotFound:
		d[word] = definition
	case nil:
		// search didn't return an error, so the word exists
		// hence word exists
		return ErrWordExists
	default:
		// some other error occurred
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrWordNotFound:
		return ErrWordDoesNotExist
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

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string
type DictionaryErr string
