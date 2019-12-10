package main

import "fmt"

type unexpectedResponseErr struct {
	statusCode int
	body       string
}

func (e *unexpectedResponseErr) Error() string {
	return fmt.Sprintf("error: unexpected response: %v %v", e.statusCode, e.body)
}

type invalidKeyTypeErr struct {
	key string
	val interface{}
}

func (e *invalidKeyTypeErr) Error() string {
	return fmt.Sprintf("error: invalid type for key: type of %v is %T: %v", e.key, e.val, e.val)
}
