package predict

const URL = "https://localhost:7000/predict"

type Request struct {
	Target Coordinate   `json:"target"`
	Source []Coordinate `json:"source"`
}

type Response []int

type Coordinate struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
