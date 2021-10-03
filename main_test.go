package main

import "testing"

func TestAdjustString(t *testing.T) {
	want := "!"
	result := adjustString([]byte(want)[0], want)
	if string(result) != want {
		t.Fatalf("Failed test return %v, want %v.", string(result), want)
	}
}
