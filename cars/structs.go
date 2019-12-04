package cars

const URL = "https://dev-api.wheely.com/fake-eta/cars"

type Coordinate struct {
	Lat   float64 `json:"lat"`
	Lng   float64 `json:"lng"`
	Limit int64   `json:"limit"`
}

// Cars
type Cars struct {
	ID  int     `json:"id"`
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
