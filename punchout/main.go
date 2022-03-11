package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/vincent87720/AutoPunch/notify"
	"github.com/vincent87720/AutoPunch/punch"
)

type Account struct {
	ID   string
	PW   string
	LINE string
}

func main() {
	info, err := ioutil.ReadFile("./info.json")
	if err != nil {
		log.Fatal(err)
	}
	var accounts []Account
	err = json.Unmarshal(info, &accounts)
	if err != nil {
		log.Fatal(err)
	}

	for _, val := range accounts {
		fmt.Println("Punch out: " + val.ID)
		statusMsg := punch.Punch(val.ID, val.PW, "B")
		if val.LINE != "" {
			notify.NotifyPunchStatus(val.ID, val.LINE, statusMsg)
		}
	}
}
