package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/vonage/vonage-go-sdk"
	"github.com/vonage/vonage-go-sdk/ncco"
)

func answer(w http.ResponseWriter, req *http.Request) {

	MyNcco := ncco.Ncco{}

	talk := ncco.TalkAction{Text: "Thank you for calling."}
	MyNcco.AddAction(talk)

	data, _ := json.Marshal(MyNcco)
	fmt.Fprintf(w, "%s", data)
}

func playAudio(w http.ResponseWriter, req *http.Request) {
	privateKey, _ := ioutil.ReadFile(PATH_TO_PRIVATE_KEY_FILE)
	auth, _ := vonage.CreateAuthFromAppPrivateKey("00001111-aaaa-bbbb-cccc-0123456789abcd", privateKey)
	client := vonage.NewVoiceClient(auth)

	result, _, _ := client.PlayAudioStream("aaaabbbb-0000-1111-2222-abcdef01234567",
		"https://raw.githubusercontent.com/nexmo-community/ncco-examples/gh-pages/assets/welcome_to_nexmo.mp3",
		vonage.PlayAudioOpts{},
	)
	// or to stop the audio
	// result, _, _:= client.StopAudioStream("aaaabbbb-0000-1111-2222-abcdef01234567")
	fmt.Println("Update message: " + result.Message)
}

func main() {

	http.HandleFunc("/webhooks/answer", answer)
	http.HandleFunc("/play-audio", playAudio)

	http.ListenAndServe(":8081", nil)
}
