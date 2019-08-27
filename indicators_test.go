package dexter

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"testing"
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

func TestHorizontalLine(t *testing.T) {
	candles := LoadFixture("./fixtures/20190827.binance.btcusdt.1h.json")
	chart := Chart{
		Exchange: "binance",
		Market: "BTC/USDT",
		Timeframe: "1h",
		Candles: candles,
	}
	horizontalLine := Indicators[0]
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
	movingAverage := Indicators[1]
	output := movingAverage.Fn([]float64{ 10 }, chart)
	if len(output) != 490 {
		t.Errorf("There should be 490 values in output")
	}
}
