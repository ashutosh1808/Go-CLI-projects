package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const balanceFileName = "balance.txt"

func writeInfoFile(balance float64) {
	balanceTxt := fmt.Sprint(balance)
	os.WriteFile(balanceFileName, []byte(balanceTxt), 0644)
}

func readFromFile() (float64, error) {
	balanceTxt, err := os.ReadFile(balanceFileName)
	if err != nil {
		return 1000, errors.New("Couldn't read file")
	}

	balanceData := string(balanceTxt)
	balanceAmt, err := strconv.ParseFloat(balanceData, 64)

	if err != nil {
		return 1000, errors.New("Couldn't convert balance")
	}

	return balanceAmt, nil
}

func main() {
	// var balanceAmt float64 = 50000
	balanceAmt, err := readFromFile()
	if err != nil {
		fmt.Println("Error: ")
		panic("Sorry, couldn't get data from file")
	}

	fmt.Println("Welcome to GO bank")
	for {
		fmt.Println("Press: ")
		fmt.Println("1: View Balance")
		fmt.Println("2: Deposit Amount")
		fmt.Println("3: Withdraw Amount")
		fmt.Println("4: Exit")

		var choice int
		fmt.Print("Enter choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			{
				fmt.Println("$", balanceAmt)
			}
		case 2:
			{
				var depositAmt float64
				fmt.Println("Enter amount to deposit: ")
				fmt.Scan(&depositAmt)

				balanceAmt += depositAmt
				fmt.Println("Account credited with $", depositAmt, ". Your new balance is $", balanceAmt)
				writeInfoFile(balanceAmt)
			}
		case 3:
			{
				var withdrawAmt float64
				fmt.Println("Enter amount to withdraw: ")
				fmt.Scan(&withdrawAmt)

				balanceAmt -= withdrawAmt
				fmt.Println("Account debited with $", withdrawAmt, ". Your new balance is $", balanceAmt)
				writeInfoFile(balanceAmt)
			}
		default:
			{
				fmt.Println("Goodbye :)")
				return
			}
		}
	}
}
