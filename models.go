package dexter

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

// Chart - a context in which analysis happens
type Chart struct {
	gorm.Model
	Exchange string
	Market string
	Alerts []Alert `gorm:"foreignkey:AlertID"`
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
	ChartID uint
	Timeframe string
	LineA postgres.Jsonb
	Condition AlertCondition
	LineB postgres.Jsonb
	Frequency NotificationFrequency
	Message string
	Webhooks []Webhook
}

// IndicatorLine - the part of the indicator we're using for comparisons
type IndicatorLine struct {
	Indicator string `json:"indicator`
	Inputs []struct {
		Name int `json:"name"`
	}
	Part string `json:"name"`
}

// IndicatorSpec - metadata for a technical analysis function
type IndicatorSpec struct {
	gorm.Model
	Name string
	Implementation string
	Source string
	Inputs postgres.Jsonb
	Lines postgres.Jsonb
	Styles postgres.Jsonb
}

/*
 * In the future, there may also be an Indicator model which would be an instance
 * of an IndicatorSpec with its inputs set.  A Chart would have-many Indicators.
 * This would mimic TradingView's model more closely, but for the first iteration,
 * I don't feel I need to go there.
 */

// Webhook - a URL to request to when an Alert is triggered.
type Webhook struct {
	gorm.Model
	AlertID uint
	Method string
	URL string
	Body string
}

