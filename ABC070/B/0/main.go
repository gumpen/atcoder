package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	var ab [100]bool
	for i := a; i < b; i++ {
		ab[i] = true
	}
	cnt := 0
	for i := c; i < d; i++ {
		if ab[i] {
			cnt++
		}
	}
	fmt.Println(cnt)
}
