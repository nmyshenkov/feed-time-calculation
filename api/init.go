package api

import (
	"feed-time-calculation/common"
	"sync"
)

// ApiType - главная структура
type TypeApi struct {
	httpCarCoordinatePool,
	httpRoutePool *sync.Pool
}

// InitApi - инициализация главной структуры
func InitApi() *TypeApi {
	var (
		httpCarCoordinatePool, httpRoutePool *sync.Pool
	)

	// Инициализация пула соединений для сервиса координат машин
	httpCarCoordinatePool = common.InitHttpPool(5)
	// Инициализация пула соединений для сервиса времени маршрута
	httpRoutePool = common.InitHttpPool(5)

	return &TypeApi{
		httpCarCoordinatePool: httpCarCoordinatePool,
		httpRoutePool:         httpRoutePool,
	}
}
