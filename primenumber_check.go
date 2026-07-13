package main

import "fmt"

func isPrime(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	num := 29

	if isPrime(num) {
		fmt.Println(num, "is Prime")
	} else {
		fmt.Println(num, "is Not Prime")
	}
}
