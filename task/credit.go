package main

import "fmt"

func nodigit(a int) (int, []int) {
	n := 0
	var b []int
	for i := 0; a != 0; i++ {
		k := a % 10
		n++
		a = a / 10
		b = append(b, k)
	}
	return n, b
}
func check(n int, a []int) {
	sum := 0
	d := 0
	for i := 1; i < n; i += 2 {
		a[i] = a[i] * 2
		if a[i] > 9 {
			a[i] = a[i] - 9

		}
	}
	for j := 0; j < n; j++ {
		d = d*10 + a[j]
		sum = sum + a[j]

	}
	if sum%10 == 0 {
		fmt.Println("YOUR CARD IS VALID")

	} else {
		fmt.Println("YOUR CARD IS INVALID")

	}

}

func main() {
	var a int

	a = 4556737586899855
	n, k := nodigit(a)

	if n >= 13 && n <= 16 {

		(check(n, k))

	} else {
		fmt.Printf("card is invalid")
	}

}
