package main

import (
	"fmt"
	"sync"

	"frontendmasters.com/go/crypto/api"
)

func main() {
	currencies := []string{"BTC", "ETH", "BCH"}

	var wg sync.WaitGroup

	for _, curreny := range currencies {
		wg.Add(1)
		go func() {
			getCurrencyData(curreny)
			wg.Done()
		}()
	}
	wg.Wait()
}

func getCurrencyData(currency string) {
	rate, err := api.GetRate(currency)
	if err == nil {
		fmt.Printf("The rate for %v is %.2f\n", rate.Currency, rate.Price)
	}
}
