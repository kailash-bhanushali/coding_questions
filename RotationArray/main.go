package main

import "fmt"

func Rotation_Array(arr []int, d int) []int {
	modFind := len(arr) / d
	for i := 0; i < modFind-1; i++ {
		temp := arr[i]
		j := i
		counter := 0
		for counter < modFind-1 {
			arr[j] = arr[j+modFind-1]
			j += modFind - 1
			counter++
		}
		arr[j] = temp
	}
	return arr
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	d := 3
	fmt.Println(Rotation_Array(arr, d))
}
