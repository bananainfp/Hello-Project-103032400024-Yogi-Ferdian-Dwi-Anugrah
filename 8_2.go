package main

import "fmt"

const NMAX int = 10

func main() {
	var A [NMAX]int
	var n, min int

	fmt.Scan(&n)

	if n > NMAX {
		n = NMAX
	}

	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}

	min = 0

	for i := 1; i < n; i++ {
		if A[i] < A[min] {
			min = i
		}
	}

	fmt.Printf("%d %d\n", A[min], min)
}
