package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/vonage/vonage-go-sdk"
)

func main() {
	godotenv.Load("../.env")

	auth := vonage.CreateAuthFromKeySecret(os.Getenv("VONAGE_API_KEY"), os.Getenv("VONAGE_API_SECRET"))
	niClient := vonage.NewNumberInsightClient(auth)

	result, _, _ := niClient.Basic(os.Getenv("INSIGHT_NUMBER"), vonage.NiOpts{})

	result_json, _ := json.MarshalIndent(result, "", "")
	fmt.Println(string(result_json))

}
