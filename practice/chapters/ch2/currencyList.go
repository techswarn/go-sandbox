//understanding constants with iota

package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	INR
	EUR
	GBP
	RMB
)

func currencyList() {
	symbol := [...]string{USD: "$", INR: "ru", EUR: "$", GBP: "gbp", RMB: "rmb"}

	fmt.Println(symbol[USD])
}
