package common

import (
	"net/http"
	"strconv"
	"sync"
	"time"
)

func FloatToString(in float64) string {
	return strconv.FormatFloat(in, 'f', 7, 64)
}

func IntToString(in int64) string {
	return strconv.FormatInt(in, 10)
}

// InitHttpPool - получение пул http соединений
func InitHttpPool(timeout int64) *sync.Pool {
	transport := &http.Transport{}

	return &sync.Pool{
		New: func() interface{} {
			return &http.Client{
				Transport: transport,
				Timeout:   time.Duration(timeout) * time.Second,
			}
		},
	}
}
