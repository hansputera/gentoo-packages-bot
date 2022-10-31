package structs

type PackageSearch struct {
	Group       string `json:"group_name"`
	Package     string `json:"package_name"`
	Description string `json:"description"`
	Url         string `json:"url"`
}

type PackageVersion struct {
	EbuildUrl    string   `json:"ebuild_url"`
	ArchStatuses []string `json:"arch_statuses"` // (e.g "amd64_stable", "x86_testing", "ppc_unknown")
}

type PackageVersions map[string]*PackageVersion

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
	Flags      []PackageFlags    `json:"package_flags"`
	License    string            `json:"license"`
	Maintainer PackageMaintainer `json:"maintainer"`
}
