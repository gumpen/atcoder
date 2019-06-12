package main

import (
	"bufio"
	"os"
)

// でか文字列
var reader = bufio.NewReader(os.Stdin)

func main() {
	s := nextLine()
}

func nextLine() string {
	var buf, line []byte
	isPrefix := true
	var err error
	for isPrefix {
		buf, isPrefix, err = reader.ReadLine()
		if err != nil {
			panic(err)
		}
		line = append(line, buf...)
	}
	return string(line)

}
