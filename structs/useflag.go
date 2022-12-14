package structs

type UseFlag struct {
	Url     string `json:"url"`
	Name    string `json:"flag"`
	Tooltip string `json:"tooltip"` // or-description
}

type UseFlagDetail struct {
	UseFlag
	Packages []PackageSearch `json:"packages"`
}
