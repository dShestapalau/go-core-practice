package main

import (
	"fmt"

	"example.com/practice-project/filemanager"
	"example.com/practice-project/prices"
)

func main() {
	var taxRates []float64 = []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%v.json", taxRate))

		priceJob := prices.New(fm, taxRate)
		priceJob.Process()
	}
}
