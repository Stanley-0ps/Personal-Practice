package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	var expextedReturnRate = 5.5

	fmt.Print("investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("years: ")
	fmt.Scan(&years)

	fmt.Print("expected Return Rate: ")
	fmt.Scan(&expextedReturnRate)

	futureValue := (investmentAmount) * math.Pow(1+expextedReturnRate/100, years)
	futureRealValue := (investmentAmount) / math.Pow(1+inflationRate/100, years)

	FormattedFV := fmt.Sprintf("future Value: %.2f\n", futureValue)
	FormattedRFV := fmt.Sprintf("Future Value (adjusted for inflation): %.2f\n", futureRealValue)
	//fmt.Println("Future Value:", futureValue)
	//fmt.Printf("future Value: %.2f\nFuture Value (adjusted for inflation): %.2f\n", futureValue, futureRealValue)
	//fmt.Println("Future Value (adjusted for inflation):", futureRealValue)
	fmt.Print(FormattedFV, FormattedRFV)
}
