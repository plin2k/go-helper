package go_helper

import "testing"

func TestNumberEvenFromInt(t *testing.T) {
	if NumberEvenFromInt(727178) != 28 {
		t.Errorf("NumberEvenFromInt work incorrect")
	}
}

func TestNumberFirstInt(t *testing.T) {
	var num = 735
	if NumberFirstInt(num) != 7 {
		t.Errorf("NumberFirstInt work incorrect")
	}
}
