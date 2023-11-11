package main

import (
	"fmt"
	"math"
)

func main() {
	var a uint16
	var b uint16
	fmt.Printf("Введите a: ")
	fmt.Scanf("%f\n", &a)

	fmt.Printf("Введите b: ")
	fmt.Scanf("%f\n", &b)

	var c float64 = math.Sqrt(float64(a)*float64(a) + float64(b)*float64(b))

	fmt.Println("Гипотенуза равна:", c)
}
