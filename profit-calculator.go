package main

import (
	"fmt"
	"os"
)

func writeInfoFile(profit float64) {
	profitText := fmt.Sprint(profit)
	os.WriteFile("profit.txt", []byte(profitText), 0644)
}

func main() {
	var income, expenses float64
	const taxRate float64 = 6.5

	fmt.Println("enter your income: ")
	fmt.Scan(&income)
	if income <= 0.0 {
		fmt.Println("Enter valid income")
		return
	}

	fmt.Println("enter your expenses: ")
	fmt.Scan(&expenses)
	if expenses <= 0.0 || expenses > income {
		fmt.Println("Enter valid expenses")
		return
	}

	savings := income - expenses
	fmt.Println("Savings: $", savings)

	eat := savings * (1 - taxRate/100)
	fmt.Println("After tax: $", eat)

	ratio := eat / savings
	fmt.Println("Ratio: ", ratio)

	writeInfoFile(savings)
}
