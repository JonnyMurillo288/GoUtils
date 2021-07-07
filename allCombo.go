package utils

import (
	"math/rand"
	"time"
	"fmt"
)

func AllCombo(arr []int) [][]int {

	var all [][]int
	poss := Factorial(len(arr))
	for i := 0; i < poss; i++ {
		rand.Seed(int64(time.Now().Second()))
		j := rand.Intn(len(arr)-1)
		l := rand.Intn(len(arr)-1)
		fmt.Println(j,l)
		arr[l], arr[j] = arr[j], arr[l]
		if !InIntArrArr(all,arr) {
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

