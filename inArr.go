package utils

func InIntArr(arr []int, a int) bool{
	for _,p := range arr {
		if p == a {
			return true
		}
	}
	return false
}

func InStringArr(arr []string, a string) bool {
	for _,p := range arr {
		if p == a {
			return true
		}
	}
	return false
}

func InIntArrArr(all [][]int, arr []int) bool{
	for _,a := range all {
		for _,i := range arr {
			if !InIntArr(a,i) {
				return false
			}
		}
	}
	return true
}