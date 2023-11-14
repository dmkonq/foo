package main

import "errors"

func foo() (string, error) {
	return "", errors.New("not implemented")
}

func bar() (string, error) {
	// asdfkjl
	return "bar", nil
}

//asdf

func baz() (string, error) {
	return "baz", nil
}

func qux() (string, error) {
	return "", errors.New("not implemented")
}

func main() {
}
