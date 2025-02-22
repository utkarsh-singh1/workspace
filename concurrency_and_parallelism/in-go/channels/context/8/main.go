package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type eventData struct {

	Name string
	Description string
	Location string
	Time time.Time
}

func main(){

	e1 := eventData{

		"Testname1", "A normal 1 day event","New york",time.Now(),

	}


	e2 := eventData{

		"Testname2", "A normal 2 day event","Delhi",time.Now(),
	}

	Event := []eventData{e1,e2}

	inJson, err := json.Marshal(Event)
	if err != nil {
		log.Fatalln("There is an Error you wanna fix it ->",err)
		
	}

	fmt.Println(string(inJson))
}
