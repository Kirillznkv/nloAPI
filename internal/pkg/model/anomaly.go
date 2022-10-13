package model

import (
	"time"
)

type Anomaly struct {
	SessionId string
	Frequency float64
	Timestamp time.Time
}
