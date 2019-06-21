package dexter

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

// Chart - a context in which analysis happens
type Chart struct {
	gorm.Model
	ExternalID uint
	Exchange string
	Market string
	Alerts []Alert `gorm:"foreignkey:AlertID"`
	Indicators []Indicator `gorm:foreignkey:ChartID`
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

// Indicator - an instance of IndicatorSpec
type Indicator struct {
	gorm.Model
	IndicatorID uint
	ChartID uint
	Parameters postgres.Jsonb // values for IndicatorSpec.Inputs
}

// IndicatorPart - what line in the indicator is being compared?
type IndicatorPart struct {
	IndicatorID int    `json:"indicator_id"`
	Part        string `json:"part"`
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
	ExternalID uint
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

/*
drop table webhooks;
drop table indicators;
drop table indicator_specs;
drop table alerts;
drop table charts;
*/
