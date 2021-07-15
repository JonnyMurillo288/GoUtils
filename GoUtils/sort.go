package utils

func Sort(arr []int) []int {
	var lowest,ind int
	for i := 0; i < len(arr); i++ {
		lowest = arr[i] +100000
		for j := 0 + i; j < len(arr); j++ {
			if arr[j] < lowest {
				lowest = arr[j]
				ind = j
			}
		}
		x := arr[i]
		arr[i] = lowest
		arr[ind] = x
	}
	return arr
}

func SortFloat(arr []float64) []float64 {
	var lowest float64
	var ind int
	for i := 0 ; i < len(arr); i++ {
		lowest = arr[i] +100000
		for j := 0+ i; j < len(arr); j++ {
			if arr[j] < lowest {
				lowest = arr[j]
				ind = j
			}
		}
		x := arr[i]
		arr[i] = lowest
		arr[ind] = x
	}
	return arr
}