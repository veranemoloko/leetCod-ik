package main

import "fmt"

func isHappy(n int) bool {
	mem := make(map[int]bool)

	for n != 1 && !mem[n] {
		mem[n] = true
		n = sumOfSquares(n)
	}

	return n == 1
}

func sumOfSquares(n int) int {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit * digit
		n /= 10
	}
	return sum
}

func main() {
	var num int
	fmt.Scanf("%d", &num)
	fmt.Print(isHappy(num))
}
