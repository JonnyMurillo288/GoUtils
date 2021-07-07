package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func AllCombo(arr []int) [][]int {

	var all [][]int
	poss := Factorial(len(arr))
	for i := 0; i < poss; i++ {
		rand.Seed(int64(time.Now().Nanosecond()))
		j := rand.Intn(len(arr))
		l := rand.Intn(len(arr))
		new := Switch(arr,j,l)
		fmt.Println("all:",all)
		fmt.Println("new:",new)
		fmt.Println("arr",arr)
		fmt.Println("\n",i)
		if !InIntArrArr(all,new) {
			all = append(all,new)
			fmt.Println("all appended:",all)
			time.Sleep(time.Second*5)
			arr = Switch(arr,j,l)
			continue
		}
		i--
	}
	return all
}

func Switch(arr []int, j int, l int) []int {
	arr[l], arr[j] = arr[j], arr[l]
	return arr
}

func Factorial(num int) int {
	if num != 0 {
		return num * Factorial(num-1)
	}
	return 1
}

