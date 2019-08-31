package dexter

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"testing"
	"time"
	"github.com/davecgh/go-spew/spew"
	//"github.com/davecgh/go-spew/spew"
)

func LoadFixture(filename string) []Candle {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	data := []Candle{}
	_ = json.Unmarshal([]byte(file), &data)
	return data
}

func TestPrice(t *testing.T) {
	candles := LoadFixture("./fixtures/20190827.binance.btcusdt.1h.json")
	chart := Chart{
		Exchange: "binance",
		Market: "BTC/USDT",
		Timeframe: "1h",
		Candles: candles,
	}
	price, _ := FindIndicatorByName("Price")
	output := price.Fn([]float64{ 10000 }, chart)
	if len(output) != 500 {
		t.Errorf("There should be 500 values in output")
	}
	if output[0][0] != 11674.38  {
		t.Errorf("The 0th price should be 11674.38 but was %f", output[0][0])
	}
	if output[499][0] != 10145.82  {
		t.Errorf("The last price should be 10145.82 but was %f", output[499][0])
	}
}

func TestHorizontalLine(t *testing.T) {
	candles := LoadFixture("./fixtures/20190827.binance.btcusdt.1h.json")
	chart := Chart{
		Exchange: "binance",
		Market: "BTC/USDT",
		Timeframe: "1h",
		Candles: candles,
	}
	horizontalLine, _ := FindIndicatorByName("Horizontal Line")
	output := horizontalLine.Fn([]float64{ 10000 }, chart)
	if len(output) != 500 {
		t.Errorf("There should be 500 values in output")
	}
	if output[0][0] != 10000 {
		t.Errorf("The first value should be 10000")
	}
	if output[499][0] != 10000 {
		t.Errorf("The last value should be 10000")
	}
}

func TestMovingAverage(t *testing.T) {
	candles := LoadFixture("./fixtures/20190827.binance.btcusdt.1h.json")
	chart := Chart{
		Exchange: "binance",
		Market: "BTC/USDT",
		Timeframe: "1h",
		Candles: candles,
	}
	movingAverage, _ := FindIndicatorByName("Moving Average")
	output := movingAverage.Fn([]float64{ 10 }, chart)
	if len(output) != 500 {
		t.Errorf("There should be 490 values in output")
	}
}

func TestTimestamp(t *testing.T) {
	// ccxt gives me time back in milliseconds
	ts := time.Unix(1566946800000 / 1000, 0)
	if ts.Unix() != 1566946800 {
		spew.Dump(ts, ts.Unix())
		t.Errorf("No good.")
	}
}
