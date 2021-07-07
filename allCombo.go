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
		new = arr
		fmt.Println("all:",all)
		fmt.Println("new:",new)
		fmt.Println("arr",arr)
		if !InIntArrArr(all,new) {
			all = append(all,new)
			fmt.Println("all appended:",all)
			time.Sleep(time.Second*5)
			continue
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

