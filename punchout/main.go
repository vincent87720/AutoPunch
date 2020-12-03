package main

import (
	punch "AutoPunch/punch"
	"io/ioutil"
	"log"
	"encoding/json"
	"fmt"
)

type Account struct{
	ID string
	PW string
}

func main() {
	info,err := ioutil.ReadFile("./info.json")
	if err != nil {
		log.Fatal(err)
	}
	var accounts []Account
	err = json.Unmarshal(info,&accounts)
	if err != nil {
		log.Fatal(err)
	}

	for _,val := range accounts{
		fmt.Println("Punch out: "+val.ID)
		punch.Punch(val.ID, val.PW, "B")
	}
}
