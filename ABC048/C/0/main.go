package main

import "fmt"

func main() {
	var N, x int
	fmt.Scan(&N, &x)

	a := make([]int, N)
	b := make([]int, N-1)
	for i := range a {
		fmt.Scan(&a[i])
		if i >= 1 {
			b[i-1] = a[i-1] + a[i]
		}
	}

	c := 0
	for i := range b {

		n := b[i] - x

		if n <= 0 {
			continue
		}

		c += n
		if i == len(b)-1 {
			break
		}

		var r int
		n -= a[i+1]
		if n > 0 {
			n -= a[i]
		} else if n < 0 {
			r = -n
		} else {
			r = 0
		}
		b[i+1] -= a[i+1] - r
		a[i+1] = r
	}
	fmt.Println(c)
}
