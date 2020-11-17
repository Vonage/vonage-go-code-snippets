package main

import (
	"fmt"

	"github.com/vonage/vonage-go-sdk"
)

func main() {
	var from string = VONAGE_BRAND_NAME
	var to string = TO_NUMBER
	var text string = "こんにちは世界"

	auth := vonage.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	smsClient := vonage.NewSMSClient(auth)
	response, err := smsClient.Send(from, to, text, vonage.SMSOpts{Type: "unicode"})

	if err != nil {
		panic(err)
	}

	if response.Messages[0].Status == "0" {
		fmt.Println("Message sent successfully.")
	}
}
