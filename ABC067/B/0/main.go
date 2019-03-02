package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func read() string {
	sc.Scan()
	return sc.Text()
}
func main() {
	sc.Split(bufio.ScanWords)
	N, _ := strconv.Atoi(read())
	K, _ := strconv.Atoi(read())

	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i], _ = strconv.Atoi(read())
	}

	sort.Sort(sort.IntSlice(a))

	var sum int
	for _, x := range a[len(a)-K:] {
		sum += x
	}
	fmt.Println(sum)
}
