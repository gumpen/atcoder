package main

import "fmt"

func main() {
	var S string
	fmt.Scan(&S)

	rs := []rune(S)
	for i := 2; i < len(rs); i += 2 {
		if checkEvenString(rs[:len(rs)-i]) {
			fmt.Println(len(rs[:len(rs)-i]))
			break
		}
	}
}

func checkEvenString(rs []rune) bool {
	if string(rs[0:(len(rs)/2)]) == string(rs[(len(rs)/2):]) {
		return true
	}
	return false
}
