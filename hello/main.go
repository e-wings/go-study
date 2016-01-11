package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3}
	for i, v := range arr {
		//如果要修改原始数组，这样写是无效的
		v = v + 1
		//这样写是可以的
		arr[i] = arr[i] + 1
	}

	fmt.Println(arr)
}
