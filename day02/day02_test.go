package day02

import "testing"

func TestSamplePartA(t *testing.T) {
	output := PartA("./samples/sample.txt")
	expected := 2
	if output != expected {
		t.Errorf("sample part a = %d; expected %d", output, expected)
	}
}

func TestPartA(t *testing.T) {
	output := PartA("./samples/parta.txt")
	expected := 230
	if output != expected {
		t.Errorf("part a = %d; expected %d", output, expected)
	}
}

func TestSamplePartB(t *testing.T) {
	output := PartB("./samples/sample.txt")
	expected := 38
	if output != expected {
		t.Errorf("sample part b = %d; expected %d", output, expected)
	}
}

func TestPartB(t *testing.T) {
	output := PartB("./samples/parta.txt")
	expected := 301
	if output != expected {
		t.Errorf("part b = %d; expected %d", output, expected)
	}
}
