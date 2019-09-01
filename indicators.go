package dexter

import (
	"fmt"
	"log"
	"time"
	"github.com/sdcoffey/big"
	//"github.com/davecgh/go-spew/spew"
	"github.com/sdcoffey/techan"
)

// Indicator is a struct that describes the inputs and outputs of an Indicator and its calculation function.
type Indicator struct {
	Name string
	Inputs []string // Ideally, this would be a struct with all the params that the Pine Script input() function takes.
	// https://www.tradingview.com/pine-script-reference/#fun_input
	Outputs []string
	Fn func(inputs []float64, chart Chart) [][]float64
}

// Indicators is a list of indicator definitions
var Indicators = []Indicator{

	// This makes the current close price of a chart comparable.
	Indicator {
		Name: "Price",
		Inputs: []string{},
		Outputs: []string{ "price" },
		Fn: func(inputs []float64, chart Chart) [][]float64 {
			var result [][]float64
			for i := 0; i < len(chart.Candles); i++ {
				price := chart.Candles[i].C
				result = append(result, []float64{ price })
			}
			return result
		},
	},

	// Did the line (usually price) cross a constant value?
	Indicator {
		Name: "Horizontal Line",
		Inputs: []string{ "value" },
		Outputs: []string{ "value" },
		Fn: func(inputs []float64, chart Chart) [][]float64 {
			var result [][]float64
			value := inputs[0]
			for i := 0; i < len(chart.Candles); i++ {
				result = append(result, []float64{ value })
			}
			return result
		},
	},

	// Simple Moving Average
	Indicator {
		Name: "Moving Average",
		Inputs: []string { "period" },
		Outputs: []string { "value" },
		Fn: func(inputs []float64, chart Chart) [][]float64 {
			var result [][]float64
			period := int(inputs[0])
			series := SeriesFromChart(chart)
			closePrices := techan.NewClosePriceIndicator(series)
			movingAverage := techan.NewSimpleMovingAverage(closePrices, period)
			for i := 0; i < len(chart.Candles); i++ {
				ma := movingAverage.Calculate(i)
				maf := ma.Float()
				//log.Println(maf)
				result = append(result, []float64{ maf })
			}
			return result
		},
	},

}

// SeriesFromChart will create a techan.TimeSeries from a dexter.Chart
func SeriesFromChart(chart Chart) *techan.TimeSeries {
	series := techan.NewTimeSeries()
	duration, err := time.ParseDuration(chart.Timeframe)
	if err != nil {
		log.Fatal(err)
	}
	for _, c := range chart.Candles {
		period := techan.NewTimePeriod(time.Unix(int64(c.Timestamp) / 1000, 0), duration)
		//spew.Dump(c.Timestamp, duration, period)
		candle := techan.NewCandle(period)
		candle.OpenPrice = big.NewDecimal(c.O)
		candle.MaxPrice = big.NewDecimal(c.H)
		candle.MinPrice = big.NewDecimal(c.L)
		candle.ClosePrice = big.NewDecimal(c.C)
		series.AddCandle(candle)
	}
	return series
}

// FindIndicatorByName looks up an indicator by its name
func FindIndicatorByName(name string) (Indicator, error) {
	for _, indicator := range Indicators {
		if indicator.Name == name {
			return indicator, nil
		}
	}
	return Indicator{}, fmt.Errorf("indicator '%s' not found", name)
}

// FindIndexOfOutput lets you know which index of the output array contains the line you're looking for
func (indicator Indicator) FindIndexOfOutput(name string) (int, error) {
	if name == "default" {
		return 0, nil
	}
	i := index(indicator.Outputs, name)
	if i == -1 {
		return -1, fmt.Errorf("line %s not found", name)
	}
	return i, nil
}

// Index finds the index of a string in an array
// https://gobyexample.com/collection-functions
func index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}
