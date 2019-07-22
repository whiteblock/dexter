package dexter

import (
	"fmt"
//	"github.com/sdcoffey/techan"
)

// Candle [timestamp, open, high, low, close, volume]
type Candle []int

// Chart - a chart is used internally when managing dexter-data candlestick streams
type Chart struct {
	Exchange string
	Market string
	Timeframe string
	Candles []Candle
	Alerts []Alert
}

// Charts is a map of Charts keyed by `${exchange},${market},${timeframe}`
var Charts map[string]Chart

// SetupChart returns a chart instance for the given exchange, market and timeframe.
func SetupChart(exchange string, market string, timeframe string) Chart {
	key := fmt.Sprintf("%s,%s,%s", exchange, market, timeframe)
	var chart Chart
	chart, ok := Charts[key]
	if !ok {
		chart = Chart{}
		Charts[key] = chart
	}
	return chart
}

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
