package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Schema struct {
	Id   string
	Name string
}

func main() {
	fmt.Println("Starting GO Server!")

	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("=> ", r.URL, r.Host)
		dbSchema := Schema{
			Id:   "0",
			Name: "schema0",
		}
		data, err := json.Marshal(dbSchema)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("response: ", data)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(data)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Err: ", err)
	}
}
