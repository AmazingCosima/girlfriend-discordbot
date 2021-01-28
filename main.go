package main

import (
	"github.com/amazingcosima/girlfriend-discordbot/girlfriend"
	"io/ioutil"
)

func main() {
	token, err := ioutil.ReadFile(".token")
	if err != nil {
	}
	girlfriend.Login(string(token))
}