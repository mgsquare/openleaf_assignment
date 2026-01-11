package tracking

import "time"

type TrackingEvent struct {
	TrackingID string
	Status     string
	Timestamp  time.Time
}
