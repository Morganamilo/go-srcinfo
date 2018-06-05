package srcinfo

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	headerNone    = -2
	headerPkgbase = -1
)

type parser struct {
	headerType int
	srcinfo    *Srcinfo
}

func (psr *parser) currentPackage() (*Package, error) {
	if psr.headerType == headerNone {
		return nil, fmt.Errorf("Not in pkgbase or pkgname")
	} else if psr.headerType == headerPkgbase {
		return &psr.srcinfo.Package, nil
	} else {
		return &psr.srcinfo.Packages[psr.headerType], nil
	}
}

func (psr *parser) setValue(key, value string) error {
	pkgbase := &psr.srcinfo.PackageBase
	seenPkgnames := map[string]struct{}{}

	switch key {
	case "pkgbase":
		if psr.headerType != headerNone {
			return fmt.Errorf("key \"%s\" can not occur after a pkgbase or pkgname", key)
		}

		pkgbase.Pkgbase = value
		psr.headerType = headerPkgbase
		return nil
	case "pkgname":
		if psr.headerType == headerNone {
			return fmt.Errorf("key \"%s\" can not occur before pkgbase", key)
		}
		if _, ok := seenPkgnames[value]; ok {
			return fmt.Errorf("pkgname \"%s\" can not occur more than once", key)
		}
		seenPkgnames[value] = struct{}{}

		pkgbase.Pkgnames = append(pkgbase.Pkgnames, value)
		psr.srcinfo.Packages = append(psr.srcinfo.Packages, Package{})
		psr.headerType = len(pkgbase.Pkgnames) - 1
		return nil
	}

	if psr.headerType == headerNone {
		return fmt.Errorf("key \"%s\" can not occur before pkgbase or pkgname", key)
	}

	pkg, err := psr.currentPackage()
	if err != nil {
		return err
	}

	found := true

	// pkgbase only
	switch key {
	case "pkgver":
		pkgbase.Pkgver = value
	case "pkgrel":
		pkgbase.Pkgrel = value
	case "epoch":
		pkgbase.Epoch = value
	case "source":
		pkgbase.Source = append(pkgbase.Source, makeArchString(key, value))
	case "validpgpkeys":
		pkgbase.ValidPGPKeys = append(pkgbase.ValidPGPKeys, value)
	case "noextract":
		pkgbase.NoExtract = append(pkgbase.NoExtract, value)
	case "md5sums":
		pkgbase.MD5Sums = append(pkgbase.MD5Sums, makeArchString(key, value))
	case "sha1sums":
		pkgbase.SHA1Sums = append(pkgbase.SHA1Sums, makeArchString(key, value))
	case "sha224sums":
		pkgbase.SHA224Sums = append(pkgbase.SHA224Sums, makeArchString(key, value))
	case "sha256sums":
		pkgbase.SHA256Sums = append(pkgbase.SHA256Sums, makeArchString(key, value))
	case "sha384sums":
		pkgbase.SHA384Sums = append(pkgbase.SHA384Sums, makeArchString(key, value))
	case "sha512sums":
		pkgbase.SHA512Sums = append(pkgbase.SHA512Sums, makeArchString(key, value))
	case "makedepends":
		pkgbase.MakeDepends = append(pkgbase.MakeDepends, makeArchString(key, value))
	case "checkdepends":
		pkgbase.CheckDepends = append(pkgbase.CheckDepends, makeArchString(key, value))
	default:
		found = false
	}

	if found {
		if psr.headerType != headerPkgbase {
			return fmt.Errorf("key \"%s\" can not after pkgname", key)
		}

		return nil
	}

	// pkgbase or pkgname
	switch key {
	case "pkgdesc":
		pkg.Pkgdesc = value
	case "url":
		pkg.URL = value
	case "license":
		pkg.License = append(pkg.License, value)
	case "install":
		pkg.Install = value
	case "changelog":
		pkg.Changelog = value
	case "groups":
		pkg.Groups = append(pkg.Groups, value)
	case "arch":
		pkg.Arch = append(pkg.Arch, value)
	case "backup":
		pkg.Backup = append(pkg.Backup, value)
	case "depends":
		pkg.Depends = append(pkg.Depends, makeArchString(key, value))
	case "optdepends":
		pkg.OptDepends = append(pkg.OptDepends, makeArchString(key, value))
	case "conflicts":
		pkg.Conflicts = append(pkg.Conflicts, makeArchString(key, value))
	case "provides":
		pkg.Provides = append(pkg.Provides, makeArchString(key, value))
	case "replaces":
		pkg.Replaces = append(pkg.Replaces, makeArchString(key, value))
	case "options":
		pkg.Options = append(pkg.Options, value)
	default:
		return fmt.Errorf("Unknown key: \"%s\"", key)
	}

	return nil
}

// splitLine splits a key value string in the form of "key = value",
// whitespace being ignored. The key and the value is returned.
func (psr *parser) splitLine(line string) (string, string, error) {
	split := strings.SplitN(line, "=", 2)

	if len(split) != 2 {
		return "", "", fmt.Errorf("Line does not contain =")
	}

	key := strings.TrimSpace(split[0])
	value := strings.TrimSpace(split[1])

	if key == ""  {
		return "", "", fmt.Errorf("Key is empty", line)
	}

	if value == ""  {
		return "", "", fmt.Errorf("value is empty", line)
	}

	return key, value, nil
}

func parse(data string) (*Srcinfo, error) {
	psr := &parser{
		headerNone,
		&Srcinfo{},
	}

	lines := strings.Split(data, "\n")

	for n, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		key, value, err := psr.splitLine(line)
		if err != nil {
			return nil, Error(n, line, err.Error())
		}

		err = psr.setValue(key, value)
		if err != nil {
			return nil, Error(n, line, err.Error())
		}

	}

	return psr.srcinfo, nil

}

func getArchFromKey(key string) string {
	split := strings.SplitN(key, "_", 2)
	arch := ""
	if len(split) == 2 {
		arch = split[1]
	}

	return arch
}

func makeArchString(key, value string) ArchString {
	return ArchString{
		getArchFromKey(key),
		value,
	}
}
