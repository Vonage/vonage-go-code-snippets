package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/webhooks/delivery-receipt", func(w http.ResponseWriter, r *http.Request) {

		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		if r.FormValue("status") == "delivered" {
			fmt.Println("Your message to " + r.FormValue("msisdn") + " message id " + r.FormValue("messageId") + ") was delivered.")
		} else if r.FormValue("status") == "accepted" {
			fmt.Println("Your message to " + r.FormValue("msisdn") + " message id " + r.FormValue("messageId") + ") was accepted by the carrier.")
		} else {
			fmt.Println("Your message to " + r.FormValue("msisdn") + " has a status of: " + r.FormValue("status") + ".")
			fmt.Println("Check err-code " + r.FormValue("err-code") + " against the documentation.")
		}
	})

	http.ListenAndServe(":8080", nil)
}
