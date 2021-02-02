package main

import (
	"bufio"
	"github.com/amazingcosima/girlfriend-discordbot/girlfriend"
	"io/ioutil"
	"os"
)

func main() {
	token, err := ioutil.ReadFile(".token")
	if err != nil {
	}
	var app = new(girlfriend.App)
	err = app.CreateSession(string(token))
	for {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
		}
	}
}
