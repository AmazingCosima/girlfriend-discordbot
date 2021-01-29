package main

import (
	"fmt"
	"github.com/amazingcosima/girlfriend-discordbot/girlfriend"
	"io/ioutil"
	"strings"
)

func main() {
	token, err := ioutil.ReadFile(".token")
	if err != nil {
	}
	var app = new(girlfriend.App)
	err = app.CreateSession(string(token))
	var input string
	for {
		_, err = fmt.Scan(&input)
		split := strings.Split(input, ",")
		app.SendMessage(split[0], split[1])
	}
}