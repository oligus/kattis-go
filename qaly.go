package main

import (
	"fmt"
)

func main() {
	var v1 int
	var q, l, ql float32
	fmt.Scanln(&v1)

	for i := 0; i < v1; i++ {
		fmt.Scanln(&q, &l)
		ql = ql + q*l
	}

	s := fmt.Sprintf("%.3f", ql)
	fmt.Println(s)
}
