package main

import "fmt"

type unexpectedResponseErr struct {
	statusCode int
	body       string
}

func (e *unexpectedResponseErr) Error() string {
	return fmt.Sprintf("error: unexpected response: %v %v", e.statusCode, e.body)
}
