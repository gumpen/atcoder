// 参考
// https://atcoder.jp/contests/abc068/submissions/4128693

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	m := 1
	for m*2 <= n {
		m *= 2
	}
	fmt.Println(m)
}
