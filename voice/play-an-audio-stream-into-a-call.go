package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/joho/godotenv"
	"github.com/vonage/vonage-go-sdk"
)

func main() {
	godotenv.Load("../.env")

	privateKey, _ := ioutil.ReadFile(os.Getenv("VONAGE_APPLICATION_PRIVATE_KEY_PATH"))
	auth, _ := vonage.CreateAuthFromAppPrivateKey(os.Getenv("VONAGE_APPLICATION_ID"), privateKey)
	client := vonage.NewVoiceClient(auth)

	result, _, _ := client.PlayAudioStream(os.Getenv("UUID"),
		"https://nexmo-community.github.io/ncco-examples/assets/voice_api_audio_streaming.mp3",
		vonage.PlayAudioOpts{},
	)

	// or to stop the audio
	// result, _, _:= client.StopAudioStream(os.Getenv("UUID"))
	fmt.Println("Update message: " + result.Message)
}
