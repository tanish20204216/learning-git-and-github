package main

import "fmt"

func main() {
	fmt.Println("GFG")
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
