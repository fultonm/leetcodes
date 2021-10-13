package main

import (
	"fmt"
	"golang/helper"
)

func main() {
	DriverMSAC()
}

func Help() {
	println("help")
}

func DriverMSAC() {
	input := helper.RandomArray(1_000_000_000, -200, 200)
	fmt.Printf("Max sum of circular sub array: %v\n", maxSubarraySumCircular(input))
}
