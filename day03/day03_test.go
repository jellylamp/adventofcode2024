package day03

import "testing"

func TestSamplePartA(t *testing.T) {
	total := PartA("./samples/sample.txt")
	var expected int64 = 161
	if total != expected {
		t.Errorf("sample part a = %d; expected %d", total, expected)
	}
}

func TestSamplePartAEdgeCases(t *testing.T) {
	total := PartA("./samples/sampleedge.txt")
	var expected int64 = 19043
	if total != expected {
		t.Errorf("sample part a = %d; expected %d", total, expected)
	}
}

func TestPartA(t *testing.T) {
	output := PartA("./samples/parta.txt")
	var expected int64 = 166357705
	if output != expected {
		t.Errorf("sample part a = %d; expected %d", output, expected)
	}
}

// func TestSamplePartB(t *testing.T) {
// 	output := PartB("./samples/sample.txt")
// 	expected := 31
// 	if output != expected {
// 		t.Errorf("sample part b = %d; expected %d", output, expected)
// 	}
// }

// func TestPartB(t *testing.T) {
// 	output := PartB("./samples/parta.txt")
// 	expected := 23126924
// 	if output != expected {
// 		t.Errorf("sample part b = %d; expected %d", output, expected)
// 	}
// }
