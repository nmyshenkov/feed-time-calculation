package api

const (
	BadRequest = "Bad request"
)

// Request - входящий запрос
type Request struct {
	Lat   float64 `json:"lat"`
	Lng   float64 `json:"lng"`
	Limit int64   `json:"limit"`
}

// Response - ответ сервиса
type Response struct {
	Error  string        `json:"error"`
	Result []CarWithTime `json:"result"`
}

// CarWithTime - стрктура с данными id машины и время подачи
type CarWithTime struct {
	ID   int `json:"car_id"`
	Time int `json:"time"`
}

type ByTime []CarWithTime

// функции для сортировки
func (s ByTime) Len() int           { return len(s) }
func (s ByTime) Less(i, j int) bool { return s[i].Time < s[j].Time }
func (s ByTime) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
