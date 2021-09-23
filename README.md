# go-helper
 go-helper

# Get Keys of map
```go
package main

import h "github.com/plin2k/go-helper"

func main () {
	var myMapSS = make(map[string]string,10)
	stringKeys := h.MapSSKeys(myMapSS) // return []string
	
	var myMapSI = make(map[string]int,10)
	stringKeys := h.MapSIKeys(myMapSI) // return []string

	var myMapSB = make(map[string]bool,10)
	stringKeys := h.MapSBKeys(myMapSB) // return []string

	var myMapSF32 = make(map[string]float32,10)
	stringKeys := h.MapSF32Keys(myMapSF32) // return []string

	var myMapSF64 = make(map[string]float64,10)
	stringKeys := h.MapSF64Keys(myMapSF64) // return []string

	var myMapIS = make(map[int]string,10)
	intKeys := h.MapISKeys(myMapIS) // return []int

	var myMapIB = make(map[int]bool,10)
	intKeys := h.MapIBKeys(myMapIB) // return []int
}
```

# Get Values of map
```go
package main

import h "github.com/plin2k/go-helper"

func main () {
	var myMapSS = make(map[string]string,10)
	stringValues := h.MapSSValues(myMapSS) // return []string
	
	var myMapSI = make(map[string]int,10)
	intValues := h.MapSIValues(myMapSI) // return []int

	var myMapSB = make(map[string]bool,10)
	boolValues := h.MapSBValues(myMapSB) // return []bool

	var myMapSF32 = make(map[string]float32,10)
	float32Values := h.MapSF32Values(myMapSF32) // return []float32

	var myMapSF64 = make(map[string]float64,10)
	float64Values := h.MapSF64Values(myMapSF64) // return []float64

	var myMapIS = make(map[int]string,10)
	stringValues := h.MapISValues(myMapIS) // return []string

	var myMapIB = make(map[int]bool,10)
	boolValues := h.MapIBValues(myMapIB) // return []bool
}
```

# Merge Slices to map
```go
package main

import h "github.com/plin2k/go-helper"

func main () {
	var myStringSlice1 = []string{"hello","bye"}
	var myStringSlice2 = []string{"hello","bye"}
	newMapSS := h.MapSSMerge(myStringSlice1,myStringSlice2) // return map[string]string

	var myIntSlice1 = []int{10,20}
	var myStringSlice2 = []string{"hello","bye"}
	newMapIS := h.MapISMerge(myIntSlice1,myStringSlice2) // return map[int]string

	var myStringSlice1 = []string{"hello","bye"}
	var myIntSlice2 = []int{10,20}
	newMapSI := h.MapSIMerge(myStringSlice1,myIntSlice2) // return map[string]int

	var myStringSlice1 = []string{"hello","bye"}
	var myFloat32Slice2 = []float32{10.1,20.31}
	newMapSF32 := h.MapSF32Merge(myStringSlice1,myFloat32Slice2) // return map[string]float32

	var myStringSlice1 = []string{"hello","bye"}
	var myFloat64Slice2 = []float64{10.1,20.31}
	newMapSF64 := h.MapSF64Merge(myStringSlice1,myFloat64Slice2) // return map[string]float64

	var myStringSlice1 = []string{"hello","bye"}
	var myBoolSlice2 = []bool{true,false}
	newMapSB := h.MapSBMerge(myStringSlice1,myBoolSlice2) // return map[string]bool

	var myIntSlice1 = []int{10,20}
	var myBoolSlice2 = []bool{true,false}
	newMapIB := h.MapIBMerge(myIntSlice1,myBoolSlice2) // return map[int]bool

}
```