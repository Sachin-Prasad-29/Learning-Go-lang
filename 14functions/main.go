package main

import "fmt"

func main() {
	fmt.Println("FUnction in Go Lang")
	fmt.Println(sum(4, 5))
	fmt.Println(proAdder(3, 3, 3, 3, 12523, 5, 234, 235, 23, 23, 4, 234, 23, 423423, 23423, 56, 234, 234, 23, 4234, 234, 234))

	resOne, resTwo := messAdd(43, 2)
	fmt.Println(resOne, "  "+resTwo)
}

func sum(a int, b int) int {
	return a + b
}

func proAdder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}
	return total
}

func messAdd(a int, b int) (int, string) {
	return a + b, "Message from messAdd function"
}
