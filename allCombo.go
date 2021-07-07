package utils

import (
	"math/rand"
	"time"
)

func AllCombo(arr []int) [][]int {
	var all [][]int
	poss := Factorial(len(arr))
	for i := 0; i < poss; i++ {
		j := rand.Intn(len(arr))+1
		arr[i], arr[j] = arr[j], arr[i]
		if !InArrArr(all,arr) {
			all = append(all,arr)
		}
	}
	return all
}

func Factorial(num int) int {
	if num != 0 {
		return num * Factorial(num-1)
	}
	return 1
}

func Shuffle(arr []int) {
	rand.Seed(int64(time.Now().Second()))
	
	return arr
}