package main

import (
	"bytes"
	"fmt"
)

type Dictionary map[string]string

const (
	ErrNotFound      = DictionaryErr("could not find the word you were looking for")
	ErrWordExists    = DictionaryErr("cannot add word because it already exists")
	ErrWordNonExists = DictionaryErr("不存在的值不能更新")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		// Map是引用类型，可以不用指针就能修改
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
	case ErrNotFound:
		return ErrWordNonExists
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

func Greet(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
