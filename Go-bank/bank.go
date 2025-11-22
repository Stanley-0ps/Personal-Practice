package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"practice.com/bank/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("_______")
		//panic("sorry, can't continue.")
	}
	fmt.Println("welcome to Go bank!")
	fmt.Println("Contact us 24/7:", randomdata.PhoneNumber())
	for { //i := 0; i < 20; i++
		presentOptions()

		fmt.Print("input your choice: ")
		var choice int
		fmt.Scan(&choice)
		fmt.Println("your choice:", choice)
		//wantsCheckBalance := choice == 1

		switch choice {
		case 1:
			fmt.Println("your balance is:", accountBalance)
		case 2:
			fmt.Print("your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("invalid deposit amount, deposit amount must be greater than 0.")
				//return
				continue
			}
			accountBalance += depositAmount // thesame as accountBalance = accountbalance + depositAmount
			fmt.Println("New account balance!:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			fmt.Print("withdrawal Amount:")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount > accountBalance {
				fmt.Println("invalid withdrawal amount, you can't withdraw more than you have!")
				continue
			}
			accountBalance -= withdrawalAmount
			fmt.Println("updated account balance!:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for Banking with Us")
			return
			// break
		}
	}
}
