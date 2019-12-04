package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func setResult(w http.ResponseWriter, result []CarWithTime) {
	var resp Response
	resp.Result = result
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Println(err)
	}
	return
}

func setError(w http.ResponseWriter, str string) {
	var resp Response
	resp.Error = str
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Println(err)
	}
	return
}

func checkHttpPost(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "POST" {
		var resp Response
		resp.Error = "Method not allow"
		w.WriteHeader(http.StatusMethodNotAllowed)
		if err := json.NewEncoder(w).Encode(r); err != nil {
			log.Println(err)
		}
		log.Println(resp.Error)
		return errors.New(resp.Error)
	}
	return nil
}
