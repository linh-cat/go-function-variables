package main

import "fmt"

func main() {
	// var revenue float64
	// var expense float64
	// var taxRate float64


	revenue := getUserInput("Revenue: ")
	expense := getUserInput("Expense: ")
	taxRate := getUserInput("Tax Rate: ")
	
	ebt, profit, ratio := calculateFinance(revenue, expense, taxRate)


	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)

}

func calculateFinance(revenue, expense, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expense
	profit := ebt * (1 - taxRate / 100)
	ratio := ebt / profit

	return ebt, profit, ratio
}


func getUserInput(infoText string) float64{
	var valueInput float64

	fmt.Print(infoText)
	fmt.Scan(&valueInput)

	return valueInput
} 