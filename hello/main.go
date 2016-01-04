package main

import (
	"fmt"
	"go-tools/form"
)

func main() {
	fmt.Println(form.InRange(1, []int{0, 1, 2, 3}))
	fmt.Println(form.InRange(3.14, []float64{0, 1, 2.1, 3.141}))
	fmt.Println(form.InRange("apple2", []string{"apple", "tree", "sss"}))
}
