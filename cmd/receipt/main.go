package main

import "github.com/brianbroderick/receipts/internal/receipt"

func main() {
	// user := "me"
	user := "brian@thelensflare.com"

	receipt.Execute(user)
}
