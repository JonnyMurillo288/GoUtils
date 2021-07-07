package utils

import "fmt"

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
	ind := 0
	match := len(arr)
	for _,a := range all {
		fmt.Println(a)
		fmt.Println(arr)
		for _,i := range arr {
			if InIntArr(a,i) {
				ind++
				if ind == match {
					ind = 0
					return true
				}
			}
		}
	}
	return false
}