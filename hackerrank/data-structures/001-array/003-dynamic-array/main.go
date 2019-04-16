package main

import "fmt"

func main() {
	var n, q, kind int64
	var arr [100000][]int64
	var x, y, lastAnswer int64

	fmt.Scanf("%d %d\n", &n, &q)

	for i := int64(0); i < q; i++ {
		fmt.Scanf("%d %d %d\n", &kind, &x, &y)

		idx := (x ^ lastAnswer) % n
		if kind == 1 {
			arr[idx] = append(arr[idx], y)
		} else if kind == 2 {
			elementIdx := y % int64(len(arr[idx]))
			lastAnswer = arr[idx][elementIdx]
			fmt.Println(lastAnswer)
		}
	}
}
