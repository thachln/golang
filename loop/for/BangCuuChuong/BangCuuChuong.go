package main

import "fmt"

func main() {
	fmt.Println("Bảng Cửu Chương.")

	for i := 2; i <= 9; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Print(i, " x ", j, " = ", i*j)
			fmt.Print("   ")
		}
		fmt.Println()
	}
}
