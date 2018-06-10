package srcinfo

import (
	"testing"
)

func TestCurrentPackage(t *testing.T) {
	srcinfo := &Srcinfo{}
	splitpkg := &Package{}
	psr := &parser{srcinfo, make(map[string]struct{})}

	_, err := psr.currentPackage()
	if err == nil {
		t.Errorf("currentPackage should because of no header but didn't")
	}

	psr.srcinfo.Pkgbase = "foo"

	pkg, err := psr.currentPackage()
	if err != nil {
		t.Error(err)
	} else if pkg != &srcinfo.Package {
		t.Errorf("currentPackage does not return srcinfo.Package")
	}

	srcinfo.Packages = append(srcinfo.Packages, *splitpkg)

	pkg, err = psr.currentPackage()
	if err != nil {
		t.Error(err)
	} else if pkg != &srcinfo.Packages[0] {
		t.Errorf("currentPackage does not return srcinfo.Packages[0]")
	}
}

func TestSplitPair(t *testing.T) {
	const keyGood string = "a"
	const valueGood string = "b"

	inputGood := [...]string{
		"a = b",
		" a = b ",
		"a=b",
		"\ta\t=\tb\t",
	}

	inputBad := [...]string{
		"",
		"=",
		//"a=",
		"=b",
	}

	for _, str := range inputGood {
		key, value, err := splitPair(str)
		if err != nil {
			t.Errorf("failed to split line \"%s\": %s", str, err)
		}

		if key != keyGood || value != valueGood {
			t.Errorf("failed to split line \"%s\": expected key=%s value=%s, got key=%s value=%s", str, keyGood, valueGood, key, value)
		}
	}

	for _, str := range inputBad {
		_, _, err := splitPair(str)
		if err == nil {
			t.Errorf("Split line should have errored but didn't: %s", str)
		}
	}
}

func TestSetField(t *testing.T) {
	srcinfo := &Srcinfo{}
	psr := &parser{srcinfo, make(map[string]struct{})}

	err := psr.setField("install", "foo")
	if err == nil {
		t.Errorf("setField should have errored due to no header but did not")
	}
}
