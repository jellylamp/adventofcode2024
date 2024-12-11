package day09

import "testing"

func TestSamplePartA(t *testing.T) {
	total := PartA("./samples/sample.txt")
	expected := 1928
	if total != expected {
		t.Errorf("sample part a = %d; expected %d", total, expected)
	}
}

func TestSamplePartAEdgeCase(t *testing.T) {
	total := PartA("./samples/edgecase.txt")
	// 15101010101010101010105
	// 0.....12345678910-11-11-11-11-11
	// 0 11 11 11 11 11 12345678910
	// 0 * 0 + 11 * 1 + 11 * 2 + 11 * 3 + 11 * 4 + 11 * 5 + 1 * 6 + 2 * 7 + 3 * 8 + 4 * 9 + 5 * 10 + 6 * 11 + 7 * 12 + 8 * 13 + 9 * 14 + 10 * 15

	expected := 825
	if total != expected {
		t.Errorf("sample part a = %d; expected %d", total, expected)
	}
}

func TestPartA(t *testing.T) {
	output := PartA("./samples/input.txt")
	expected := 6461526305538
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

func TestSamplePartB(t *testing.T) {
	output := PartB("./samples/sample.txt")
	expected := 2858
	if output != expected {
		t.Errorf("sample part b = %d; expected %d", output, expected)
	}
}

// func TestPartB(t *testing.T) {
// 	output := PartB("./samples/input.txt")
// 	expected := 912
// 	if output != expected {
// 		t.Errorf("sample part b = %d; expected %d", output, expected)
// 	}
// }
