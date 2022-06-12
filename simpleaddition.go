package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a, b string

	fmt.Scanln(&a)
	fmt.Scanln(&b)

	n := new(big.Int)
	m := new(big.Int)
	n, ok1 := n.SetString(a, 10)
	m, ok2 := m.SetString(b, 10)
	if !ok1 || !ok2 {
		fmt.Println("SetString: error")
		return
	}

	sum := big.NewInt(0)
	sum = sum.Add(n, m)

	fmt.Println(sum)
}
