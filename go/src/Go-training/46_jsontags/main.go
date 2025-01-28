package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type User struct {
	Name string `json:"name"`

	// The swd is sensitive if it leaks, say goodbye to your daata, in this case we can have special struct tag so that we can have swd field omitted from the json output, the tag is `json:" - "`. Keep note of the space.
	Swd string `json:"swd"`

	// Somehow i don't want to have this field any kind of value currently, in future it may change but in present na-na, don't want any value here when creating its instance, so i will suffix ",omitempty" to the json struct tag like this `json:"age,omitempty"` to make this field omitted from json output after marshal or if not used`json:"age",omitempty`, json output will have age field with null value init.
	Age int `json:"age,omitempty"`

	CreatedAt time.Time `json:"createdAt"`

	// There are more special json struct tags to deal with the json output but that is limited to encoding/json package.
}

type op struct {
	name string
	swd  string
	age  int
}

func main() {
	u := &User{
		Name:      "baburao",
		Swd:       "champaklal",
		CreatedAt: time.Now(),
		//Age:       32,
	}

	out, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(out))
}
