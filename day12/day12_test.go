package day12

import "testing"

func TestEasySamplePartA(t *testing.T) {
	total := PartA("./samples/easysample.txt")
	expected := 140
	if total != expected {
		t.Errorf("sample part a = %d; expected %d", total, expected)
	}
}

func TestMediumSamplePartA(t *testing.T) {
	total := PartA("./samples/mediumsample.txt")
	expected := 772
	if total != expected {
		t.Errorf("sample part a = %d; expected %d", total, expected)
	}
}

func TestSamplePartA(t *testing.T) {
	total := PartA("./samples/sample.txt")
	expected := 1930
	if total != expected {
		t.Errorf("sample part a = %d; expected %d", total, expected)
	}
}

func TestPartA(t *testing.T) {
	output := PartA("./samples/input.txt")
	expected := 1489582
	if output != expected {
		t.Errorf("sample part a = %d; expected %d", output, expected)
	}
}

// func TestPartB(t *testing.T) {
// 	output := PartB("./samples/input.txt")
// 	expected := 1340
// 	if output != expected {
// 		t.Errorf("sample part b = %d; expected %d", output, expected)
// 	}
// }
