package main

import (
	"fmt"
)

func main() {
	// var investmentAmount = 1000
	// var inflationRate = 5.5
	// var years = 10

	// fmt.Println("Final value of the investment is: ", float64(investmentAmount)*(1+inflationRate/100)*float64(years))

	// explicitly telling go the datatype.
	var investmentAmount float64 = 1000
	var years float64 = 10

	var inflationRate float64
	// Can also be initialized as
	// inflationRate := 5.5

	// Can also be multiple initialized in same line as
	// investmentAmount, years, inflationRate := 1000.0, 10.0, 5.5

	// constants
	const expectedReturnRate float64 = 8

	fmt.Print("Enter inflation rate: ")
	fmt.Scan(&inflationRate)

	// printf for formatted string output.
	fmt.Printf("For inflation rate: %v which is of type %T\n", inflationRate, inflationRate)

	// Controlling digits after decimal.
	fomrattedStringOutput := fmt.Sprintf("For expectedReturnRate rate: %.2f\n", expectedReturnRate)

	fmt.Print(fomrattedStringOutput)

	fmt.Println("Final value of the investment is: ", calculateFinalValue(expectedReturnRate, inflationRate, investmentAmount, years))

	valueA, valueB := returnMultipleValues()

	fmt.Printf("Two values are %.1f, %v\n", valueA, valueB)

	valueANamed, valueBNamed := returnMultipleValuesNamed()

	fmt.Printf("Two values are %.1f, %v\n", valueANamed, valueBNamed)

}

func calculateFinalValue(returnRate float64, inflationRate float64, amt float64, years float64) float64 {
	return (amt * ((1 + returnRate - inflationRate) / 100) * years)
}

func returnMultipleValues() (float64, int64) {
	return 1.0, 987234786
}

func returnMultipleValuesNamed() (a float64, b int64) {
	a, b = 2.0, 987234783236
	return
}
