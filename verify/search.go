package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/kr/pretty"
	"github.com/vonage/vonage-go-sdk"
)

func main() {
	godotenv.Load("../.env")

	auth := vonage.CreateAuthFromKeySecret(os.Getenv("VONAGE_API_KEY"), os.Getenv("VONAGE_API_SECRET"))
	verifyClient := vonage.NewVerifyClient(auth)

	response, errResp, err := verifyClient.Search(os.Getenv("REQUEST_ID"))

	if err != nil {
		fmt.Printf("%#v\n", err)
	} else if response.Status == "" {
		fmt.Println("Error status " + errResp.Status + ": " + errResp.ErrorText)
	} else {
		fmt.Printf("%# v", pretty.Formatter(response))
	}
}
