package structs

type PackageSearch struct {
	Group       string `json:"group_name"`
	Package     string `json:"package_name"`
	Description string `json:"description"`
	Url         string `json:"url"`
}

type PackageVersion map[string]string // "stable", "testing", "alpha", and "unsupported"
type PackageVersions map[string]PackageVersion

type PackageFlags struct {
	Category string    `json:"category"` // "Local Use Flags", "Global Use Flags", "cpu_flags_arm (Use Expand)", "l10n (Use Expand)", etc..
	Flags    []UseFlag `json:"flags"`
}

type PackageMaintainer struct {
	Name  string `json:"name"`
	Url   string `json:"url"`
	Email string `json:"email"`
}

type Package struct {
	PackageSearch
	Versions   PackageVersions   `json:"versions"`
	Flags      PackageFlags      `json:"package_flags"`
	License    string            `json:"license"`
	Maintainer PackageMaintainer `json:"maintainer"`
}
