package main

import "fmt"

func main() {
	fmt.Println("Loop in Go lang")

	days := []string{"Sunday", "Monday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	// type One
	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// type Two
	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// type Three
	// for index, day := range days {
	// 	fmt.Printf("Index is %v and Value is %v\n", index, day)
	// }

	rougueVal := 1

	for rougueVal < 10 {

		if rougueVal == 2 {
			goto lco
		}

		if rougueVal == 3 {
			continue
		}

		fmt.Println("Value is:", rougueVal)
		rougueVal++
	}

	lco:
		fmt.Println("Jumping a Google.com")

}
