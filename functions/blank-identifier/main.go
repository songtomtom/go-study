package main

import (
	"fmt"
	"math"
)

func getStockPriceChange(prevPrice, currentPrice float64) (float64, float64, int) {
	change := currentPrice - prevPrice
	percentChange := (change / prevPrice) * 100
	x := 100
	return change, percentChange, x
}

func main() {
	prevStockPrice := 80000.0
	currentStockPrice := 120000.0

	change, _, _ := getStockPriceChange(prevStockPrice, currentStockPrice)
	if change < 0 {
		fmt.Printf("The Stock Price decreased by $%.2f\n", math.Abs(change))
	} else {
		fmt.Printf("The Stock Price increased by $%.2f\n", change)
	}
}
