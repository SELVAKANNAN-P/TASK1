package main

import "fmt"

func main() {
	s1 := "this is go language go go go go go go go go"
	c := 0
	s := []byte(s1)
	for i := 0; i < len(s); i++ {
		j := i + 1
		k := len(s)
		k--
		if s[i] == 'g' && s[i+1] == 'o' && (j == k || s[i+2] == ' ') {
			c += 1

		}

	}
	fmt.Println(c)
}
