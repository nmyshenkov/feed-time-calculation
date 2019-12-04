package api

import (
	"encoding/json"
	"feed-time-calculation/cars"
	"feed-time-calculation/predict"
	"io/ioutil"
	"log"
	"net/http"
)

// GetFeedTime - метод получения списка
func (t TypeApi) GetFeedTime(w http.ResponseWriter, r *http.Request) {

	var (
		in Request
	)

	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	// проверям, что пришел POST
	if checkHttpPost(w, r) != nil {
		return
	}
	// читаем Входящий запрос
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		setError(w, BadRequest)
		return
	}
	// попытка распарсить входящий запрос
	if err := json.Unmarshal(reqBody, &in); err != nil {
		log.Println(err)
		setError(w, BadRequest)
		return
	}

	// валидируем входные парамеры
	if in.Lat == 0 {
		setError(w, "Не передан обязательный параметр lat")
		return
	}

	if in.Lng == 0 {
		setError(w, "Не передан обязательный параметр lng")
		return
	}

	if in.Limit > 100 || in.Limit < 0 {
		setError(w, "Параметр limit имеет не валидный формат")
		return
	}

	// получаем список машин поблизости
	nearbyCars, err := cars.GetCars(t.httpCarCoordinatePool, cars.Coordinate{
		Lat:   in.Lat,
		Lng:   in.Lng,
		Limit: in.Limit,
	})
	if err != nil {
		setError(w, err.Error())
		return
	}

	// если нет машин поблизости - отдаем пустой результат
	if len(nearbyCars) == 0 {
		setResult(w, []CarWithTime{})
		return
	}

	// инициализируем данные
	var carsCooord []predict.Coordinate
	var result []CarWithTime

	// подготавливаем данные
	for _, car := range nearbyCars {
		// проверям, что координаты корректные
		if car.Lng == 0 || car.Lat == 0 {
			continue
		}
		carsCooord = append(carsCooord, predict.Coordinate{
			Lat: car.Lat,
			Lng: car.Lng,
		})
		result = append(result, CarWithTime{ID: car.ID})
	}

	// если не получилось собрать массив с координатами - отдаем пустой результат
	if len(carsCooord) == 0 {
		setResult(w, []CarWithTime{})
		return
	}

	// получаем список пргнозов
	predicts, err := predict.GetPredict(t.httpRoutePool, predict.Coordinate{
		Lat: in.Lat,
		Lng: in.Lng,
	}, carsCooord)
	if err != nil {
		setError(w, err.Error())
		return
	}

	// если данных не равно резултатом - делам проверку, чтобы не было ошибки индексирования
	if len(predicts) > len(result) {
		setError(w, "Ошибка при расчете")
		return
	}

	// сохраняем координаты в той же последовательности
	for i, pred := range predicts {
		result[i].Time = pred
	}

	setResult(w, result)
	return
}
