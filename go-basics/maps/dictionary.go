package maps

import "errors"

var (
	ErrNotFound      = errors.New("could not find the word you were looking for")
	ErrAlreadyExists = errors.New("word already exists")
)

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	if _, err := d.Search(word); err == nil {
		return ErrAlreadyExists
	}
	d[word] = definition
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	if _, err := d.Search(word); err != nil {
		return err
	}
	d[word] = definition
	return nil
}

func (d Dictionary) Delete(word string)  {
	delete(d, word)
}

type Dictionary map[string]string
