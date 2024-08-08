package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type tableSchema struct {
	Key   string
	Value string
}

var list = []tableSchema{}

func main() {
	fmt.Println("Starting GO Server!")

	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("=> ", r.URL, r.Host)

		newEntry := tableSchema{
			Key:   "0",
			Value: "Hello",
		}
		data, err := json.Marshal(newEntry)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("response: ", data)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(200)
			w.Write(data)
		}
	})

	http.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("=> ", r.URL, r.Body)
		var entry tableSchema
		err := json.NewDecoder(r.Body).Decode(&entry)

		if err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
			w.Write([]byte(`{"key" : -1, "value": "failed!"}`))
		} else {
			list = append(list, entry)
			fmt.Println("received: ", list)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(201)
			w.Write([]byte(`{"key" : 0, "value": "received!"}`))
		}

	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Err: ", err)
	}
}
