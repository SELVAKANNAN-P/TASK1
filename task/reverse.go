package main

import "fmt"

func main() {
	var tring string
	tring = "ABCDEFGHIJKL"

	s := []byte(tring)
	for i := 2; i < len(s); i += 4 {
		j := i + 1
		s[i], s[j] = s[j], s[i]

	}
	fmt.Println(string(s))

}
