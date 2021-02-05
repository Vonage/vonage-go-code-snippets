package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/vonage/vonage-go-sdk/ncco"
)

type Dtmf struct {
	Digits    string
	Timed_out bool
}

type Response struct {
	Speech            []string
	Dtmf              Dtmf
	From              string
	To                string
	Uuid              string
	Conversation_uuid string
	Timestamp         string
}

func answer(w http.ResponseWriter, req *http.Request) {
	MyNcco := ncco.Ncco{}

	talk := ncco.TalkAction{Text: "Hello please press any key to continue."}
	MyNcco.AddAction(talk)

	log.Println()
	log.Println(req.Host)
	inputAction := ncco.InputAction{EventUrl: []string{"https://demo.ngrok.io/webhooks/dtmf"}, Dtmf: &ncco.DtmfInput{MaxDigits: 1}}
	MyNcco.AddAction(inputAction)

	data, _ := json.Marshal(MyNcco)

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func dtmf(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)
	var t Response
	json.Unmarshal(data, &t)

	MyNcco := ncco.Ncco{}
	talk := ncco.TalkAction{Text: "You pressed " + t.Dtmf.Digits + ", Goodbye"}
	MyNcco.AddAction(talk)

	responseData, _ := json.Marshal(MyNcco)
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseData)
}

func main() {

	http.HandleFunc("/webhooks/answer", answer)
	http.HandleFunc("/webhooks/dtmf", dtmf)

	http.ListenAndServe(":3000", nil)
}
