package dexter

import (
	"context"
	"fmt"
	"log"
	"time"
	dataPb "github.com/whiteblock/dexter/api/data"
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

// Create a key for an exchange,market,timeframe triplet from an alert.
func makeChartKey(alert Alert) string {
	key := fmt.Sprintf("%s,%s,%s", alert.Exchange, alert.Market, alert.Timeframe)
	return key
}

// SetupChart returns a chart instance for the given exchange, market and timeframe.
func SetupChart(alert Alert, client dataPb.DataClient) Chart {
	key := makeChartKey(alert)
	var chart Chart
	chart, ok := Charts[key]
	if !ok {
		chart = Chart{}
		Charts[key] = chart
		chart.InitializeCandles(client)
		chart.StreamCandles(client)
	}
	return chart
}

// InitializeCandles uses a dexter-data client to load an initial set of candles for this chart.
func (chart Chart) InitializeCandles(client dataPb.DataClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := client.GetCandles(ctx, &dataPb.CandlesRequest{
		Exchange: chart.Exchange,
		Market: chart.Market,
		Timeframe: chart.Timeframe,
	})
	if err != nil {
		log.Fatalln("Error", err)
	}
	for i, c := range res.Candles {
		chart.Candles = append(chart.Candles, Candle{c.Timestamp, c.O, c.H. c.L, c.C, c.V})
	}
}

// StreamCandles starts starts 
func (chart Chart) StreamCandles(client dataPb.DataClient) {
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
