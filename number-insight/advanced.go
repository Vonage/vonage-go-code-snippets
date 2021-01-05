package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/vonage/vonage-go-sdk"
)

func main() {
	godotenv.Load("../.env")

	auth := vonage.CreateAuthFromKeySecret(os.Getenv("VONAGE_API_KEY"), os.Getenv("VONAGE_API_SECRET"))
	niClient := vonage.NewNumberInsightClient(auth)

	result, _, _ := niClient.AdvancedAsync(os.Getenv("INSIGHT_NUMBER"), "https://example.com/number-insight-data", vonage.NiOpts{})
	if result.Status == 0 {
		http.HandleFunc("/number-insight-data", func(w http.ResponseWriter, r *http.Request) {
			data, _ := ioutil.ReadAll(r.Body)
			fmt.Println(string(data))
		})

		http.ListenAndServe(":8080", nil)
	} else {
		fmt.Println("Request status " + string(result.Status) + ": " + result.StatusMessage)
	}
}
