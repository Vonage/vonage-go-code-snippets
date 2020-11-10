package main

import (
	"fmt"
	"io/ioutil"

	"github.com/vonage/vonage-go-sdk"
	"github.com/vonage/vonage-go-sdk/ncco"
)

func main() {
	privateKey, _ := ioutil.ReadFile(PATH_TO_PRIVATE_KEY_FILE)
	auth, _ := vonage.CreateAuthFromAppPrivateKey(APPLICATION_ID, privateKey)
	client := vonage.NewVoiceClient(auth)

	from := vonage.CallFrom{Type: "phone", Number: FROM_NUMBER}
	to := vonage.CallTo{Type: "phone", Number: TO_NUMBER}

	MyNcco := ncco.Ncco{}
	talk := ncco.TalkAction{Text: "This is the Go code snippets calling to say hello."}
	MyNcco.AddAction(talk)

	// NCCO example
	result, _, _ := client.CreateCall(vonage.CreateCallOpts{From: from, To: to, Ncco: MyNcco})
	// alternate version with answer URL
	//result, _, _ := client.CreateCall(CreateCallOpts{From: from, To: to, AnswerUrl: []string{"https://example.com/answer"}})

	fmt.Println(result.Uuid + " call ID started")
}
