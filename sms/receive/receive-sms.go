package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// Sms defines parameters returned in Webhook request, consisting of SMS content.
type Sms struct {
	Msisdn           string
	To               string
	MessageID        string
	Text             string
	Type             string
	Keyword          string
	MessageTimestamp string
	APIKey           string
}

func main() {
	// Loaded .env file, to compare API keys
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	http.HandleFunc("/webhook/receive-sms", receiveSms)

	log.Print("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

func receiveSms(w http.ResponseWriter, r *http.Request) {
	receivedSmsDetails := Sms{
		Msisdn:           r.FormValue("msisdn"),
		To:               r.FormValue("to"),
		MessageID:        r.FormValue("messageId"),
		Text:             r.FormValue("text"),
		Type:             r.FormValue("type"),
		Keyword:          r.FormValue("keyword"),
		MessageTimestamp: r.FormValue("message-timestamp"),
		APIKey:           r.FormValue("api-key"),
	}

	// Checking to ensure "api-key" returned in webhook matches your API key defined in your .env file.
	if receivedSmsDetails.APIKey != os.Getenv("VONAGE_API_KEY") {
		fmt.Println("APIKey doesn't match your VONAGE_API_KEY environment variable.")

		return
	}

	log.Printf("%+v\n", receivedSmsDetails)

	return
}
