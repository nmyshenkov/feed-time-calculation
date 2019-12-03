package api

import (
	"net/http"
	"sync"
	"time"
)

// ApiType - главная структура
type TypeApi struct {
	DoneChan <-chan struct{}
	httpCarCoordinatePool,
	httpRoutePool *sync.Pool
}

func InitApi() *TypeApi {
	var (
		httpCarCoordinatePool, httpRoutePool *sync.Pool
	)

	// Инициализация пула соединений для ЕФСП
	httpCarCoordinatePool = initHttpPool(5)
	// Инициализация пула соединений для Мосэнергосбыта
	httpRoutePool = initHttpPool(5)

	return &TypeApi{
		// DoneChan:              doneChan,
		httpCarCoordinatePool: httpCarCoordinatePool,
		httpRoutePool:         httpRoutePool,
	}
}

// initHttpPool - получение пул http соединений
func initHttpPool(timeout int64) *sync.Pool {
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
