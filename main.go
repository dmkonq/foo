package main

import "errors"

func foo() (string, error) {
	return "", errors.New("not implemented")
}

func bar() (string, error) {
	return "", errors.New("not implemented")
}

func baz() (string, error) {
	return "foo", nil
}

func main() {
}
