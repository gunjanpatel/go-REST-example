/*
Copyright © 2023 Gunjan Patel
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/gunjanpatel/norlys-price/price"
	"github.com/spf13/cobra"
)

// todayCmd represents the today command
var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "Fetch today's lowest, heighest and current price.",
	Long:  `Show today's Norlys price`,
	Run: func(cmd *cobra.Command, args []string) {

		todayPrice := price.GetPrice(time.Now())

		fmt.Println(todayPrice.Date.Format("January 02, 2006"))
		min, max := findMinAndMax(todayPrice.Prices)

		fmt.Println(fmt.Sprintf("%s price is %.2f at time %s", "Lowest", min.Value, min.Hour))
		fmt.Println(fmt.Sprintf("%s price is %.2f at time %s", "Highest", max.Value, max.Hour))
	},
}

func init() {
	rootCmd.AddCommand(todayCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// todayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// todayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func findMinAndMax(price []price.Price) (min price.Price, max price.Price) {
	min = price[0]
	max = price[0]

	for _, p := range price {

		if p.Value < min.Value {
			min = p
		}
		if p.Value > max.Value {
			max = p
		}
	}
	return min, max
}
