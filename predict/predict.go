package predict

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func GetPredict(clientPool *sync.Pool, target Coordinate, source []Coordinate) (Response, error) {

	var (
		jsonPostData []byte
		httpResponse *http.Response
		httpRequest  *http.Request
		err          error
		respBody     []byte
		response     Response
	)

	if jsonPostData, err = json.Marshal(Request{Target: target, Source: source}); err != nil {
		return nil, err
	}

	log.Println(string(jsonPostData))

	if httpRequest, err = http.NewRequest("POST", URL, bytes.NewBuffer(jsonPostData)); err != nil {
		return nil, err
	}

	httpRequest.Header.Set("content-type", "application/json")

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
		return nil, errors.New("ошибка получения времени маршрута")
	}

	if err = json.Unmarshal(respBody, &response); err != nil {
		return nil, err
	}

	return response, nil
}
