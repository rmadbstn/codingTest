package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var T int
	fmt.Fscanln(reader, &T)

	for i := 0; i < T; i++ {
		var str string
		var N int
		fmt.Fscanln(reader, &N, &str)

		var newStr string
		for j := 0; j < len(str); j++ {
			for k := 0; k < N; k++ {
				newStr += string(str[j])
			}
		}

		fmt.Fprintln(writer, newStr)
	}

}
