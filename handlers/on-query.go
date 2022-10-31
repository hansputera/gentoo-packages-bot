package handlers

import (
	"fmt"

	"gopkg.in/telebot.v3"
)

func OnQueryFunc(ctx telebot.Context) error {
	fmt.Println(ctx.Args())

	return nil
}
