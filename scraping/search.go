package scraping

import (
	"gentoo-packages-bot/structs"

	"github.com/gocolly/colly/v2"
)

func SearchPackage(packageQuery string) *[]structs.PackageSearch {
	pkgs := []structs.PackageSearch{}
	collector := colly.NewCollector()

	collector.OnHTML(".row > .col-12 > .panel > .list-group", func(h *colly.HTMLElement) {
		h.ForEach("a", func(i int, h *colly.HTMLElement) {
			pkgs = append(pkgs, structs.PackageSearch{
				Url:   getBaseUrl() + h.Attr("href"),
				Group: h.ChildText("h3 > .text-muted"),
			})
		})
	})

	collector.Visit(GetSearchQueryUrl(packageQuery))

	return &pkgs
}
