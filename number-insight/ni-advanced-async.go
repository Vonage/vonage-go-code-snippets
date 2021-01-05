package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/vonage/vonage-go-sdk"
)

func main() {
	godotenv.Load("../.env")

	auth := vonage.CreateAuthFromKeySecret(os.Getenv("VONAGE_API_KEY"), os.Getenv("VONAGE_API_SECRET"))
	niClient := vonage.NewNumberInsightClient(auth)

	niClient.AdvancedAsync(os.Getenv("INSIGHT_NUMBER"), os.Getenv("SERVER_BASE_URL")+"/webhooks/insight", vonage.NiOpts{})
}
