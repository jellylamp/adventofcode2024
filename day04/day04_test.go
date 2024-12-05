package day04

import "testing"

func TestSamplePartA(t *testing.T) {
	total := PartA("./samples/sample.txt")
	var expected int = 18
	if total != expected {
		t.Errorf("sample part a = %d; expected %d", total, expected)
	}
}

func TestEastSamplePartA(t *testing.T) {
	total := PartA("./samples/eastSample.txt")
	var expected int = 1
	if total != expected {
		t.Errorf("sample part a = %d; expected %d", total, expected)
	}
}

func TestWestSamplePartA(t *testing.T) {
	total := PartA("./samples/westSample.txt")
	var expected int = 1
	if total != expected {
		t.Errorf("sample part a = %d; expected %d", total, expected)
	}
}

func TestNorthSamplePartA(t *testing.T) {
	total := PartA("./samples/northSample.txt")
	var expected int = 1
	if total != expected {
		t.Errorf("sample part a = %d; expected %d", total, expected)
	}
}

func TestPartA(t *testing.T) {
	output := PartA("./samples/parta.txt")
	var expected = 2591
	if output != expected {
		t.Errorf("sample part a = %d; expected %d", output, expected)
	}
}

func TestSamplePartB(t *testing.T) {
	output := PartB("./samples/partbSample.txt")
	var expected = 9
	if output != expected {
		t.Errorf("sample part b = %d; expected %d", output, expected)
	}
}

func TestXSamplePartB(t *testing.T) {
	output := PartB("./samples/xSample.txt")
	var expected = 1
	if output != expected {
		t.Errorf("sample part b = %d; expected %d", output, expected)
	}
}

func TestXBottomRowSamplePartB(t *testing.T) {
	output := PartB("./samples/xBottomRowSample.txt")
	var expected = 4
	if output != expected {
		t.Errorf("sample part b = %d; expected %d", output, expected)
	}
}

func TestPartB(t *testing.T) {
	output := PartB("./samples/parta.txt")
	var expected = 1880
	if output != expected {
		t.Errorf("sample part b = %d; expected %d", output, expected)
	}
}
