package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type WebResponse struct {
	Code        int    `json:"Code"`
	Description string `json:"Description"`
}

//var WRs = []WebResponse{}

func main() {
	var Data []WebResponse

	rcvd := `[{"Code":200,"Description":"StatusOK"},{"Code":301,"Description":"StatusMovedPermanently"},{"Code":302,"Description":"StatusFound"},{"Code":303,"Description":"StatusSeeOther"},{"Code":307,"Description":"StatusTemporaryRedirect"},{"Code":400,"Description":"StatusBadRequest"},{"Code":401,"Description":"StatusUnauthorized"},{"Code":402,"Description":"StatusPaymentRequired"},{"Code":403,"Description":"StatusForbidden"},{"Code":404,"Description":"StatusNotFound"},{"Code":405,"Description":"StatusMethodNotAllowed"},{"Code":418,"Description":"StatusTeapot"},{"Code":500,"Description":"StatusInternalServerError"}]`

	err := json.Unmarshal([]byte(rcvd), &Data)
	if err != nil {
		fmt.Printf("%v\n", err)
		log.Fatalln(err)
	}

	for _, t := range Data {
		fmt.Println(t.Code, " - ", t.Description)
	}
}
