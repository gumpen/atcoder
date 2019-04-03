package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, d float64
	fmt.Scan(&a, &b, &c, &d)
	fmt.Println(math.Max(math.Min(b, d)-math.Max(a, c), 0))
}
