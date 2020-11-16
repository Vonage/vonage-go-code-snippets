package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vonage/vonage-go-sdk/ncco"
)

func answer(w http.ResponseWriter, req *http.Request) {

	paramKeys, _ := req.URL.Query()["from"]

	MyNcco := ncco.Ncco{}

	talk := ncco.TalkAction{Text: "Thank you for calling." + string(paramKeys[0])}
	MyNcco.AddAction(talk)

	data, _ := json.Marshal(MyNcco)
	fmt.Fprintf(w, "%s", data)
}

func main() {

	http.HandleFunc("/webhooks/answer", answer)

	http.ListenAndServe(":3000", nil)
}
