package go_helper

import (
	"testing"
)

func TestSliceRemoveIntItem(t *testing.T) {
	var slice = []int{1, 2, 3, 4, 5}
	if SliceRemoveIntItem(slice, 3)[3] == slice[3] {
		t.Errorf("SliceRemoveIntItem work incorrect!")
	}
}
func TestSliceRemoveStringItem(t *testing.T) {
	var slice = []string{"one", "two", "three", "four", "five"}
	if SliceRemoveStringItem(slice, 3)[3] == slice[3] {
		t.Errorf("SliceRemoveStringItem work incorrect!")
	}
}
