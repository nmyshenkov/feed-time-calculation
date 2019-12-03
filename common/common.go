package common

import "strconv"

func FloatToString(in float64) string {
	return strconv.FormatFloat(in, 'f', 7, 64)
}

func IntToString(in int64) string {
	return strconv.FormatInt(in, 10)
}
