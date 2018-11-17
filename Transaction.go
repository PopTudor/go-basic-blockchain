package main

type Transaction struct {
	Sender    string  `json:"sender" binding:"required"`
	Recipient string  `json:"recipient" binding:"required"`
	Amount    float32 `json:"amount" binding:"required"`
}
