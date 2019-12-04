package predict

const URL = "https://dev-api.wheely.com/fake-eta/predict"

type Request struct {
	Target Coordinate   `json:"target"`
	Source []Coordinate `json:"source"`
}

type Coordinate struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Response []int
