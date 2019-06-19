package dexter

import (
//	"github.com/sdcoffey/techan"
)

// Candle [timestamp, open, high, low, close, volume]
type Candle []int

// InitializeCandles - set the intial price data for a chart
func (chart Chart) InitializeCandles(candles []Candle) {
}

// AddAlert - adds an Alert to a chart
func (chart Chart) AddAlert(alert Alert) {
}

// UpdateCandle - Update the price data of a chart.
func (chart Chart) UpdateCandle(candle Candle) {
}

// Analyze - Go through every alert set for the chart and check to see if any conditions have been met
func (chart Chart) Analyze() {
}
