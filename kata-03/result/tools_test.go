// Copyright 2017 The OLX Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package result

import (
	"testing"
)

func TestPalindrome_should_return_true_if_empty_string(t *testing.T) {
	assertTrue(t, palindrome(""))
}

func TestPalindrome_should_return_true_with_one_character(t *testing.T) {
	assertTrue(t, palindrome("a"))
}

func TestPalindrome_should_return_false_with_2_diff_characters(t *testing.T) {
	assertFalse(t, palindrome("ab"))
}

func TestPalindrome_should_return_true_with_2_equal_characters(t *testing.T) {
	assertTrue(t, palindrome("aa"))
}

func TestPalindrome_should_return_true_with_3_equal_characters(t *testing.T) {
	assertTrue(t, palindrome("aba"))
}

func TestPalindrome_should_return_True_with_1space_1characters(t *testing.T) {
	assertTrue(t, palindrome(" a"))
}

func TestPalindrome_should_return_true_with_4_different_characters(t *testing.T) {
	assertFalse(t, palindrome("abca"))
}

func TestPalindrome_should_return_true_with_many_different_characters(t *testing.T) {
	assertTrue(t, palindrome("anotaram a data da maratona"))
}

func TestPalindrome_should_return_true_with_many_different_characters_2(t *testing.T) {
	assertTrue(t, palindrome("LUZA ROCELINA A NAMORADA DO MANUEL LEU NA MODA DA ROMANA ANIL E COR AZUL"))
}

func assertFalse(t *testing.T, b bool) {
	if b {
		t.Fatalf("Assertion failed")
	}
}

func assertTrue(t *testing.T, b bool) {
	if !b {
		t.Fatalf("Assertion failed: must return true")
	}
}
