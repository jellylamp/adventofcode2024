package day11

import "testing"

func TestSamplePartA(t *testing.T) {
	total := PartA("./samples/sample.txt")
	expected := 55312
	if total != expected {
		t.Errorf("sample part a = %d; expected %d", total, expected)
	}
}

func TestPartA(t *testing.T) {
	output := PartA("./samples/input.txt")
	expected := 175006
	if output != expected {
		t.Errorf("sample part a = %d; expected %d", output, expected)
	}
}

func TestSamplePartB25(t *testing.T) {
	total := PartB("./samples/sample.txt")
	expected := 55312
	if total != expected {
		t.Errorf("sample part a = %d; expected %d", total, expected)
	}
}

func TestPartB(t *testing.T) {
	output := PartB("./samples/input.txt")
	expected := 207961583799296
	if output != expected {
		t.Errorf("sample part b = %d; expected %d", output, expected)
	}
}
