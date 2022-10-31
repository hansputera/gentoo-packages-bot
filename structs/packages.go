package structs

type PackageSearch struct {
	Group       string `json:"group_name"`
	Package     string `json:"package_name"`
	Description string `json:"description"`
	Url         string `json:"url"`
}
