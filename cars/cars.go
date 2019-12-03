package cars

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"whelly/common"
)

func GetCars(clientPool *sync.Pool, coord Coordinate) ([]Cars, error) {

	var (
		httpResponse *http.Response
		u            *url.URL
		httpRequest  *http.Request
		err          error
		respBody     []byte
		response     []Cars
	)

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

	if httpResponse.StatusCode != http.StatusOK {
		return nil, errors.New("ошибка получения списка машин")
	}

	if respBody, err = ioutil.ReadAll(httpResponse.Body); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(respBody, response); err != nil {
		return nil, err
	}

	return response, nil
}
