package scraping

import (
	"gentoo-packages-bot/structs"
	"gentoo-packages-bot/utils"
	"strings"

	"github.com/gocolly/colly/v2"
)

func GetUseFlag(flag string) *structs.UseFlagDetail {
	useflag := &structs.UseFlagDetail{}
	collector := colly.NewCollector()

	collector.OnHTML(".kk-header-container", func(h *colly.HTMLElement) {
		useflag.Tooltip = utils.StandardizeSpaces(h.ChildText(".kk-package-maindesc"))
		useflag.Name = utils.StandardizeSpaces(h.ChildText(".kk-package-name"))
	})

	collector.OnHTML(".tab-content tbody", func(h *colly.HTMLElement) {
		h.ForEach("tr", func(i int, h *colly.HTMLElement) {
			name := utils.StandardizeSpaces(h.ChildText("th"))

			if name == "" || name == useflag.Name {
				return
			}

			useflag.Packages = append(useflag.Packages, structs.PackageSearch{
				Group:       strings.Split(name, "/")[0],
				Package:     strings.Split(name, "/")[1],
				Url:         getBaseUrl() + h.DOM.Find("th a").AttrOr("href", "FAIL_PKG_URL"),
				Description: utils.StandardizeSpaces(h.ChildText("td")),
			})
		})
	})

	collector.Visit(GetUseFlagUrl(flag))
	return useflag
}
