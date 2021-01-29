package main

import (
	"bufio"
	"github.com/amazingcosima/girlfriend-discordbot/girlfriend"
	"io/ioutil"
	"os"
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
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input = scanner.Text()
			split := strings.Split(input, " ")
			app.SendMessage(split[0], split[1])
		}
	}
}