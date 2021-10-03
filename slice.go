package go_helper

func SliceRemoveIntItem(slice []int, i int) []int {
	var newSlice = make([]int, 0, len(slice)-1)
	for k, v := range slice {
		if k == i {
			continue
		}
		newSlice = append(newSlice, v)
	}
	return newSlice
}

func SliceRemoveStringItem(slice []string, i int) []string {
	var newSlice = make([]string, 0, len(slice)-1)
	for k, v := range slice {
		if k == i {
			continue
		}
		newSlice = append(newSlice, v)
	}
	return newSlice
}
