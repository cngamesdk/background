package utils

import (
	"fmt"
	"github.com/spf13/cast"
)

// Percent 百分比
func Percent(numerator, denominator int) float64 {
	if denominator == 0 {
		return 0
	}
	return cast.ToFloat64(numerator) / cast.ToFloat64(denominator)
}

func FloatDecimal2Str(req float64) string {
	if req == 0 {
		return "0.00%"
	}
	return fmt.Sprintf("%.2f", req*100) + "%"
}
