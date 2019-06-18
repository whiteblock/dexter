package dexter

import (
//	"github.com/sdcoffey/techan"
)

// Chart - a context in which analysis happens
type Chart struct {
	Exchange string
	Market string
	Timeframe string
	Alerts []Alert
}

// InitializeCandles - set the intial price data for a chart
func (chart Chart) InitializeCandles(candles [][]int) {
}

// AddAlert - adds an Alert to a chart
func (chart Chart) AddAlert(alert Alert) {
}

// UpdateCandle - Update the price data of a chart.
func (chart Chart) UpdateCandle(candle []int) {
}

// Analyze - Go through every alert set for the chart and check to see if any conditions have been met
func (chart Chart) Analyze() {
}
