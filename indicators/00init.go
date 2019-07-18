package indicators

import (
	"github.com/whiteblock/dexter"
)

// Indicators is sn in-memory hashmap of all Indicators known to dexter.
var Indicators map[string]dexter.Indicator

// Register an Indicator with dexter.
func Register(in dexter.Indicator) (error) {
	Indicators[in.Name] = in
	return nil
}
