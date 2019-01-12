package main

import (
	"log"

	"github.com/hugoamvieira/intercom-users-inviter/config"
)

func main() {
	c, err := config.NewJSON("config/config.json")
	if err != nil {
		log.Fatalln("Couldn't obtain configuration for this program, cannot continue | Error:", err)
	}

	i := NewInviter(c)
	err = i.Start()
	if err != nil {
		log.Fatalln("Issue with running inviter | Error:", err)
	}
}
