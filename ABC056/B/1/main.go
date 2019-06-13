package main

import "fmt"

func main() {
	var w, a, b int
	fmt.Scan(&w, &a, &b)
	if abs(a-b) < w {
		fmt.Println(0)
		return
	}
	fmt.Println(abs(a-b) - w)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
