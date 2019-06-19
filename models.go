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
	Alerts []Alert
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
	Timeframe string
	LineA string
	Condition AlertCondition
	LineB string
	Frequency NotificationFrequency
	Message string
	Webhook Webhook
}

// Indicator - a technical analysis function
type Indicator struct {
	gorm.Model
	Name string
	Implementation string
	Source string
	Inputs postgres.Jsonb
	Lines postgres.Jsonb
	Styles postgres.Jsonb
}

// Webhook - a URL to request to when an Alert is triggered.
type Webhook struct {
	gorm.Model
	Method string
	URL string
	Body string
}

