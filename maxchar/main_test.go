package main

import "testing"

func TestMaxChar1(t *testing.T) {
	got := MaxChar("ab1c1d1e1f1g1")
	res := "1"
	if got != res {
		t.Errorf("Expected %s got %s", res, got)
	}
}

func TestMaxChar2(t *testing.T) {
	got := MaxChar("a")
	res := "a"
	if got != res {
		t.Errorf("Expected %s got %s", res, got)
	}
}

func TestMaxChar3(t *testing.T) {
	got := MaxChar("abcdefghijklmnaaaaa")
	res := "a"
	if got != res {
		t.Errorf("Expected %s got %s", res, got)
	}
}
