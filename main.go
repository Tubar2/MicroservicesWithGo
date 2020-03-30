package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(_ http.ResponseWriter, r *http.Request) {
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal("Error reading body")
		}
		log.Println(string(d))
	})

	http.ListenAndServe("localhost:9090", nil)
}
