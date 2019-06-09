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
	N, M := nextInt(), nextInt()
	a := make([]int, M)
	if M > 0 {
		a[0] = nextInt()
		for i := 1; i < M; i++ {
			a[i] = nextInt()
			if a[i]-1 == a[i-1] {
				fmt.Println(0)
				return
			}
		}
	}

	s := 0
	var list []int
	for n := 0; n < M+1; n++ {
		m := 0
		if n == len(a) {
			m = N
		} else {
			m = a[n] - 1
		}
		section := m - s
		c := 1
		for i := 1; i*2 <= section; i++ {
			c += combination(section-i, i)
		}
		s = m + 2
		list = append(list, c)
	}
	fmt.Printf("%#v\n", mul(list))
	return
}

func mul(list []int) int {
	r := 1
	for _, e := range list {
		r = r * e % 1000000007
	}
	return r
}

func permutation(n int, k int) int {
	v := 1
	if 0 < k && k <= n {
		for i := 0; i < k; i++ {
			v *= (n - i)
		}
	} else if k > n {
		v = 0
	}
	return v
}

func factorial(n int) int {
	return permutation(n, n-1)
}

func combination(n int, k int) int {
	return permutation(n, k) / factorial(k)
}
