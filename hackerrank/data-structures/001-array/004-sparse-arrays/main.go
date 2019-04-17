package main

import (
	"fmt"
)

func main() {
	maps := make(map[string]int)

	var n, q int
	var str string

	fmt.Scanf("%d\n", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s\n", &str)
		maps[str]++
	}

	fmt.Scanf("%d\n", &q)
	for i := 0; i < q; i++ {
		fmt.Scanf("%s\n", &str)
		fmt.Println(maps[str])
	}
}
