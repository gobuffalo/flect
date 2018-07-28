package name

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Package will attempt to return a package version of the name
//	$GOPATH/src/foo/bar = foo/bar
//	$GOPATH\src\foo\bar = foo/bar
//	foo/bar = foo/bar
func Package(s string) string {
	return New(s).Package().String()
}

func (i Ident) Package() Ident {
	gp := goPath()
	s := i.Original
	slash := string(filepath.Separator)
	trims := []string{gp, slash, "src", slash}
	for _, pre := range trims {
		s = strings.TrimPrefix(s, pre)
	}
	return New(strings.Replace(s, "\\", "/", -1))
}

func goPath() string {
	cmd := exec.Command("go", "env", "GOPATH")
	b, err := cmd.CombinedOutput()
	if err != nil {
		return filepath.Join(os.Getenv("HOME"), "go")
	}
	return strings.TrimSpace(string(b))
}
