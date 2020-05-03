package main

import "testing"

func TestReverseInt1(t *testing.T) {
	got := ReverseInt(123)
	res := 321
	if got != res {
		t.Errorf("Expected %d got %d", res, got)
	}
}

func TestReverseInt2(t *testing.T) {
	got := ReverseInt(-15)
	res := -51
	if got != res {
		t.Errorf("Expected %d got %d", res, got)
	}
}

func TestReverseInt3(t *testing.T) {
	got := ReverseInt(121121)
	res := 121121
	if got != res {
		t.Errorf("Expected %d got %d", res, got)
	}
}

func TestReverseInt4(t *testing.T) {
	got := ReverseInt(-90)
	res := -9
	if got != res {
		t.Errorf("Expected %d got %d", res, got)
	}
}
