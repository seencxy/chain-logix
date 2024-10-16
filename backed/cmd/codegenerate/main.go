package main

import (
	"backed/app/models"
	"backed/internal/pkg"
)

func main() {
	generator := pkg.NewGenerator("database")

	generator.ApplyBasic(&models.UserInfo{})
	generator.ApplyBasic(&models.LogisticsInfo{})
	generator.ApplyBasic(&models.TradeHash{})

	generator.Execute()
}
