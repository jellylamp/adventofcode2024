package day01

import "testing"

func TestSamplePartA(t *testing.T) {
	output := PartA("./samples/sample.txt")
	expected := 11
	if output != expected {
		t.Errorf("sample part a = %d; expected %d", output, expected)
	}
}

func TestPartA(t *testing.T) {
	output := PartA("./samples/parta.txt")
	expected := 1506483
	if output != expected {
		t.Errorf("sample part a = %d; expected %d", output, expected)
	}
}

func TestSamplePartB(t *testing.T) {
	output := PartB("./samples/sample.txt")
	expected := 31
	if output != expected {
		t.Errorf("sample part a = %d; expected %d", output, expected)
	}
}

func TestPartB(t *testing.T) {
	output := PartB("./samples/parta.txt")
	expected := 23126924
	if output != expected {
		t.Errorf("sample part a = %d; expected %d", output, expected)
	}
}