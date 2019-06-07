package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)

	N, x := nextInt(), nextInt()
	l := nextInt()
	c := 0

	for i := 1; i < N; i++ {
		r := nextInt()
		if n := l + r - x; n > 0 {
			c += n
			if l = r - n; l < 0 {
				l = 0
			}
		} else {
			l = r
		}
	}
	fmt.Println(c)
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
