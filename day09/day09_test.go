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

func TestEdgeCaseSamplePartB(t *testing.T) {
	output := PartB("./samples/edgecase2.txt")
	expected := 169
	// 0...1...2......33333
	// 0...1...233333......
	// 0...1...233333......
	// 02..1....33333......
	// 021......33333......
	// 021......33333......
	if output != expected {
		t.Errorf("sample part b = %d; expected %d", output, expected)
	}
}

func TestEdgeCase1SamplePartB(t *testing.T) {
	output := PartB("./samples/edgecase.txt")
	expected := 825
	// 15101010101010101010105
	// 0.....123456789 10 11 11 11 11 11
	// 0 11 11 11 11 11 11 123456789 10

	if output != expected {
		t.Errorf("sample part b = %d; expected %d", output, expected)
	}
}

func TestEdgeCase3SamplePartB(t *testing.T) {
	output := PartB("./samples/edgecase3.txt")
	expected := 6204

	if output != expected {
		t.Errorf("sample part b = %d; expected %d", output, expected)
	}
}

func TestEdgeCase4SamplePartB(t *testing.T) {
	output := PartB("./samples/edgecase4.txt")
	expected := 1338

	// 80893804751608292
	// 0000000011111111.........222............3333333.....4..............55.........66
	// 000000001111111166554....2223333333.............................................
	// 000000001111111166554222....3333333.............................................
	// 0+0+0+0+0+0+0+0+ 8*1 + 9*1 + 10*1 + 11*1 + 12*1 + 13*1 + 14*1 + 15*1 + 16*6 + 17*6 + 18*6 + 19*5 + 20*5 + 21*5 + 22*4 + 23*2 + 24*2 + 25*2 + 30*3 + 31*3 + 32*3 + 33*3 + 34*3 + 35*3 + 36*3

	if output != expected {
		t.Errorf("sample part b = %d; expected %d", output, expected)
	}
}

func TestSamplePartB(t *testing.T) {
	output := PartB("./samples/sample.txt")
	expected := 2858
	if output != expected {
		t.Errorf("sample part b = %d; expected %d", output, expected)
	}
}

func TestPartB(t *testing.T) {
	output := PartB("./samples/input.txt")

	expected := 6488291456470
	if output != expected {
		t.Errorf("sample part b = %d; expected %d", output, expected)
	}
}
