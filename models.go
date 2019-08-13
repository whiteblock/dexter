package dexter

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

// Indicator - metadata for a technical analysis function
type Indicator struct {
	gorm.Model
	Name string
	Implementation string // native or pinescript
	Source string         // pinescript source code
	Inputs postgres.Jsonb // the parameters this indicator takes
	Lines postgres.Jsonb  // the lines this indicator offers
	Styles postgres.Jsonb // unused for now but anything that's drawn gets to set visual parameters of its own
}

// Line is a line offered by an Indicator for comparison.
type Line struct {
	Name string `json:"name"`
	// What else?
}

/*
Note that indicators that only have one line (like HorizontalLine) don't need to
list any lines in the UI. Maybe in the database, it'll have a Default line to
make it so that I don't need to create a special case for Indicators that have
only one line on the server side.
*/

// Input is a paremter for an indicator
type Input struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Default string `json:"default"`
}

// IndicatorPart - what line in the indicator is being compared?
type IndicatorPart struct {
	IndicatorID int    `json:"indicatorId"`
	Line        string `json:"line"`
	// InputValues
}

// AlertCondition describes how lines can interact with each other.
type AlertCondition int

// The different ways lines can interact with each other
const (
	Crossing AlertCondition = iota + 1
	CrossingUp
	CrossingDown
	GreaterThan
	LessThan
	EnteringChannel
	ExitingChannel
	InsideChannel
	OutsideChannel
	MovingUp
	MovingDown
	MovingUpPercent
	MovingDownPercent
)

// NotificationFrequency - how often should an alert notification fire
type NotificationFrequency int

// The different frequencies of alert notifications
const (
	OnlyOnce NotificationFrequency = iota + 1
	OncePerBar
	OncePerBarClose
	OncePerMinute
)

// Alert - describes market condition that should trigger a notification.
type Alert struct {
	gorm.Model
	Exchange string
	Market string
	Timeframe string
	ExternalID uint64
	LineA postgres.Jsonb // IndicatorPart
	Condition AlertCondition
	LineB postgres.Jsonb // IndicatorPart
	Frequency NotificationFrequency
	MessageBody string
	Webhooks []Webhook
}

// Webhook - a URL to request to when an Alert is triggered.
type Webhook struct {
	gorm.Model
	AlertID uint
	Method string
	URL string
	Body string
}

// Compare checks an alerts conditions and returns true if the conditions are met.
func (alert Alert) Compare(chart Chart) bool {
	return true
}

// Send an alert which currently means fire a webhook
func (alert Alert) Send() {

}

/*
drop table webhooks;
drop table indicators;
drop table indicator_specs;
drop table alerts;
drop table charts;
*/
