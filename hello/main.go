package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a float32 = 3.0
	b := float64(a)
	fmt.Println(reflect.TypeOf(b))
}
