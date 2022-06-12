package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	fmt.Scanln(&n)
	list := getList(n)

	for _, r := range reverse(list) {
		fmt.Println(r)
	}

}

func reverse(list []int) []int {
	l := len(list)
	m := len(list) - 1
	result := make([]int, l)

	for i := 0; i < l; i++ {
		result[i] = list[m]
		m--
	}

	return result
}

func getList(n int) []int {
	list := make([]int, n)
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < n; i++ {
		scanner.Scan()
		list[i], _ = strconv.Atoi(scanner.Text())
	}

	return list
}
