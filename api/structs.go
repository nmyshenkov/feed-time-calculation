package api

const (
	BadRequest = "Bad request"
)

type Request struct {
	Lat   float64 `json:"lat"`
	Lng   float64 `json:"lng"`
	Limit int64   `json:"limit"`
}

type Response struct {
	Error  string        `json:"error"`
	Result []CarWithTime `json:"result"`
}

type CarWithTime struct {
	ID   int `json:"id"`
	Time int `json:"time"`
}
