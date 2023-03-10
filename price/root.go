/*
Copyright © 2023 Gunjan Patel
*/
package price

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"golang.org/x/exp/slices"
)

// Prices struct which contains
// an array of prices
type PricesByDate struct {
	Date   time.Time `json:"PriceDate"`
	Prices []Price   `json:"DisplayPrices"`
}

type Price struct {
	Hour  string  `json:"time"`
	Value float32 `json:"value"`
}

func GetPrice(FilterDate ...time.Time) PricesByDate {
	// Open our jsonFile
	jsonFile, err := os.Open("data/prices.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	fileData, _ := ioutil.ReadAll(jsonFile)

	var pricesByDate []PricesByDate

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(fileData, &pricesByDate)

	idx := slices.IndexFunc(pricesByDate, func(c PricesByDate) bool { return c.Date.Format("2006-01-02") == FilterDate[0].Format("2006-01-02") })

	return pricesByDate[idx]
}
