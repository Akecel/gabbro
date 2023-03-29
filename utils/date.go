package utils

import (
	"time"
)

func ParseTimeStampToString(timestamp int) string {
	i := int64(timestamp)
	tm := time.Unix(i, 0)
	return tm.Format("2006-01-02")
}
