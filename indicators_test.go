package dexter

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"testing"
	"github.com/davecgh/go-spew/spew"
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
	spew.Dump(candles)
}

func TestMovingAverage(t *testing.T) {

}
