package indicators

import (
	"github.com/whiteblock/dexter"
)

// Indicator is an interface for all technical analysis algorithms.
type Indicator interface {
	Name() string
	// Update current set of candles
	Update(candles []float64)
}

// Indicators is sn in-memory hashmap of all Indicators known to dexter.
var Indicators map[string]dexter.Indicator

// Register an Indicator with dexter.
func Register(in dexter.Indicator) (error) {
	Indicators[in.Name] = in
	return nil
}

// Compare will test lines against each other
func Compare(a dexter.Line, criteria dexter.AlertCondition, b dexter.Line) bool {
	return false
}
