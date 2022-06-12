package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	var n, m int
	fmt.Scanln(&n)
	fmt.Scanln(&m)

	shares := distribute(m, n)

	symbol := "*"
	for i := range shares {
		fmt.Println(strings.Repeat(symbol, shares[i]))

	}

}

func distribute(total int, size int) []int {
	shares := make([]int, size)
	floor := math.Floor(float64(total) / float64(size))

	for i := range shares {
		shares[i] = int(floor)
	}

	remain := total % size

	if remain == 0 {
		return shares
	}

	for i := 0; i < remain; i++ {
		shares[i]++
	}

	return shares
}
