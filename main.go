package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/raaaaaaaay86/time-now-go-unit-test/service"
)

func main() {
	svc := service.BetterMessageService{
		Now: func() time.Time {
			return time.Now()
		},
	}

	title := "Help!!"
	content := "I forgot to take my wallet. Would you help me to bring it to the restaurant?"
	msg, err := svc.Send(title, content)
	if err != nil {
		panic(err)
	}

	prettyJSON, err := toPrettyJSON(msg)
	if err != nil {
		panic(err)
	}

	fmt.Println(prettyJSON)
}

func toPrettyJSON(o any) (string, error) {
	slice, err := json.MarshalIndent(o, "", "\t")
	if err != nil {
		return "", nil
	}

	return string(slice), nil
}
