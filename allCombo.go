package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func AllCombo(arr []int) [][]int {

	var all [][]int
	var new []int
	poss := Factorial(len(arr))
	for i := 0; i < poss; i++ {
		rand.Seed(int64(time.Now().Nanosecond()))
		j := rand.Intn(len(arr))
		l := rand.Intn(len(arr))
		arr[l], arr[j] = arr[j], arr[l]
		fmt.Println(all)
		fmt.Println(arr)
		fmt.Print("\n\n")
		if !InIntArrArr(all,arr) {
			new = arr
			all = append(all,new)
		}
		i--
	}
	return all
}

func Factorial(num int) int {
	if num != 0 {
		return num * Factorial(num-1)
	}
	return 1
}

