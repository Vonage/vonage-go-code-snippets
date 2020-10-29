package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/vonage/vonage-go-sdk"
)

func main() {
	godotenv.Load("../.env")

	auth := vonage.CreateAuthFromKeySecret(os.Getenv("VONAGE_API_KEY"), os.Getenv("VONAGE_API_SECRET"))
	verifyClient := vonage.NewVerifyClient(auth)

	response, errResp, err := verifyClient.Request(os.Getenv("RECIPIENT_NUMBER"), os.Getenv("BRAND_NAME"), vonage.VerifyOpts{})

	if err != nil {
		fmt.Printf("%#v\n", err)
	} else if response.Status != "0" {
		fmt.Println("Error status " + errResp.Status + ": " + errResp.ErrorText)
	} else {
		fmt.Println(response.RequestId)
	}
}
