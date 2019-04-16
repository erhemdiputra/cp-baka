package main

import (
	"fmt"
	"math"
)

func main() {
	const n = 6
	const m = 6
	var arr [n][m]int

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scanf("%d", &arr[i][j])
		}
	}

	var max = math.MinInt32
	var val int
	for i := 0; i < n-2; i++ {
		for j := 0; j < m-2; j++ {
			val = arr[i][j] + arr[i][j+1] + arr[i][j+2] +
				arr[i+1][j+1] +
				arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]

			if val > max {
				max = val
			}
		}
	}

	fmt.Println(max)
}
