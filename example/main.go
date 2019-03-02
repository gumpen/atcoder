package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	c := math.Pow(2, 3)
	fmt.Println(reflect.TypeOf(c))
	fmt.Printf("%#v\n", c)
}
