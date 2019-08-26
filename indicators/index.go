package indicators

// Indicator is ...
type Indicator struct {
	Name string
	Inputs []string
	Outputs []string
	Fn func(inputs []float64) [][]float64
}

// Index is a list of indicator definitions
var Index = []Indicator{

	Indicator {
		Name: "Horizontal Line",
		Inputs: []string{ "value" },
		Outputs: []string{ "value" },
		Fn: func(inputs []float64) [][]float64 {
			return [][]float64{[]float64{}}
		},
	},
	Indicator {
		Name: "Stochastics",
		Inputs: []string{ "k", "d", "Smooth"},
		Outputs: []string{
			"percentK",
			"percentD",
		},
		Fn: func(inputs []float64) [][]float64 {
			return [][]float64{[]float64{}}
		},
	},
}
