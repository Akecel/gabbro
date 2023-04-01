package utils

import (
	"time"
)

func ParseTimeStampToString(timestamp int, onlyDate bool) string {
	i := int64(timestamp)
	tm := time.Unix(i, 0)

	var format string
	if onlyDate {
		format = "2006"
	} else {
		format = "January 2, 2006"
	}

	return tm.Format(format)
}
