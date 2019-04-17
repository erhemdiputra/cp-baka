package main

import (
	"fmt"
	"math"
)

func main() {
	var n, m, x int
	fmt.Scanf("%d %d\n", &n, &m)

	arr := make([]int, n+5)

	for i := 0; i < m; i++ {
		var a, b, k int
		fmt.Scanf("%d %d %d\n", &a, &b, &k)

		arr[a] += k
		if b+1 <= n {
			arr[b+1] -= k
		}
	}

	max := math.MinInt32
	for i := 1; i < len(arr); i++ {
		x += arr[i]
		if x > max {
			max = x
		}
	}

	fmt.Println(max)
}
