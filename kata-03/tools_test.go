package tools

import (
	"testing"
)

func TestPalindrome_should_return_true_if_empty_string(t *testing.T) {
	res := palindrome("")

	if !res {
		t.Fatalf("Assertion failed")
	}
}