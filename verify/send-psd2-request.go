package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/vonage/vonage-go-sdk"
)

func main() {
	godotenv.Load("../.env")

	auth := vonage.CreateAuthFromKeySecret(os.Getenv("VONAGE_API_KEY"), os.Getenv("VONAGE_API_SECRET"))
	verifyClient := vonage.NewVerifyClient(auth)

	amount, _ := strconv.ParseFloat(os.Getenv("AMOUNT"), 32)
	response, errResp, _ := verifyClient.Psd2(
		os.Getenv("RECIPIENT_NUMBER"),
		os.Getenv("PAYEE_NAME"),
		amount,
		vonage.VerifyPsd2Opts{},
	)

	if response.Status != "0" {
		fmt.Println("Error status " + errResp.Status + ": " + errResp.ErrorText)
	} else {
		fmt.Println(response.RequestId)
	}
}
