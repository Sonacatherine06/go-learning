package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50}

	sum := 0

	for _, value := range numbers {
		sum += value
	}

	fmt.Println("Sum =", sum)
}
