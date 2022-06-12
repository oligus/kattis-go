package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	var line string
	fmt.Scanln(&n)

	scanner := bufio.NewScanner(os.Stdin)

	i := 0
	for scanner.Scan() {
		line = scanner.Text()
		for _, str := range strings.Fields(line) {
			x, _ := strconv.Atoi(str)
			if x < 0 {
				i = i + 1
			}
		}
	}

	fmt.Println(i)
}
