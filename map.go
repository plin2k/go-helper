package go_helper

// KEYS AND VALUES

// map[string]string
func MapSSKeys (input map[string]string) []string {
	var keys = make([]string,0,len(input))
	for key,_ := range input{
		keys = append(keys, key)
	}
	return keys
}

// map[string]string
func MapSSValues (input map[string]string) []string {
	var values = make([]string,0,len(input))
	for _,value := range input{
		values = append(values, value)
	}
	return values
}

// map[int]string
func MapISKeys (input map[int]string) []int {
	var keys = make([]int,0,len(input))
	for key,_ := range input{
		keys = append(keys, key)
	}
	return keys
}

// map[int]string
func MapISValues (input map[int]string) []string {
	var values = make([]string,0,len(input))
	for _,value := range input{
		values = append(values, value)
	}
	return values
}

// map[string]int
func MapSIKeys (input map[string]int) []string {
	var keys = make([]string,0,len(input))
	for key,_ := range input{
		keys = append(keys, key)
	}
	return keys
}

// map[string]int
func MapSIValues (input map[string]int) []int {
	var values = make([]int,0,len(input))
	for _,value := range input{
		values = append(values, value)
	}
	return values
}

// map[string]float32
func MapSF32Keys (input map[string]float32) []string {
	var keys = make([]string,0,len(input))
	for key,_ := range input{
		keys = append(keys, key)
	}
	return keys
}

// map[string]float32
func MapSF32Values (input map[string]float32) []float32 {
	var values = make([]float32,0,len(input))
	for _,value := range input{
		values = append(values, value)
	}
	return values
}

// map[string]float64
func MapSF64Keys (input map[string]float64) []string {
	var keys = make([]string,0,len(input))
	for key,_ := range input{
		keys = append(keys, key)
	}
	return keys
}

// map[string]float64
func MapSF64Values (input map[string]float64) []float64 {
	var values = make([]float64,0,len(input))
	for _,value := range input{
		values = append(values, value)
	}
	return values
}

// map[string]bool
func MapSBKeys (input map[string]bool) []string {
	var keys = make([]string,0,len(input))
	for key,_ := range input{
		keys = append(keys, key)
	}
	return keys
}

// map[string]bool
func MapSBValues (input map[string]bool) []bool {
	var values = make([]bool,0,len(input))
	for _,value := range input{
		values = append(values, value)
	}
	return values
}

// map[int]bool
func MapIBKeys (input map[int]bool) []int {
	var keys = make([]int,0,len(input))
	for key,_ := range input{
		keys = append(keys, key)
	}
	return keys
}

// map[int]bool
func MapIBValues (input map[int]bool) []bool {
	var values = make([]bool,0,len(input))
	for _,value := range input{
		values = append(values, value)
	}
	return values
}

// MERGE

// map[string]string
func MapSSMerge (keys []string, values []string) map[string]string {
	var outMap = make(map[string]string,len(keys))
	for key, value := range keys {
		outMap[value] = values[key]
	}
	return outMap
}

// map[int]string
func MapISMerge (keys []int, values []string) map[int]string {
	var outMap = make(map[int]string,len(keys))
	for key, value := range keys {
		outMap[value] = values[key]
	}
	return outMap
}

// map[string]int
func MapSIMerge (keys []string, values []int) map[string]int {
	var outMap = make(map[string]int,len(keys))
	for key, value := range keys {
		outMap[value] = values[key]
	}
	return outMap
}

// map[string]float32
func MapSF32Merge (keys []string, values []float32) map[string]float32 {
	var outMap = make(map[string]float32,len(keys))
	for key, value := range keys {
		outMap[value] = values[key]
	}
	return outMap
}

// map[string]float64
func MapSF64Merge (keys []string, values []float64) map[string]float64 {
	var outMap = make(map[string]float64,len(keys))
	for key, value := range keys {
		outMap[value] = values[key]
	}
	return outMap
}

// map[string]bool
func MapSBMerge (keys []string, values []bool) map[string]bool {
	var outMap = make(map[string]bool,len(keys))
	for key, value := range keys {
		outMap[value] = values[key]
	}
	return outMap
}

// map[int]bool
func MapIBMerge (keys []int, values []bool) map[int]bool {
	var outMap = make(map[int]bool,len(keys))
	for key, value := range keys {
		outMap[value] = values[key]
	}
	return outMap
}