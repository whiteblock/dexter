package dexter

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	dataPb "github.com/whiteblock/dexter/api/data"
)

// Candle [timestamp, open, high, low, close, volume]
type Candle struct {
	Timestamp uint64
	O         float64
	H         float64
	L         float64
	C         float64
	V         float64
}

// Chart - a chart is used internally when managing dexter-data candlestick streams
type Chart struct {
	Exchange  string
	Market    string
	Timeframe string
	Candles   []Candle
	Alerts    []Alert
}

// Charts is a map of Charts keyed by `${exchange},${market},${timeframe}`
var Charts = make(map[string]Chart)

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
		chart = Chart{
			Exchange:  alert.Exchange,
			Market:    alert.Market,
			Timeframe: alert.Timeframe,
		}
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
		Exchange:  chart.Exchange,
		Market:    chart.Market,
		Timeframe: chart.Timeframe,
	})
	if err != nil {
		log.Fatalln("Error", err)
	}
	for _, c := range res.Candles {
		chart.Candles = append(chart.Candles, Candle{c.Timestamp, c.O, c.H, c.L, c.C, c.V})
	}
	log.Println("end init")
}

// StreamCandles starts getting realtime candlestick updates and runs analysis on every updated candlestick.
func (chart *Chart) StreamCandles(client dataPb.DataClient) error {
	log.Println("begin streaming")
	ctx, cancel := context.WithCancel(context.Background())
	stream, err := client.StreamCandles(ctx, &dataPb.CandlesRequest{
		Exchange:  chart.Exchange,
		Market:    chart.Market,
		Timeframe: chart.Timeframe,
	})
	if err != nil {
		log.Fatalln("Error", err)
		return errors.New("Could not stream candlesticks")
	}
	go func() {
		log.Println("streaming")
		for {
			candle, err := stream.Recv()
			if err == io.EOF {
				break
				cancel()
			}
			if err != nil {
				log.Fatalln("Streaming error", err)
				cancel()
			}
			chart.UpdateCandle(candle)
			chart.Analyze()
		}
	}()
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
	i, err := chart.FindAlertIndex(alert.ID)
	if err != nil {
		return err
	}
	chart.Alerts[i] = alert
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
	for _, alert := range chart.Alerts {
		hit := alert.Compare(chart)
		if hit {
			go alert.Send()
		}
	}
}

// Compare checks an alerts conditions and returns true if the conditions are met.
func (alert Alert) Compare(chart Chart) bool {
	// So we grab the two lines.
	var lineA Line
	var lineB Line
	json.Unmarshal([]byte(alert.LineA.RawMessage), &lineA)
	json.Unmarshal([]byte(alert.LineB.RawMessage), &lineB)
	indicatorA, _ := FindIndicatorByName(lineA.Name)
	indicatorB, _ := FindIndicatorByName(lineB.Name)
	log.Printf("Comparing %s to %s", indicatorA.Name, indicatorB.Name)
	// Calculate their values.
	outputA := indicatorA.Fn(lineA.Inputs, chart)
	outputB := indicatorB.Fn(lineB.Inputs, chart)
	// Compare based on comparison style
	//   which will require a little bit of state for a few of them.
	// Return true if conditions have been met and return false otherwise.
	az := len(outputA) - 1
	ay := len(outputA) - 2
	bz := len(outputB) - 1
	by := len(outputB) - 2
	var ai, bi int
	ai, err := indicatorA.FindIndexOfOutput(lineA.Output)
	if err != nil {
		log.Fatalf("lineA.Output %s not found", lineA.Output)
	}
	bi, err = indicatorB.FindIndexOfOutput(lineB.Output)
	if err != nil {
		log.Fatalf("lineB.Output %s not found", lineB.Output)
	}
	switch alert.Condition {
	case Crossing:
		if outputA[ay][ai] < outputB[by][bi] {
			if outputA[az][ai] > outputB[bz][bi] {
				return true
			}
		}
		if outputA[ay][ai] > outputB[by][bi] {
			if outputA[az][ai] < outputB[bz][bi] {
				return true
			}
		}
		return false
	case CrossingUp:
		if outputA[ay][ai] < outputB[by][bi] {
			if outputA[az][ai] > outputB[bz][bi] {
				return true
			}
		}
		return false
	case CrossingDown:
		if outputA[ay][ai] > outputB[by][bi] {
			if outputA[az][ai] < outputB[bz][bi] {
				return true
			}
		}
		return false
	case GreaterThan:
		return false
	case LessThan:
		return false
	case EnteringChannel:
		return false
	case ExitingChannel:
		return false
	case InsideChannel:
		return false
	case OutsideChannel:
		return false
	case MovingUp:
		return false
	case MovingDown:
		return false
	case MovingUpPercent:
		return false
	case MovingDownPercent:
		return false
	}
	return false
}

// Send an alert which currently means fire a webhook
func (alert Alert) Send() {
	// grab the HTTP request description
	client := &http.Client{}
	req, err := http.NewRequest(alert.Webhook.Method, alert.Webhook.URL, strings.NewReader(alert.Webhook.Body))
	if err != nil {
		log.Println("couldn't creat request")
	}
	client.Do(req)
}
