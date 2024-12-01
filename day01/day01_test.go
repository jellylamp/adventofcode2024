package day01

import "testing"

func TestSamplePartA(t *testing.T) {
    output := Run("./samples/sample.txt")
	expected := 11
	if output != expected {
		t.Errorf("sample part a = %d; expected %d", output, expected)
	}
}

func TestPartA(t *testing.T) {
    output := Run("./samples/parta.txt")
	expected := 1506483
	if output != expected {
		t.Errorf("sample part a = %d; expected %d", output, expected)
	}
}

