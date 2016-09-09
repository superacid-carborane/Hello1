package Hello

import (
	"flag"
	"path/filepath"
)

var (
	path = flag.String("path", "", "The path to the file")
)

func sumFunc() string {
	flag.Parse()
	return *path
}

func hello(s string) string {
	outpud := filepath.Ext(s)
	return outpud
}
