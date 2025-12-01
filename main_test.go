package main

import "testing"

func TestCalculateNoOfZeros_Case2(t *testing.T) {
	list := []Movement{
		{distance: 68, side: "L"},
		{distance: 30, side: "L"},
		{distance: 48, side: "R"},
		{distance: 5, side: "L"},
		{distance: 60, side: "R"},
		{distance: 55, side: "L"},
		{distance: 1, side: "L"},
		{distance: 99, side: "L"},
		{distance: 14, side: "R"},
		{distance: 82, side: "L"},
	}

	got, err := calculateNoOfZeros(list)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	want := int32(6)
	if got != want {
		t.Errorf("Expected %d zeros, but got %d", want, got)
	}
}
