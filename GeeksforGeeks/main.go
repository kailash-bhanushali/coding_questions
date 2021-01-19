package main

import (
	"fmt"

	"./rotationArray"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	d := 3
	fmt.Println(rotationArray.Rotation_Array(arr, d))
}
