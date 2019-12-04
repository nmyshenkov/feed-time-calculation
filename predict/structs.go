package predict

const URL = "https://localhost:8081/predict"

type Request struct {
	Target Coordinate   `json:"target"`
	Source []Coordinate `json:"source"`
}

type Coordinate struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Response []int
