package cars

const URL = "https://localhost:8081/cars"

type Coordinate struct {
	Lat   float64 `json:"lat"`
	Lng   float64 `json:"lng"`
	Limit int64   `json:"limit"`
}

// Cars - структура с мащинами
type Cars struct {
	ID  int     `json:"id"`
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
