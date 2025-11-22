package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err := GetUserInput("Revenue: ")
	if err != nil {
		fmt.Println(err)
		return
		//panic(err)
	}
	expenses, err := GetUserInput("Expenses: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	taxRate, err := GetUserInput("Tax Rate: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	//if err1 != nil || err2 != nil || err3 != nil {
	//fmt.Println(err1)
	//return
	//}
	ebt, profit, ratio := calculations(revenue, expenses, taxRate)
	fmt.Printf("earningsBeforeTax: %.2f\nearningsAftertax: %.2f\nratio: %.2f\n", ebt, profit, ratio)
	storeResults(ebt, profit, ratio)
}
func GetUserInput(infoText string) (float64, error) {
	var UserInput float64
	fmt.Print(infoText)
	fmt.Scan(&UserInput)

	if UserInput <= 0 {
		return 0, errors.New("value must be a postive number")
	}
	return UserInput, nil
}
func calculations(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := (ebt) * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}
func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.2f\nprofit: %.2f\nratio: %.2f", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}
