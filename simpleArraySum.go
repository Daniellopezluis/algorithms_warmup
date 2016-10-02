package main

import "fmt"

func main() {
	size := readSize()
	N := readArray(size)
	sum := arraySum(N)
	fmt.Println(sum)

}

func readSize() int {
	var n int
	fmt.Scanf("%d\n", &n)
	return n
}

func readArray(size int) []int {
	a := make([]int, size)
	for i := 0; i < size; i++ {
		var x int
		fmt.Scan(&x)
		a[i] = x
	}
	return a
}

func arraySum(a []int) int {
	sum := 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	return sum
}
