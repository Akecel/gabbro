package tests

import (
	"testing"

	"github.com/akecel/gabbro/utils"
)

func TestParseTimeStampToString(t *testing.T) {
    // Test case 1
    input := 1617235200 // April 1, 2021
    expectedOutput := "April 1, 2021"
    result := utils.ParseTimeStampToString(input)
    if result != expectedOutput {
        t.Errorf("ParseTimeStampToString(%d) = %s; expected %s", input, result, expectedOutput)
    }
    
    // Test case 2
    input = 1648713600 // March 31, 2022
    expectedOutput = "March 31, 2022"
    result = utils.ParseTimeStampToString(input)
    if result != expectedOutput {
        t.Errorf("ParseTimeStampToString(%d) = %s; expected %s", input, result, expectedOutput)
    }
    
    // Test case 3
    input = 1661971200 // August 31, 2022
    expectedOutput = "August 31, 2022"
    result = utils.ParseTimeStampToString(input)
    if result != expectedOutput {
        t.Errorf("ParseTimeStampToString(%d) = %s; expected %s", input, result, expectedOutput)
    }
    
    // Test case 4 (edge case)
    input = -62135596800 // December 31, 0000 (the earliest time possible)
    expectedOutput = "December 31, 0000"
    result = utils.ParseTimeStampToString(input)
    if result != expectedOutput {
        t.Errorf("ParseTimeStampToString(%d) = %s; expected %s", input, result, expectedOutput)
    }
}
