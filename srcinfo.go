package srcinfo

import (
	"fmt"
)

type ArchString struct {
	Arch  string
	Value string
}

type Package struct {
	Pkgdesc    string
	Arch       []string
	URL        string
	License    []string
	Groups     []string
	Depends    []ArchString
	OptDepends []ArchString
	Provides   []ArchString
	Conflicts  []ArchString
	Replaces   []ArchString
	Backup     []string
	Options    []string
	Install    string
	Changelog  string
}

type PackageBase struct {
	Pkgbase      string
	Pkgnames     []string
	Pkgver       string
	Pkgrel       string
	Epoch        string
	Source       []ArchString
	ValidPGPKeys []string
	NoExtract    []string
	MD5Sums      []ArchString
	SHA1Sums     []ArchString
	SHA224Sums   []ArchString
	SHA256Sums   []ArchString
	SHA384Sums   []ArchString
	SHA512Sums   []ArchString
	MakeDepends  []ArchString
	CheckDepends []ArchString
}

type Srcinfo struct {
	PackageBase
	Package
	Packages []Package
}

func (si *Srcinfo) SplitPackage(pkgname string) (*Package, error) {
	for n, name := range si.Pkgnames {
		if name == pkgname {
			return mergeSplitPackage(&si.Package, &si.Packages[n]), nil
		}
	}

	return nil, fmt.Errorf("Package \"%s\" is not part of this package base", pkgname)
}

func mergeSplitPackage(base, split *Package) *Package {
	pkg := &Package{}
	*pkg = *base

	if split.Pkgdesc != "" {
		pkg.Pkgdesc = split.Pkgdesc
	}

	if len(split.Arch) != 0 {
		pkg.Arch = split.Arch
	}

	if split.URL != "" {
		pkg.URL = split.URL
	}

	if len(split.License) != 0 {
		pkg.License = split.License
	}

	if len(split.Groups) != 0 {
		pkg.Groups = split.Groups
	}

	if len(split.License) != 0 {
		pkg.License = split.License
	}

	if len(split.Depends) != 0 {
		pkg.Depends = split.Depends
	}

	if len(split.OptDepends) != 0 {
		pkg.OptDepends = split.OptDepends
	}

	if len(split.Provides) != 0 {
		pkg.Provides = split.Provides
	}

	if len(split.Conflicts) != 0 {
		pkg.Conflicts = split.Conflicts
	}

	if len(split.Replaces) != 0 {
		pkg.Replaces = split.Replaces
	}

	if len(split.Backup) != 0 {
		pkg.Backup = split.Backup
	}

	if len(split.Options) != 0 {
		pkg.Options = split.Options
	}

	if split.Changelog != "" {
		pkg.Changelog = split.Changelog
	}

	if split.Install != "" {
		pkg.Install = split.Install
	}

	return pkg
}
