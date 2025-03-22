package main

import (
	"fmt"
	"math"
)

func main() {
	var investedAmount, expectedReturnRate, years float64

	fmt.Println("Enter invested amount: ")
	fmt.Scan(&investedAmount)
	if investedAmount <= 0.0 {
		fmt.Println("Invalid invested amount")
		return
	}

	fmt.Println("Enter expected return rate: ")
	fmt.Scan(&expectedReturnRate)
	if expectedReturnRate <= 0.0 {
		fmt.Println("Invalid return rate")
		return
	}

	fmt.Println("Enter years: ")
	fmt.Scan(&years)
	if years <= 0.0 {
		fmt.Println("Invalid years")
		return
	}

	var expectedAmount float64 = investedAmount * math.Pow(1+(expectedReturnRate/100), years)
	fmt.Printf("Expected return amount: %.2f", expectedAmount)
}
