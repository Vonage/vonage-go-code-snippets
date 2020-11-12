package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vonage/vonage-go-sdk/ncco"
)

func answer(w http.ResponseWriter, req *http.Request) {

	MyNcco := ncco.Ncco{}

	talk := ncco.TalkAction{Text: "Thank you for calling."}
	MyNcco.AddAction(talk)

	data, _ := json.Marshal(MyNcco)
	fmt.Fprintf(w, "%s", data)
}

func main() {

	http.HandleFunc("/answer", answer)

	http.ListenAndServe(":8081", nil)
}
