package main

import "fmt"

func main() {
	fmt.Println(pow(2, 4))
}

func pow(n int, p int) int {
	r := n
	for i := 1; i < p; i++ {
		r = r * n
	}
	return r
}
