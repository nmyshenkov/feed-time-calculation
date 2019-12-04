package cars

import (
	"encoding/json"
	"errors"
	"feed-time-calculation/common"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sync"
)

// Функция получения координат машин
func GetCars(clientPool *sync.Pool, coord Coordinate) ([]Cars, error) {

	var (
		httpResponse *http.Response
		u            *url.URL
		httpRequest  *http.Request
		err          error
		respBody     []byte
		response     []Cars
	)

	// если не передан limit - задаем как 5 машин
	if coord.Limit == 0 {
		coord.Limit = 5
	}

	u, err = url.Parse(URL)
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Add("lat", common.FloatToString(coord.Lat))
	q.Add("lng", common.FloatToString(coord.Lat))
	q.Add("limit", common.IntToString(coord.Limit))
	u.RawQuery = q.Encode()

	if httpRequest, err = http.NewRequest("GET", u.String(), nil); err != nil {
		return nil, err
	}

	client := clientPool.Get().(*http.Client)
	defer clientPool.Put(client)

	if httpResponse, err = client.Do(httpRequest); err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()

	if respBody, err = ioutil.ReadAll(httpResponse.Body); err != nil {
		return nil, err
	}

	if httpResponse.StatusCode != http.StatusOK {
		log.Println(string(respBody))
		return nil, errors.New("ошибка получения списка машин")
	}

	if err = json.Unmarshal(respBody, &response); err != nil {
		return nil, err
	}

	return response, nil
}
