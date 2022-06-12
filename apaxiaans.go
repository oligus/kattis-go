package main

import "fmt"

func main() {
	var name string
	var fixed = make([]byte, 0)
	fmt.Scanln(&name)

	fixed = append(fixed, name[0])
	for i := 0; i < len(name); i++ {
		if i > 0 && name[i-1] != name[i] {
			fixed = append(fixed, name[i])
		}
	}
	str := string(fixed[:])
	fmt.Println(str)
}
