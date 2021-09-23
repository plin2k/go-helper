package go_helper

import (
	"testing"
)

func dataProviderSS() (map[string]string, []string, []string) {
	return map[string]string{
			"hello":  "rise",
			"not":    "one",
			"plin2k": "plin3k",
			"only":   "key",
		},
		[]string{"hello", "not", "plin2k", "only"},
		[]string{"rise", "one", "plin3k", "key"}
}

func dataProviderSI() (map[string]int, []string, []int) {
	return map[string]int{
			"hello":  10,
			"not":    3123421,
			"plin2k": 90,
			"only":   2,
		},
		[]string{"hello", "not", "plin2k", "only"},
		[]int{10, 3123421, 90, 2}
}

func dataProviderSF() (map[string]float32, []string, []float32) {
	return map[string]float32{
			"hello":  20.3,
			"not":    1.312312,
			"plin2k": 5343,
			"only":   0.312,
		},
		[]string{"hello", "not", "plin2k", "only"},
		[]float32{20.3, 1.312312, 5343, 0.312}
}

func dataProviderIS() (map[int]string, []int, []string) {
	return map[int]string{
			10:          "rise",
			345334:      "one",
			1:           "plin3k",
			43123123121: "key",
		},
		[]int{10, 345334, 1, 43123123121},
		[]string{"rise", "one", "plin3k", "key"}
}

func dataProviderSB() (map[string]bool, []string, []bool) {
	return map[string]bool{
			"hello":  true,
			"not":    false,
			"plin2k": false,
			"only":   true,
		},
		[]string{"hello", "not", "plin2k", "only"},
		[]bool{true, false, false, true}
}

func dataProviderIB() (map[int]bool, []int, []bool) {
	return map[int]bool{
			12:    true,
			100:   false,
			43232: false,
			0:     true,
		},
		[]int{12, 100, 43232, 0},
		[]bool{true, false, false, true}
}

func TestMapSS(t *testing.T) {
	arr, keys, values := dataProviderSS()
	outKeys := MapSSKeys(arr)
	var compare bool

	for _, v := range outKeys {
		compare = false
		for _, key := range keys {
			if v == key {
				compare = true
			}
		}
		if compare == false {
			t.Fatal("MapSSKeys work incorrect")
		}
	}

	outValues := MapSSValues(arr)
	for _, v := range outValues {
		compare = false
		for _, value := range values {
			if v == value {
				compare = true
			}
		}
		if compare == false {
			t.Fatal("MapSSValues work incorrect")
		}
	}
}

func TestMapSI(t *testing.T) {
	arr, keys, values := dataProviderSI()
	outKeys := MapSIKeys(arr)
	var compare bool

	for _, v := range outKeys {
		compare = false
		for _, key := range keys {
			if v == key {
				compare = true
			}
		}
		if compare == false {
			t.Fatal("MapSIKeys work incorrect")
		}
	}

	outValues := MapSIValues(arr)
	for _, v := range outValues {
		compare = false
		for _, value := range values {
			if v == value {
				compare = true
			}
		}
		if compare == false {
			t.Fatal("MapSIValues work incorrect")
		}
	}
}

func TestMapSF(t *testing.T) {
	arr, keys, values := dataProviderSF()
	outKeys := MapSF32Keys(arr)
	var compare bool

	for _, v := range outKeys {
		compare = false
		for _, key := range keys {
			if v == key {
				compare = true
			}
		}
		if compare == false {
			t.Fatal("MapSF32Keys work incorrect")
		}
	}

	outValues := MapSF32Values(arr)
	for _, v := range outValues {
		compare = false
		for _, value := range values {
			if v == value {
				compare = true
			}
		}
		if compare == false {
			t.Fatal("MapSF32Values work incorrect")
		}
	}
}

func TestMapSB(t *testing.T) {
	arr, keys, values := dataProviderSB()
	outKeys := MapSBKeys(arr)
	var compare bool

	for _, v := range outKeys {
		compare = false
		for _, key := range keys {
			if v == key {
				compare = true
			}
		}
		if compare == false {
			t.Fatal("MapSBKeys work incorrect")
		}
	}

	outValues := MapSBValues(arr)
	for _, v := range outValues {
		compare = false
		for _, value := range values {
			if v == value {
				compare = true
			}
		}
		if compare == false {
			t.Fatal("MapSBValues work incorrect")
		}
	}
}

func TestMapIS(t *testing.T) {
	arr, keys, values := dataProviderIS()
	outKeys := MapISKeys(arr)
	var compare bool

	for _, v := range outKeys {
		compare = false
		for _, key := range keys {
			if v == key {
				compare = true
			}
		}
		if compare == false {
			t.Fatal("MapISKeys work incorrect")
		}
	}

	outValues := MapISValues(arr)
	for _, v := range outValues {
		compare = false
		for _, value := range values {
			if v == value {
				compare = true
			}
		}
		if compare == false {
			t.Fatal("MapISValues work incorrect")
		}
	}
}

func TestMapIB(t *testing.T) {
	arr, keys, values := dataProviderIB()
	outKeys := MapIBKeys(arr)
	var compare bool

	for _, v := range outKeys {
		compare = false
		for _, key := range keys {
			if v == key {
				compare = true
			}
		}
		if compare == false {
			t.Fatal("MapIBKeys work incorrect")
		}
	}

	outValues := MapIBValues(arr)
	for _, v := range outValues {
		compare = false
		for _, value := range values {
			if v == value {
				compare = true
			}
		}
		if compare == false {
			t.Fatal("MapIBValues work incorrect")
		}
	}
}

func TestMapSSMerge(t *testing.T) {
	arr, keys, values := dataProviderSS()
	result := MapSSMerge(keys, values)

	for key, value := range result {
		if arr[key] != value {
			t.Fatal("MapSSMerge work incorrect")
		}
	}
}

func TestMapSIMerge(t *testing.T) {
	arr, keys, values := dataProviderSI()
	result := MapSIMerge(keys, values)

	for key, value := range result {
		if arr[key] != value {
			t.Fatal("MapSIMerge work incorrect")
		}
	}
}

func TestMapSBMerge(t *testing.T) {
	arr, keys, values := dataProviderSB()
	result := MapSBMerge(keys, values)

	for key, value := range result {
		if arr[key] != value {
			t.Fatal("MapSBMerge work incorrect")
		}
	}
}

func TestMapSF32Merge(t *testing.T) {
	arr, keys, values := dataProviderSF()
	result := MapSF32Merge(keys, values)

	for key, value := range result {
		if arr[key] != value {
			t.Fatal("MapSF32Merge work incorrect")
		}
	}
}

func TestMapISMerge(t *testing.T) {
	arr, keys, values := dataProviderIS()
	result := MapISMerge(keys, values)

	for key, value := range result {
		if arr[key] != value {
			t.Fatal("MapISMerge work incorrect")
		}
	}
}

func TestMapIBMerge(t *testing.T) {
	arr, keys, values := dataProviderIB()
	result := MapIBMerge(keys, values)

	for key, value := range result {
		if arr[key] != value {
			t.Fatal("MapIBMerge work incorrect")
		}
	}
}
