package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var line string
	set := []int{1, 1, 2, 2, 2, 8}
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line = scanner.Text()
		for i, str := range strings.Fields(line) {
			x, _ := strconv.Atoi(str)
			set[i] = set[i] - x
		}
	}

	fmt.Println(strings.Trim(fmt.Sprint(set), "[]"))
}
