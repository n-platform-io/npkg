package types

type PackageBuild struct {
	PkgName         string      `hcl:"pkgname"`
	PkgVer          string      `hcl:"pkgver"`
	PkgRel          int         `hcl:"pkgrel"`
	BuildStyle      string      `hcl:"build_style"`
	MakeBuildArgs   []string    `hcl:"make_build_args"`
	HostMakeDepends []string    `hcl:"hostmakedepends"`
	PkgDesc         string      `hcl:"pkgdesc"`
	Maintainer      string      `hcl:"maintainer"`
	License         string      `hcl:"license"`
	URL             string      `hcl:"url"`
	Source          string      `hcl:"source"`
	Sha256          string      `hcl:"sha256"`
	BuildScript     BuildScript `hcl:"build_script,block"`
}

type BuildScript struct {
	Commands  [][]string `hcl:"commands"`
	Overrides Overrides  `hcl:"overrides,block"`
}

type Overrides struct {
	// Add fields here if you have specific overrides in the future
}
