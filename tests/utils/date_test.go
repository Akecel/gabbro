package tests

import (
    "testing"

    "github.com/akecel/gabbro/utils"
)

func TestParseTimeStampToString(t *testing.T) {
    timestamp := 1559001600
    result := utils.ParseTimeStampToString(timestamp)

    if result != "2019-05-28" {
        t.Errorf("Result=%s; Expected=%s", result, "2019-05-28")
    }
}