package scraping

import (
	"net/url"
)

func getBaseUrl() string {
	return "https://packages.gentoo.org"
}

func GetSearchQueryUrl(query string) string {
	return getBaseUrl() + "/search?" + url.Values{
		"q": {query},
	}.Encode()
}

func GetPackageUrl(group string, pkg string) string {
	return getBaseUrl() + "/packages/" + group + "/" + pkg
}

func GetUseFlagUrl(flag string) string {
	return getBaseUrl() + "/useflags/" + flag
}
