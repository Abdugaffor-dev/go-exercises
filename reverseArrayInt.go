package main

import (
	"fmt"
)

func main() {
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	arr = reverseArray(arr)
	fmt.Println(arr)

}
func reverseArray(arr [5]int) [5]int {
	var reversedArr [5]int
	for index := range arr {
		reversedArr[index] = arr[len(arr)-1-index]
	}
	for index := range reversedArr {
		arr[index] = reversedArr[index]
	}
	return arr
}
