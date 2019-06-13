package main

import "fmt"

func main() {
	var w, a, b int
	fmt.Scan(&w, &a, &b)
	if abs(a-b) < w {
		fmt.Println(0)
		return
	}
	fmt.Println(min(abs(b+w-a), abs(a+w-b)))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
