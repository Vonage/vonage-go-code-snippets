package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/vonage/vonage-go-sdk"
)

func main() {
	godotenv.Load("../.env")

	var from string = os.Getenv("VONAGE_BRAND_NAME")
	var to string = os.Getenv("TO_NUMBER")
	var text string = "A text message sent using the Vonage SMS API"

	auth := vonage.CreateAuthFromKeySecret(os.Getenv("VONAGE_API_KEY"), os.Getenv("VONAGE_API_SECRET"))
	smsClient := vonage.NewSMSClient(auth)
	response, errResp, err := smsClient.Send(from, to, text, vonage.SMSOpts{})

	if err != nil {
		panic(err)
	}

	if response.Messages[0].Status == "0" {
		fmt.Println("Message sent successfully.")
	} else {
		fmt.Println("Error code " + errResp.Messages[0].Status + ": " + errResp.Messages[0].ErrorText)
	}
}
