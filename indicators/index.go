package indicators

import (
	"fmt"
	"log"
	"time"
	"github.com/whiteblock/dexter"
	"github.com/sdcoffey/techan"
	"github.com/sdcoffey/big"
)

// Indicator is a struct that describes the inputs and outputs of an Indicator and its calculation function.
type Indicator struct {
	Name string
	Inputs []string // Ideally, this would be a struct with all the params that the Pine Script input() function takes.
	// https://www.tradingview.com/pine-script-reference/#fun_input
	Outputs []string
	Fn func(inputs []float64, chart dexter.Chart) [][]float64
}

// Index is a list of indicator definitions
var Index = []Indicator{

	Indicator {
		Name: "Horizontal Line",
		Inputs: []string{ "value" },
		Outputs: []string{ "value" },
		Fn: func(inputs []float64, chart dexter.Chart) [][]float64 {
			var result [][]float64
			value := inputs[0]
			for i := 0; i < len(chart.Candles); i++ {
				result = append(result, []float64{ value })
			}
			return result
		},
	},

	Indicator {
		Name: "Moving Average",
		Inputs: []string { "period" },
		Outputs: []string { "value" },
		Fn: func(inputs []float64, chart dexter.Chart) [][]float64 {
			var result [][]float64
			period := int(inputs[0])
			series := SeriesFromChart(chart)
			closePrices := techan.NewClosePriceIndicator(series)
			movingAverage := techan.NewSimpleMovingAverage(closePrices, period)
			res := movingAverage.Calculate(0)
			fmt.Println(res)
			return result
		},
	},

}

// SeriesFromChart will create a techan.TimeSeries from a dexter.Chart
func SeriesFromChart(chart dexter.Chart) *techan.TimeSeries {
	series := techan.NewTimeSeries()
	duration, err := time.ParseDuration(chart.Timeframe)
	if err != nil {
		log.Fatal(err)
	}
	for _, c := range chart.Candles {
		period := techan.NewTimePeriod(time.Unix(int64(c.Timestamp), 0), duration)
		candle := techan.NewCandle(period)
		candle.OpenPrice = big.NewDecimal(c.O)
		candle.MaxPrice = big.NewDecimal(c.H)
		candle.MinPrice = big.NewDecimal(c.L)
		candle.ClosePrice = big.NewDecimal(c.C)
		series.AddCandle(candle)
	}
	return series
}

