package day08

import "testing"

func TestSamplePartA(t *testing.T) {
	total := PartA("./samples/sample.txt")
	expected := 14
	if total != expected {
		t.Errorf("sample part a = %d; expected %d", total, expected)
	}
}

func TestPartA(t *testing.T) {
	output := PartA("./samples/input.txt")
	expected := 244
	if output != expected {
		t.Errorf("sample part a = %d; expected %d", output, expected)
	}
}

func TestEasySamplePartB(t *testing.T) {
	output := PartB("./samples/partbsample.txt")
	expected := 9
	if output != expected {
		t.Errorf("sample part b = %d; expected %d", output, expected)
	}
}

func TestSamplePartB(t *testing.T) {
	output := PartB("./samples/sample.txt")
	expected := 34
	if output != expected {
		t.Errorf("sample part b = %d; expected %d", output, expected)
	}
}

func TestPartB(t *testing.T) {
	output := PartB("./samples/input.txt")
	expected := 912
	if output != expected {
		t.Errorf("sample part b = %d; expected %d", output, expected)
	}
}
