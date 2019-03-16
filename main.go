package main

import (
	"encoding/json"
	"net/http"
)

type weatherData struct {
	Cod  string `json:"cod"`
	City struct {
		Name string `json:"name"`
	} `json:"city"`
}

func main() {
	http.HandleFunc("/", hello)

	http.HandleFunc("/weather/", func(w http.ResponseWriter, r *http.Request) {
		data, err := query()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(data)
	})

	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello!"))
}

func query() (weatherData, error) {
	resp, err := http.Get("https://samples.openweathermap.org/data/2.5/forecast?id=524901&appid=b1b15e88fa797225412429c1c50c122a1")
	if err != nil {
		return weatherData{}, err
	}

	defer resp.Body.Close()

	var d weatherData

	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return weatherData{}, err
	}

	return d, nil
}
