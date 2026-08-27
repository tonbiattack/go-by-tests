package timevalue

import "time"

func CompareSameInstant() (directEqual bool, instantEqual bool) {
	utc := time.Date(2026, time.August, 27, 0, 0, 0, 0, time.UTC)
	jst := utc.In(time.FixedZone("JST", 9*60*60))
	return utc == jst, utc.Equal(jst)
}
