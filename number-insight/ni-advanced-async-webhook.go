package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/webhooks/insight", func(w http.ResponseWriter, r *http.Request) {
		data, _ := ioutil.ReadAll(r.Body)
		fmt.Println(string(data))
	})

	http.ListenAndServe(":3000", nil)
}
