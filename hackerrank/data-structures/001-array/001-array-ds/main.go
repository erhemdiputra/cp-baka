package main

import "fmt"

func main() {
	var n, num int
	var arr []int
	fmt.Scanf("%d", &n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &num)
		arr = append(arr, num)
	}

	for i := n - 1; i >= 0; i-- {
		fmt.Printf("%d ", arr[i])
	}

	fmt.Println()
}
