package helpers

import (
	"math/rand"
	"time"
)

func GenerateTrackingID() string {
	return "TRK-" + time.Now().Format("20060102150405") + "-" + string(rune(rand.Intn(1000)))
}

func ExpectedDeliveryDate() time.Time {
	return time.Now().AddDate(0, 0, rand.Intn(5)+2)
}
