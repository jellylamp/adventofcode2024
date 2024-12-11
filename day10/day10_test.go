package day10

import "testing"

func TestEasySamplePartA(t *testing.T) {
	total := PartA("./samples/easySample.txt")
	expected := 1
	if total != expected {
		t.Errorf("sample part a = %d; expected %d", total, expected)
	}
}

func TestSplitSamplePartA(t *testing.T) {
	total := PartA("./samples/splitsample.txt")
	expected := 2
	if total != expected {
		t.Errorf("sample part a = %d; expected %d", total, expected)
	}
}

func TestMediumSamplePartA(t *testing.T) {
	total := PartA("./samples/mediumsample.txt")
	expected := 4
	if total != expected {
		t.Errorf("sample part a = %d; expected %d", total, expected)
	}
}

func TestTwoTrailheadsSamplePartA(t *testing.T) {
	total := PartA("./samples/twoTrailheads.txt")
	expected := 3
	if total != expected {
		t.Errorf("sample part a = %d; expected %d", total, expected)
	}
}

func TestSamplePartA(t *testing.T) {
	total := PartA("./samples/sample.txt")
	expected := 36
	if total != expected {
		t.Errorf("sample part a = %d; expected %d", total, expected)
	}
}

func TestPartA(t *testing.T) {
	output := PartA("./samples/input.txt")
	expected := 587
	if output != expected {
		t.Errorf("sample part a = %d; expected %d", output, expected)
	}
}

// func TestEasySamplePartB(t *testing.T) {
// 	output := PartB("./samples/partbsample.txt")
// 	expected := 9
// 	if output != expected {
// 		t.Errorf("sample part b = %d; expected %d", output, expected)
// 	}
// }

// func TestSamplePartB(t *testing.T) {
// 	output := PartB("./samples/sample.txt")
// 	expected := 34
// 	if output != expected {
// 		t.Errorf("sample part b = %d; expected %d", output, expected)
// 	}
// }

// func TestPartB(t *testing.T) {
// 	output := PartB("./samples/input.txt")
// 	expected := 912
// 	if output != expected {
// 		t.Errorf("sample part b = %d; expected %d", output, expected)
// 	}
// }
