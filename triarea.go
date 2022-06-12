package main

import (
	"fmt"
)

func main() {
	var v1, v2 float32
	var res float32
	fmt.Scanln(&v1, &v2)

	res = .5 * (v1 * v2)

	s := fmt.Sprintf("%.7f", res)
	fmt.Println(s)
}
