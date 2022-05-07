package main

import (
	"fmt"
	"log"
	"time"

	"github.com/abdfnx/botwaygo"
	tele "gopkg.in/telebot.v3"
)

func main() {
	pref := tele.Settings{
		Token:  botwaygo.GetToken(),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello!")
	})

	fmt.Println("Bot started")
	b.Start()
}
