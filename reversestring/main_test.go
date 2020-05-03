package main

import "testing"

func TestReverse1(t *testing.T) {
	got := Reverse("abc")
	res := "cba"
	if got != res {
		t.Errorf("Expected %s got %s", res, got)
	}
}
func TestReverse2(t *testing.T) {
	got := Reverse("  abcd")
	res := "dcba  "
	if got != res {
		t.Errorf("Expected %s got %s", res, got)
	}
}

func TestReverse3(t *testing.T) {
	got := Reverse1("abc")
	res := "cba"
	if got != res {
		t.Errorf("Expected %s got %s", res, got)
	}
}
func TestReverse4(t *testing.T) {
	got := Reverse1("  abcd")
	res := "dcba  "
	if got != res {
		t.Errorf("Expected %s got %s", res, got)
	}
}
