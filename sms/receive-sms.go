package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/webhooks/inbound-sms", func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		fmt.Println("From: " + params["msisdn"][0] + ", message: " + string(params["text"][0]))
	})

	http.ListenAndServe(":8080", nil)
}
