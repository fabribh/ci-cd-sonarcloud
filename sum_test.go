package main

import "testing"

func TestSum(t *testing.T) {
	result := sum(2, 2)
	expected := 4
	if result != expected {
		t.Errorf("sum(2, 2) = %d, expected %d", result, expected)
	}

	result = sum(-2, 3)
	expected = 1
	if result != expected {
		t.Errorf("sum(-2, 3) = %d, expected %d", result, expected)
	}
}

func TestTimes(t *testing.T) {
	result := times(2, 2)
	expected := 4
	if result != expected {
		t.Errorf("times(2, 2) = %d, expected %d", result, expected)
	}

	result = times(-2, 3)
	expected = -6
	if result != expected {
		t.Errorf("times(-2, 3) = %d, expected %d", result, expected)
	}
}

func TestSub(t *testing.T) {
	result := sub(10, 2)
	expected := 8
	if result != expected {
		t.Errorf("sub(10, 2) = %d, expected %d", result, expected)
	}
}

func TestDiv(t *testing.T) {
	result := div(8, 2)
	expected := 4
	if result != expected {
		t.Errorf("div(8, 2) = %d, expected %d", result, expected)
	}
}
