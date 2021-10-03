package go_helper

func StringReverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

// [3]string{"sobaka","sobaki","sobak"}
func StringDeclension(i int, s [3]string) string {
	if i%100 < 11 || i%100 > 14 {
		switch i % 10 {
		case 1:
			return s[0]
		case 2:
			fallthrough
		case 3:
			fallthrough
		case 4:
			return s[1]
		}
	}
	return s[2]
}
