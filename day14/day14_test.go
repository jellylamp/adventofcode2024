package day14

import "testing"

func TestSingleSamplePartA(t *testing.T) {
	total := PartA("./samples/singlesample.txt", 11, 7, 3)
	expected := 0
	if total != expected {
		t.Errorf("sample part a = %d; expected %d", total, expected)
	}
}

func TestSamplePartA(t *testing.T) {
	total := PartA("./samples/sample.txt", 11, 7, 100)
	expected := 12
	if total != expected {
		t.Errorf("sample part a = %d; expected %d", total, expected)
	}
}

func TestPartA(t *testing.T) {
	output := PartA("./samples/input.txt", 101, 103, 100)
	expected := 231221760
	if output != expected {
		t.Errorf("sample part a = %d; expected %d", output, expected)
	}
}

// func TestSingleSamplePartB(t *testing.T) {
// 	total := PartB("./samples/singlesample.txt")
// 	expected := 0
// 	if total != expected {
// 		t.Errorf("sample part a = %d; expected %d", total, expected)
// 	}
// }

// func TestTwoSamplesPartB(t *testing.T) {
// 	total := PartB("./samples/twosamples.txt")
// 	expected := 280
// 	if total != expected {
// 		t.Errorf("sample part a = %d; expected %d", total, expected)
// 	}
// }

// func TestSamplePart3PartB(t *testing.T) {
// 	total := PartB("./samples/thirdsample.txt")
// 	expected := 0
// 	if total != expected {
// 		t.Errorf("sample part a = %d; expected %d", total, expected)
// 	}
// }

// func TestSamplePartB(t *testing.T) {
// 	total := PartB("./samples/sample.txt")
// 	expected := 480
// 	if total != expected {
// 		t.Errorf("sample part a = %d; expected %d", total, expected)
// 	}
// }

// func TestPartB(t *testing.T) {
// 	output := PartB("./samples/input.txt")
// 	expected := 207961583799296
// 	if output != expected {
// 		t.Errorf("sample part b = %d; expected %d", output, expected)
// 	}
// }
