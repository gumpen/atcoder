package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}
	sort.Ints(a)
	sum := 0
	for _, x := range a[len(a)-k:] {
		sum += x
	}
	fmt.Println(sum)
}
