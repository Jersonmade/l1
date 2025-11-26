package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	a.SetInt64(1_257_789)
	b := new(big.Int)
	b.SetInt64(8_768_123)

	fmt.Println("произведение", new(big.Int).Mul(a, b))
	fmt.Println("сумма", new(big.Int).Add(a, b))
	fmt.Println("деление", new(big.Int).Div(b, a))
	fmt.Println("разница", new(big.Int).Sub(b, a))
}