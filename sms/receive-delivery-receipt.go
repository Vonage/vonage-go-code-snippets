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

		fmt.Println("Delivery receipt status: " + r.FormValue("status"))
	})

	http.ListenAndServe(":8080", nil)
}
