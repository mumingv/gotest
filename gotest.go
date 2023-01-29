package gotest

import (
	"github.com/mumingv/golib"
)

// Hello returns a greeting for the named person.
func TestHello(name string) string {
	message, err := golib.Hello(name)
	if err != nil {
		return "err occur"
	}
	return message
}
