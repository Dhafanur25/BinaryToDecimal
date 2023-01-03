package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Binary Number: ")
	fmt.Scan(&x)
	binaryToDecimal(x, &y)
	fmt.Print("Decimal Number: ")
	fmt.Println(y)
}

func squared(a, b int) int {
	var total int
	total = 1
	for i := 1; i <= b; i++ {
		total *= a
	}
	return total
}

func binaryToDecimal(binary int, decimal *int) {
	var x int
	*decimal = 0
	i := 0
	for binary > 0 {
		x = binary % 10
		*decimal += x * squared(2, i)
		i++
		binary = binary / 10
	}
}
