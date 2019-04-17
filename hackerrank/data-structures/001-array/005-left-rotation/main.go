package main

import "fmt"

func main() {
	var n, d int
	fmt.Scanf("%d %d\n", &n, &d)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		idx := (i + n - d) % n
		fmt.Scanf("%d", &arr[idx])
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}
}
