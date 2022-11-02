package controllers

import (
	"fmt"
	"gentoo-packages-bot/scraping"
	"strings"

	"gopkg.in/telebot.v3"
)

func UseflagController(args []string, ctx telebot.Context) {
	result := scraping.GetUseFlag(strings.Join(args, " "))
	if result.Name == "" {
		go ctx.Answer(&telebot.QueryResponse{
			CacheTime: (24 * 60) * 60, // 1 day
			Results: []telebot.Result{
				&telebot.ArticleResult{
					Title:       "404 Not Found",
					Description: "I couldn't find package " + strings.Join(args, " "),
					Text:        "I couldn't find " + strings.Join(args, "") + ", try it again in a day :)",
					ThumbURL:    "https://www.gentoo.org/assets/img/logo/gentoo-logo.png",
				},
			},
		})
	} else {
		len_slice := 50
		if len(result.Tooltip) < len_slice {
			len_slice = len(result.Tooltip)
		}

		teleResult := &telebot.ArticleResult{
			Title:       result.Name,
			Description: result.Tooltip[:len_slice] + "...",
			URL:         result.Url,
			ThumbURL:    "https://www.gentoo.org/assets/img/logo/gentoo-logo.png",
		}

		pkgs := []string{}

		for _, pkg := range result.Packages {
			pkgs = append(pkgs, fmt.Sprintf(`<a href="%s">%s/%s</a>`, pkg.Url, pkg.Group, pkg.Package))
		}

		teleResult.SetContent(&telebot.InputTextMessageContent{
			Text:      fmt.Sprintf("Useflag <strong>%s</strong> - <i>%s</i>\n\nPackages: %s", result.Name, result.Tooltip, strings.Join(pkgs, ", ")),
			ParseMode: "HTML",
		})

		go ctx.Answer(&telebot.QueryResponse{
			CacheTime: 2 * 60,
			Results: []telebot.Result{
				teleResult,
			},
		})
	}
}
