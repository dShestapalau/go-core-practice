package main

import (
	"fmt"

	"example.com/practice-project/filemanager"
	"example.com/practice-project/prices"
)

func main() {
	var taxRates []float64 = []float64{0, 0.07, 0.1, 0.15}

	doneChans := make([]chan bool, len(taxRates))

	errorChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)

		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%v.json", taxRate))
		// cmd := cmdmanager.New()

		priceJob := prices.New(fm, taxRate)
		go priceJob.Process(doneChans[index], errorChans[index])

	}

	for index, _ := range taxRates {
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case done := <-doneChans[index]:
			fmt.Println(done)
		}
	}
}
