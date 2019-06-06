package main

import "fmt"

func main() {
	var N, x int
	fmt.Scan(&N, &x)

	// この配列の受け取り方が非常に遅い可能性がある
	a := make([]int, N)
	for i := range a {
		fmt.Scan(&a[i])
	}

	c := 0
	for i := 1; i < N; i++ {
		b := a[i-1] + a[i]
		n := b - x

		if n <= 0 {
			continue
		}

		c += n

		n -= a[i]
		if n < 0 {
			a[i] = -n
		} else {
			a[i] = 0
		}
	}
	fmt.Println(c)
}
