package main

import (
	"encoding/json"
	"net/http"

	"github.com/vonage/vonage-go-sdk/ncco"
)

func answer(w http.ResponseWriter, req *http.Request) {

	paramKeys, _ := req.URL.Query()["from"]

	MyNcco := ncco.Ncco{}

	talk := ncco.TalkAction{Text: "Thank you for calling " + string(paramKeys[0])}
	MyNcco.AddAction(talk)

	data, _ := json.Marshal(MyNcco)

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func main() {

	http.HandleFunc("/webhooks/answer", answer)

	http.ListenAndServe(":8080", nil)
}
