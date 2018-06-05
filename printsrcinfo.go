package srcinfo

import (
	"bytes"
)

func appendHeader(buffer *bytes.Buffer, key string, value string) {
	if value == "" {
		return
	}

	buffer.WriteString(key + " = " + value + "\n")
}

func appendValue(buffer *bytes.Buffer, key string, value string) {
	if value == "" {
		return
	}

	buffer.WriteString("\t" + key + " = " + value + "\n")
}

func appendMultiValue(buffer *bytes.Buffer, key string, values []string) {
	if len(values) == 0 {
		return
	}

	for _, value := range values {
		buffer.WriteString("\t" + key + " = " + value + "\n")
	}
}

func appendMultiArchValue(buffer *bytes.Buffer, key string, values []ArchString) {
	if len(values) == 0 {
		return
	}

	for _, value := range values {
		if value.Arch == "" {
			buffer.WriteString("\t" + key + " = " + value.Value + "\n")
		} else {
			buffer.WriteString("\t" + key + "_" + value.Arch + " = " + value.Value + "\n")
		}
	}
}

func appendArchValue(buffer *bytes.Buffer, key string, value ArchString) {
	if value.Value == "" {
		return
	}

	if value.Arch == "" {
		if value.Arch == "" {
			buffer.WriteString("\t" + key + " = " + value.Value + "\n")
		} else {
			buffer.WriteString("\t" + key + "_" + value.Arch + " = " + value.Value + "\n")
		}
	}
}

func (si *Srcinfo) String() string {
	var buffer bytes.Buffer

	appendHeader(&buffer, "pkgbase", si.Pkgbase)

	appendValue(&buffer, "pkgdesc", si.Pkgdesc)
	appendValue(&buffer, "pkgver", si.Pkgver)
	appendValue(&buffer, "pkgrel", si.Pkgrel)
	appendValue(&buffer, "epoch", si.Epoch)
	appendValue(&buffer, "url", si.URL)
	appendValue(&buffer, "install", si.Install)
	appendValue(&buffer, "changelog", si.Changelog)
	appendMultiValue(&buffer, "arch", si.Arch)
	appendMultiValue(&buffer, "groups", si.Groups)
	appendMultiValue(&buffer, "license", si.License)
	appendMultiArchValue(&buffer, "checkDepends", si.CheckDepends)
	appendMultiArchValue(&buffer, "makedepends", si.MakeDepends)
	appendMultiArchValue(&buffer, "depends", si.Depends)
	appendMultiArchValue(&buffer, "optdepends", si.OptDepends)
	appendMultiArchValue(&buffer, "provides", si.Provides)
	appendMultiArchValue(&buffer, "conflicts", si.Conflicts)
	appendMultiArchValue(&buffer, "replaces", si.Replaces)
	appendMultiValue(&buffer, "noextract", si.NoExtract)
	appendMultiValue(&buffer, "options", si.Options)
	appendMultiValue(&buffer, "backup", si.Backup)
	appendMultiValue(&buffer, "validpgpkeys", si.ValidPGPKeys)
	appendMultiArchValue(&buffer, "source", si.Source)
	appendMultiArchValue(&buffer, "md5sums", si.MD5Sums)
	appendMultiArchValue(&buffer, "sha1sums", si.SHA1Sums)
	appendMultiArchValue(&buffer, "sha224sums", si.SHA224Sums)
	appendMultiArchValue(&buffer, "sha256sums", si.SHA256Sums)
	appendMultiArchValue(&buffer, "sha384sums", si.SHA384Sums)
	appendMultiArchValue(&buffer, "sha512sums", si.SHA512Sums)

	for n, pkg := range si.Packages {
		appendHeader(&buffer, "\npkgname", si.Pkgnames[n])

		appendValue(&buffer, "pkgdesc", pkg.Pkgdesc)
		appendValue(&buffer, "url", pkg.URL)
		appendValue(&buffer, "install", pkg.Install)
		appendValue(&buffer, "changelog", pkg.Changelog)
		appendMultiValue(&buffer, "arch", pkg.Arch)
		appendMultiValue(&buffer, "groups", pkg.Groups)
		appendMultiValue(&buffer, "license", pkg.License)
		appendMultiArchValue(&buffer, "depends", pkg.Depends)
		appendMultiArchValue(&buffer, "optdepends", pkg.OptDepends)
		appendMultiArchValue(&buffer, "provides", pkg.Provides)
		appendMultiArchValue(&buffer, "conflicts", pkg.Conflicts)
		appendMultiArchValue(&buffer, "replaces", pkg.Replaces)
		appendMultiValue(&buffer, "options", pkg.Options)
		appendMultiValue(&buffer, "backup", pkg.Backup)
	}

	return buffer.String()
}

// Order: global
// pkgdesc
// pkgver
// pkgrel
// epoch
// url
// install
// changelog
// arch
// groups
// license
// checkdepends
// makedepends
// depends
// optdepends
// provides
// conflicts
// replaces
// noextract
// options
// backup
// source
// validpgpkeys
// md5suns
// sha1sums
// sha224sums
// sha256sums
// sha384sums
// sha512sums
//
// Order: package
// pkgdesc
// url
// install
// changelog
// arch
// groups
// license
// checkdepends
// depends
// optdepends
// provides
// conflicts
// replaces
// options
// backup
