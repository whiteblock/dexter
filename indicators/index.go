package indicators

import (
	"github.com/whiteblock/dexter"
	//"github.com/sdcoffey/techan"
)

// Indicator is a struct that describes the inputs and outputs of an Indicator and its calculation function.
type Indicator struct {
	Name string
	Inputs []string // Ideally, this would be a struct with all the params that the Pine Script input() function takes.
	// https://www.tradingview.com/pine-script-reference/#fun_input
	Outputs []string
	Fn func(inputs []float64, candles []dexter.Candle) [][]float64
}

// Index is a list of indicator definitions
var Index = []Indicator{

	Indicator {
		Name: "Horizontal Line",
		Inputs: []string{ "value" },
		Outputs: []string{ "value" },
		Fn: func(inputs []float64, candles []dexter.Candle) [][]float64 {
			var result [][]float64
			value := inputs[0]
			for i := 0; i < len(candles); i++ {
				result = append(result, []float64{ value })
			}
			return result
		},
	},

	Indicator {
		Name: "Moving Average",
		Inputs: []string { "period" },
		Outputs: []string { "value" },
		Fn: func(inputs []float64, candles []dexter.Candle) [][]float64 {
			var result [][]float64
			return result
		},
	},

	Indicator {
		Name: "Stochastics",
		Inputs: []string{ "k", "d", "smooth" },
		Outputs: []string{
			"percentK",
			"percentD",
		},
		Fn: func(inputs []float64, candles []dexter.Candle) [][]float64 {
			/*
			k := inputs[0]
			d := inputs[1]
			smooth := inputs[2] */
			return [][]float64{[]float64{}}
		},
	},

}
