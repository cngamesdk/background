package utils

import (
	"fmt"
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/spf13/cast"
	"net/url"
	"strings"
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

// ConvertStructToQueryString 结构体转化请求串
func ConvertStructToQueryString(req interface{}) (resp string, err error) {
	reqMap, reqMapErr := convertor.StructToMap(req)
	if reqMapErr != nil {
		err = reqMapErr
		return
	}
	var result []string
	for key, value := range reqMap {
		result = append(result, fmt.Sprintf("%s=%s", key, url.QueryEscape(convertor.ToString(value))))
	}
	resp = strings.Join(result, "&")
	return
}
