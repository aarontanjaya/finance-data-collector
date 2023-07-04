package main

import "github.com/aarontanjaya/finance-data-collector/collector"

func main() {
	c := collector.NewPasardanaSnaphotCollector("https://www.pasardana.id/api/")

	c.GetAll()
}
