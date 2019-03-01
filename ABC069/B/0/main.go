package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Printf("%s%d%s\n", s[:1], len([]rune(s))-2, s[len([]rune(s))-1:])
}
