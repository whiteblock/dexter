package dexter

import (
	"github.com/jinzhu/gorm"
)

// AlertAction describes how lines can interact with each other.
type AlertAction int

// The different ways lines can interact with each other
const (
	Crossing AlertAction = iota + 1
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
	LineA string
	Action AlertAction
	LineB string
	Frequency NotificationFrequency
	Message string
	Webhook Webhook
}

// Webhook - a URL to request to when an Alert is triggered.
type Webhook struct {
	gorm.Model
	Method string
	URL string
	Body string
}


// TODO - Create a model for indicators
