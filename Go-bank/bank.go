package main

import "fmt"

func main() {
	var accountBalance = 10000.00
	fmt.Println("welcome to Go bank!")

	for { //i := 0; i < 20; i++
		fmt.Println("what do you want to do?")
		fmt.Println("1. check balance")
		fmt.Println("2. deposit money")
		fmt.Println("3. withdraw money")
		fmt.Println("4. exit")

		fmt.Print("input your choice: ")
		var choice int
		fmt.Scan(&choice)
		fmt.Println("your choice:", choice)
		//wantsCheckBalance := choice == 1

switch choice {
case 1: 
fmt.Println("your balance is:", accountBalance)
case 2:
*fmt.Print("your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("invalid deposit amount, deposit amount must be greater than 0.")
				//return
				continue
			}

			accountBalance += depositAmount // thesame as accountBalance = accountbalance + depositAmount
			fmt.Println("updated account balance!:", accountBalance)
		case 3:
fmt.Print("withdrawal Amount:")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount > accountBalance {
				fmt.Println("invalid withdrawal amount, you can't withdraw more than you have!")
				return
			}
			accountBalance -= withdrawalAmount
			fmt.Println("updated account balance!:", accountBalance)
		default: 
fmt.Println("Goodbye!")
break
}
		/*if choice == 1 {
			//fmt.Println("your balance is:", accountBalance)
		} else if choice == 2 {
			fmt.Print("your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("invalid deposit amount, deposit amount must be greater than 0.")
				//return
				continue
			}
			accountBalance += depositAmount // thesame as accountBalance = accountbalance + depositAmount
			fmt.Println("updated account balance!:", accountBalance)
		} else if choice == 3 {
			fmt.Print("withdrawal Amount:")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount > accountBalance {
				fmt.Println("invalid withdrawal amount, you can't withdraw more than you have!")
				return
			}
			accountBalance -= withdrawalAmount
			fmt.Println("updated account balance!:", accountBalance)
		}else {
			fmt.Println("Goodbye!")
			//return
			break
		}
	}*/
	fmt.Println("Thanks for banking with us")
}
