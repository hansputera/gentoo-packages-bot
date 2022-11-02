package handlers

import (
	"gentoo-packages-bot/handlers/controllers"
	"strings"

	"gopkg.in/telebot.v3"
)

func send_usage(ctx *telebot.Context) {
	go (*ctx).Answer(&telebot.QueryResponse{
		Results: []telebot.Result{
			&telebot.ArticleResult{
				Title:       "Invalid",
				Text:        "Correct format: @" + (*ctx).Bot().Me.Username + " <useflag | package> <package/useflag name>\n\nExample: @" + (*ctx).Bot().Me.Username + " package firefox",
				Description: "Please input data type, and the query to search!",
			},
		},
		CacheTime: 5,
	})
}

func OnQueryFunc(ctx telebot.Context) error {
	args := ctx.Args()

	if len(args) < 2 {
		go send_usage(&ctx)
	}

	switch strings.ToLower(args[0]) {
	case "useflag":
		go controllers.UseflagController(args[1:], ctx)
	default:
		go send_usage(&ctx)
	}
	return nil
}
