package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

// TestHellos calls greetings.Hellos with a list of names,
// checking for valid return values.
func TestHellos(t *testing.T) {
	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err := Hellos(names)
	if err != nil {
		t.Fatalf(`Hellos(%v) returned an error: %v`, names, err)
	}
	for _, name := range names {
		want := regexp.MustCompile(`\b` + name + `\b`)
		msg, ok := messages[name]
		if !ok || !want.MatchString(msg) {
			t.Fatalf(`Hellos(%v) = %v, want match for %#q`, names, messages, want)
		}
	}
}

// TestHellosEmpty calls greetings.Hellos with an empty list,
// checking for an empty map and no error.
func TestHellosEmpty(t *testing.T) {
	names := []string{}
	messages, err := Hellos(names)
	if err != nil {
		t.Fatalf(`Hellos(%v) returned an error: %v`, names, err)
	}
	if len(messages) != 0 {
		t.Fatalf(`Hellos(%v) = %v, want empty map`, names, messages)
	}
}
