package dexter

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"time"
	dataPb "github.com/whiteblock/dexter/api/data"
)

// Candle [timestamp, open, high, low, close, volume]
type Candle struct {
	Timestamp uint64
	O float64
	H float64
	L float64
	C float64
	V float64
}

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
		go chart.StreamCandles(client)
	}
	return chart
}

// InitializeCandles uses a dexter-data client to load an initial set of candles for this chart.
func (chart *Chart) InitializeCandles(client dataPb.DataClient) {
	if len(chart.Candles) > 0 {
		return
	}
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
	for _, c := range res.Candles {
		chart.Candles = append(chart.Candles, Candle{c.Timestamp, c.O, c.H, c.L, c.C, c.V})
	}
}

// StreamCandles starts getting realtime candlestick updates and runs analysis on every updated candlestick.
func (chart *Chart) StreamCandles(client dataPb.DataClient) error {
	if len(chart.Candles) > 0 {
		return nil
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	stream, err := client.StreamCandles(ctx, &dataPb.CandlesRequest{
		Exchange: chart.Exchange,
		Market: chart.Market,
		Timeframe: chart.Timeframe,
	})
	if err != nil {
		log.Fatalln("Error", err)
		return errors.New("Could not stream candlesticks")
	}
	for {
		candle, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Streaming error", err)
		}
		chart.UpdateCandle(candle)
		chart.Analyze()
	}
	return nil
}

// FindAlertIndex - find an alert in a Chart by its database id
func (chart Chart) FindAlertIndex(id uint) (int, error) {
	for i, alert := range chart.Alerts {
		if alert.ID == id {
			return i, nil
			break
		}
	}
	return -1, errors.New("Alert not found")
}

// AddAlert - adds an Alert to a Chart
func (chart *Chart) AddAlert(alert Alert) error {
	exists, err := chart.FindAlertIndex(alert.ID)
	if err != nil {
		return err
	}
	if exists >= 0 {
		chart.Alerts[exists] = alert
	} else {
		chart.Alerts = append(chart.Alerts, alert)
	}
	return nil
}

// UpdateAlert - update an Alert in a Chart
func (chart *Chart) UpdateAlert(alert Alert) error {
	return nil
}

// RemoveAlert - remove an Alert from a Chart
func (chart *Chart) RemoveAlert(alert Alert) error {
	i, err := chart.FindAlertIndex(alert.ID)
	if err != nil {
		return err
	}
	chart.Alerts = append(chart.Alerts[:i], chart.Alerts[i+1:]...)
	return nil
}

// UpdateCandle - Update the price data of a chart.
func (chart *Chart) UpdateCandle(candle *dataPb.Candle) error {
	last := len(chart.Candles) - 1
	lastCandle := chart.Candles[last]
	if lastCandle.Timestamp == candle.Timestamp {
		chart.Candles[last] = Candle{candle.Timestamp, candle.O, candle.H, candle.L, candle.C, candle.V}
	} else {
		chart.Candles = append(chart.Candles, Candle{candle.Timestamp, candle.O, candle.H, candle.L, candle.C, candle.V})
	}
	return nil
}

// Analyze - Go through every alert set for the chart and check to see if any conditions have been met
func (chart Chart) Analyze() {
	// iterate through every alert and check conditions
}
