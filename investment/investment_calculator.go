package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {

	var investmentAmount float64
	var years float64
	var expectedReturnRate = 5.5

	//fmt.Print("investment Amount: ")
	GetUserInput("investment Amount: ")
	fmt.Scan(&investmentAmount)

	//fmt.Print("years: ")
	GetUserInput("years: ")
	fmt.Scan(&years)

	//fmt.Print("expected Return Rate: ")
	GetUserInput("expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := CalculatedValues(investmentAmount, expectedReturnRate, years)
	//futureValue := (investmentAmount) * math.Pow(1+expectedReturnRate/100, years)
	//futureRealValue := (futureValue) / math.Pow(1+inflationRate/100, years)

	FormattedFV := fmt.Sprintf("future Value: %.2f\n", futureValue)
	FormattedRFV := fmt.Sprintf("Future Value (adjusted for inflation): %.2f\n", futureRealValue)
	//fmt.Println("Future Value:", futureValue)
	//fmt.Printf("future Value: %.2f\nFuture Value (adjusted for inflation): %.2f\n", futureValue, futureRealValue)
	//fmt.Println("Future Value (adjusted for inflation):", futureRealValue)
	fmt.Print(FormattedFV, FormattedRFV)
}

func GetUserInput(text string) {
	fmt.Print(text)
}

func CalculatedValues(investmentAmount, expectedReturnRate, years float64) (fv float64, frv float64) {
	fv = (investmentAmount) * math.Pow(1+expectedReturnRate/100, years)
	frv = (fv) / math.Pow(1+inflationRate/100, years)
	//return fv, frv
	return
}
