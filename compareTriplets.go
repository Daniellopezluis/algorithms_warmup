package main

import "fmt"

func main() {
	alice := readScore()
	bob := readScore()
	s := compareScores(alice, bob)
	fmt.Printf("%d %d", s[0], s[1])

}

func readScore() []int {
	a := make([]int, 3)
	fmt.Scanf("%d %d %d\n", &a[0], &a[1], &a[2])
	return a
}

func compareScores(a []int, b []int) [2]int {
	s := [2]int{0, 0}
	for i := 0; i < 3; i++ {
		if a[i] > b[i] {
			s[0]++
		} else if a[i] < b[i] {
			s[1]++
		}
	}
	return s
}
