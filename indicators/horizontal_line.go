package indicators

// HorizontalLine tests against a single constant value.
type HorizontalLine struct {
	Level float64
}

// Name returns the human readable name of an indicator.
func (*HorizontalLine) Name() string {
	return "Horizontal Line"
}


type IndicatorConfig struct {
	Name string
	Inputs map[string]float
	Outputs array of array of float
}
