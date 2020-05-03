package main

import "testing"

func TestPalindrom1(t *testing.T) {
	got := IsPalindrome("abba")
	if got != true {
		t.Errorf("Expected a palindrome")
	}
}

func TestPalindrom2(t *testing.T) {
	got := IsPalindrome("abcdefg")
	if got != false {
		t.Errorf("Not Expecting a palindrome")
	}
}
func TestPalindrom3(t *testing.T) {
	got := IsPalindrome("ablalba")
	if got != true {
		t.Errorf("Expected a palindrome")
	}
}

func TestPalindrom4(t *testing.T) {
	got := IsPalindrome("1000000001")
	if got != true {
		t.Errorf("Expected a palindrome")
	}
}

func TestPalindrom5(t *testing.T) {
	got := IsPalindrome("aba ")
	if got != false {
		t.Errorf("Not Expecting a palindrome")
	}
}
