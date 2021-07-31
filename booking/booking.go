package booking

import "fmt"

const vatRate = 20.0

func printVatRate() {
	fmt.Println("%.2f %%", vatRate)
}