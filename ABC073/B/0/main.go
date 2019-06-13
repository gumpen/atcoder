package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
func main() {
	sc.Split(bufio.ScanWords)
	N := nextInt()
	sum := 0
	for i := 0; i < N; i++ {
		l, r := nextInt(), nextInt()
		sum += r - l + 1
	}
	fmt.Println(sum)
}
