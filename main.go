package main

import (
	"fmt"
	"log"

	"github.com/aarontanjaya/finance-data-collector/collector"
	"github.com/aarontanjaya/finance-data-collector/entity"
	util "github.com/aarontanjaya/finance-data-collector/utils"
)

func main() {
	c, err := collector.NewPasardanaSnaphotCollector("https://www.pasardana.id/api/")
	if err != nil {
		log.Fatalln(err)
	}
	q := entity.PasardanaPagination{
		Start:     1,
		Length:    100,
		SortField: "Name",
		SortOrder: "ASC",
	}
	res, err := c.GetAll(q)
	fmt.Printf("result %v %s", util.PrettifyObject(res), err)
}
