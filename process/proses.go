package process

import (
	"encoding/json"
	"fmt"
	"os"
)

func getBiller() {
	type Person struct {
		Fn string
		Ln string
	}
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
		P      Person `json:"Person"`
	}

	per := Person{Fn: "John",
		Ln: "Doe",
	}

	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
		P:      per,
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)

	type dataListVar struct {
		billerId   string
		name       string
		BillerType string
	}

	dataList := dataListVar{
		billerId:   "45039",
		name:       "PLN - PREPAID",
		BillerType: "Electricity",
	}

	aa, err := json.Marshal(dataList)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(aa)
}
