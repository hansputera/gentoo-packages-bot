package scraping

import (
	"gentoo-packages-bot/structs"
	"gentoo-packages-bot/utils"
	"strings"

	"github.com/gocolly/colly/v2"
)

func handleMetadata(dst *structs.Package, e *colly.HTMLElement) {
	name := utils.StandardizeSpaces(e.ChildText(".kk-metadata-key"))

	if strings.Contains(strings.ToLower(name), "license") {
		dst.License = utils.StandardizeSpaces(e.DOM.After(".kk-metadata-key").Text())
	} else {
		dst.Maintainer = structs.PackageMaintainer{
			Name:  e.DOM.After(".kk-metadata-key").Find("a").Eq(0).AttrOr("title", "-"),
			Url:   getBaseUrl() + e.DOM.After(".kk-metadata-key").Find("a").Eq(0).AttrOr("href", "FAIL_GET_USER"),
			Email: e.DOM.After(".kk-metadata-key").Find("a").Eq(1).AttrOr("href", "-"),
		}
	}
}

func GetPackage(group string, pkg string) *structs.Package {
	pkg_struct := structs.Package{}
	collector := colly.NewCollector()

	pkg_struct.Versions = make(structs.PackageVersions)

	// detail track
	collector.OnHTML(".kk-header-container", func(h *colly.HTMLElement) {
		pkg_struct.Group = strings.TrimSuffix(h.ChildText(".kk-package-cat"), "/")
		pkg_struct.Package = utils.StandardizeSpaces(h.ChildText(".kk-package-name"))
		pkg_struct.Url = GetPackageUrl(group, pkg)
	})

	// available versions track
	collector.OnHTML("tbody tr", func(h *colly.HTMLElement) {
		// skip header
		if h.Index == 0 {
			return
		}

		h.ForEach("td", func(_ int, versionEl *colly.HTMLElement) {
			if versionEl.DOM.HasClass("kk-version") {
				pkg_struct.Versions[utils.StandardizeSpaces(versionEl.DOM.Find("strong").Text())] = &structs.PackageVersion{
					EbuildUrl: versionEl.DOM.Find("strong a").AttrOr("href", "-"),
				}
			} else if versionEl.DOM.HasClass("kk-keyword") {
				matchs := strings.Split(
					utils.StandardizeSpaces(
						versionEl.Attr("title"),
					), " ",
				)
				if len(pkg_struct.Versions[matchs[0]].EbuildUrl) > 0 {
					pkg_struct.Versions[matchs[0]].ArchStatuses = append(
						pkg_struct.Versions[matchs[0]].ArchStatuses,
						matchs[len(matchs)-1]+"_"+matchs[2],
					)
				}
			}
		})
	})

	// metadata useflag track
	collector.OnHTML(".kk-useflag-group", func(h *colly.HTMLElement) {
		pkg_struct.Flags = append(pkg_struct.Flags, structs.PackageFlags{
			Category: utils.StandardizeSpaces(h.Text),
			Flags:    make([]structs.UseFlag, 0),
		})
	})

	// metadata useflag track
	collector.OnHTML("ul.kk-useflag-container", func(h *colly.HTMLElement) {
		if len(pkg_struct.Flags) >= h.Index {
			h.ForEach("li.kk-useflag a", func(_ int, flagel *colly.HTMLElement) {
				pkg_struct.Flags[h.Index].Flags = append(pkg_struct.Flags[h.Index].Flags, structs.UseFlag{
					Name:    utils.StandardizeSpaces(flagel.Text),
					Url:     getBaseUrl() + flagel.Attr("href"),
					Tooltip: flagel.Attr("title"),
				})
			})
		}
	})

	// another useflag metadata track
	collector.OnHTML("ul.kk-metadata-list", func(h *colly.HTMLElement) {
		h.ForEach("li.kk-metadata-item", func(_ int, mel *colly.HTMLElement) {
			// we skip [0] if 'len(pkg_struct.Flags) > 0', the flag is tracked
			if len(pkg_struct.Flags) > 0 && mel.Index == 0 {
				return
			}

			handleMetadata(&pkg_struct, mel)
		})
	})

	collector.Visit(GetPackageUrl(group, pkg))

	return &pkg_struct
}
