package dexter

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

/*
// Indicator - metadata for a technical analysis function
type Indicator struct {
	gorm.Model
	Name string
	Implementation string // native or pinescript
	Source string         // pinescript source code
	Inputs postgres.Jsonb // the parameters this indicator takes
        /*
        {
          value: "integer"
        }

        ex:

        { value: 3 }


        alternatively...

        period

        input [ 1 ]
        output [ [ 1.34 ] ]


        horizontal line

        input [ 1.23 ]
        output [ [ 1.234 ] ]

        stochastics

        input [ 1, 2, 3 ] (K, D, Smooth)
        output [
          [3, 4]
          [5, 6]
        ]
        */
        /*
	Lines postgres.Jsonb  // the lines this indicator offers
	Styles postgres.Jsonb // unused for now but anything that's drawn gets to set visual parameters of its own
}

type IndicatorInput []float64
type IndicatorOutput [][]float64

type IndicatorFn func(IndicatorInput) IndicatorOutput
*/

// Line is a line offered by an Indicator for comparison.
type Line struct {
	Name string `json:"name"`
	Inputs []float64 `json:"inputs"`
	Output string `json:"output"`
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
	LineA postgres.Jsonb // Line
        /*
        {
          name: "Horizontal Line",
          inputs: [ 10000 ],
          output: "default"
        }

        {
          name: "Simple Moving Average",
          inputs: [ 10 ], // 10 period MA
          output: "default" // Some indicators have more than one output, but default is a synonym for the 0th index
        }
        */
	Condition AlertCondition
	LineB postgres.Jsonb // Line
	Frequency NotificationFrequency
	MessageBody string
	Webhook Webhook
}

// Webhook - a URL to request to when an Alert is triggered.
type Webhook struct {
	gorm.Model
	AlertID uint
	Method string
	URL string
	Body string
}

/*
drop table webhooks;
drop table indicators;
drop table indicator_specs;
drop table alerts;
drop table charts;
*/
