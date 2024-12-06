package day05

import "testing"

func TestSamplePartA(t *testing.T) {
	total := PartA("./samples/sample.txt")
	expected := 41
	if total != expected {
		t.Errorf("sample part a = %d; expected %d", total, expected)
	}
}

func TestPartA(t *testing.T) {
	output := PartA("./samples/input.txt")
	expected := 5305
	if output != expected {
		t.Errorf("sample part a = %d; expected %d", output, expected)
	}
}

func TestSamplePartB(t *testing.T) {
	output := PartB("./samples/sample.txt")
	expected := 6
	if output != expected {
		t.Errorf("sample part b = %d; expected %d", output, expected)
	}
}

func TestPartB(t *testing.T) {
	output := PartB("./samples/input.txt")
	// 2106 is too low
	// 2170 is too high

	// 100 - 1885
	// 200 - 1884
	// 500 - 1884
	// 530 - 1884 // all too low
	expected := 2106 // too low
	if output != expected {
		t.Errorf("sample part b = %d; expected %d", output, expected)
	}
}
