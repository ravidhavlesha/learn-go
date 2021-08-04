package greetings

import (
	"regexp"
	"testing"
)

func TestGreetEmpty(t *testing.T) {
	msg, err := Greet("")
	if msg != "" || err == nil {
		t.Fatalf(`Greet("") = %q, %v, want "", error`, msg, err)
	}
}

func TestGreet(t *testing.T) {
	name := "Ravi"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Greet(name)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Greet("Ravi)=%q, %v, want match for %#q,nil`, msg, err, want)
	}
}
