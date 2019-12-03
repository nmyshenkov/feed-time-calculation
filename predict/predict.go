package predict

import (
	"bytes"
	"encoding/json"
	"net/http"
	"sync"
)

func GetPredict(clientPool *sync.Pool, target Coordinate, source []Coordinate) (Response, error) {

	var (
		jsonPostData []byte
		httpResponse *http.Response
		httpRequest  *http.Request
		err          error
	)

	if jsonPostData, err = json.Marshal(Request{Target: target, Source: source}); err != nil {
		return nil, err
	}

	if httpRequest, err = http.NewRequest("POST", URL, bytes.NewBuffer(jsonPostData)); err != nil {
		return nil, err
	}

	client := clientPool.Get().(*http.Client)
	defer clientPool.Put(client)

	if httpResponse, err = client.Do(httpRequest); err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()

	return nil, nil
}
