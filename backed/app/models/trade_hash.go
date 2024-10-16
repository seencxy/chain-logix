package models

import "gorm.io/gorm"

type TradeHash struct {
	gorm.Model
	PackageId       string `json:"packageId"`
	TransactionHash string `json:"transactionHash"`
}
