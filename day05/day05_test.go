package day05

import "testing"

func TestSamplePartA(t *testing.T) {
	total := PartA("./samples/sample.txt")
	expected := 143
	if total != expected {
		t.Errorf("sample part a = %d; expected %d", total, expected)
	}
}

func TestPartA(t *testing.T) {
	output := PartA("./samples/input.txt")
	expected := 7198
	if output != expected {
		t.Errorf("sample part a = %d; expected %d", output, expected)
	}
}

// func TestSamplePartB(t *testing.T) {
// 	output := PartB("./samples/partbsample.txt")
// 	var expected int64 = 48
// 	if output != expected {
// 		t.Errorf("sample part b = %d; expected %d", output, expected)
// 	}
// }

// func TestPartB(t *testing.T) {
// 	output := PartB("./samples/parta.txt")
// 	var expected int64 = 88811886
// 	if output != expected {
// 		t.Errorf("sample part a = %d; expected %d", output, expected)
// 	}
// }
