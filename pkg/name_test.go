package pkg

import (
	"fmt"
	"regexp"
	"testing"
)

//test for a valid value
func TestNameValid(t *testing.T) {
	name := "Manish"
	want := regexp.MustCompile(`\b` + name + `\b`)
	fmt.Println(want)
	msg, err := Name("Manish")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Name("Manish") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestNameEmpty(t *testing.T) {
	msg, err := Name("")
	if msg != "" || err == nil {
		t.Fatalf(`Name("") = %q, %v, want "" error`, msg, err)
	}
}
