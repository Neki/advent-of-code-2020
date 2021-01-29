package day2

import "testing"

func TestParsePasswordInfo(t *testing.T) {
	_, err := ParsePasswordInfo("nope")
	if err == nil {
		t.Errorf("ParsePasswordInfo did not return an error on invalid input")
	}
	input := "1-12 a: bkjsdficc"
	got, err := ParsePasswordInfo(input)
	if err != nil {
		t.Errorf("ParsePasswordInfo returned an error on input '%s': '%s'", input, err)
	}
	if got.MandatoryChar != 'a' {
		t.Errorf("ParsePasswordInfo('%s'): got MandatoryChar = '%v', expected 'a'", input, got.MandatoryChar)
	}
	if got.Password != "bkjsdficc" {
		t.Errorf("ParsePasswordInfo('%s'): got Password = '%v', expected 'bkjsdficc'", input, got.Password)
	}
	if got.OccurenceRange.max != 12 {
		t.Errorf("ParsePasswordInfo('%s'): got Range.max = '%v', expected '12'", input, got.OccurenceRange.min)
	}
}

func TestPasswordMatchesOldPolicy(t *testing.T) {
	val, _ := ParsePasswordInfo("1-2 b: abbbbca")
	if val.MatchesOldPolicy() {
		t.Errorf("MatchesOldPolicy() returned true, expected false")
	}
	val, _ = ParsePasswordInfo("1-2 c: abbbbca")
	if !val.MatchesOldPolicy() {
		t.Errorf("MatchesOldPolicy() returned false, expected true")
	}
}
