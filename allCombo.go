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


func Permutations(arr []int)[][]int{
    var helper func([]int, int)
    res := [][]int{}

    helper = func(arr []int, n int){
        if n == 1{
            tmp := make([]int, len(arr))
            copy(tmp, arr)
            res = append(res, tmp)
        } else {
            for i := 0; i < n; i++{
                helper(arr, n - 1)
                if n % 2 == 1{
                    tmp := arr[i]
                    arr[i] = arr[n - 1]
                    arr[n - 1] = tmp
                } else {
                    tmp := arr[0]
                    arr[0] = arr[n - 1]
                    arr[n - 1] = tmp
                }
            }
        }
    }
    helper(arr, len(arr))
    return res
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

