package main

import "fmt"

func main() {
	var N, x int
	fmt.Scan(&N, &x)
	var c, l int
	fmt.Scan(&l)
	for i := 1; i < N; i++ {
		var r int
		fmt.Scan(&r)
		n := l + r - x
		if n < 0 {
			l = r
			continue
		}
		c += n
		l = r - n
		if l < 0 {
			l = 0
		}
	}
	fmt.Println(c)
}
