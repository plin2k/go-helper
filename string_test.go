package go_helper

import "testing"

func TestStringReverse(t *testing.T) {
	var str = "hello"
	var strNew = "olleh"
	if StringReverse(str) != strNew {
		t.Errorf("StringReverse work incorrect")
	}
}

func TestStringDeclension(t *testing.T) {
	var arr = [3]string{"sobaka", "sobaki", "sobak"}
	if StringDeclension(10, arr) != arr[2] {
		t.Errorf("StringDeclension work incorrect")
	}
	if StringDeclension(21, arr) != arr[0] {
		t.Errorf("StringDeclension work incorrect")
	}
	if StringDeclension(999923, arr) != arr[1] {
		t.Errorf("StringDeclension work incorrect")
	}
}
