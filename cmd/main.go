package main

import (
	"flag"
	"fmt"
	"github.com/ttacon/goji-filesize/filesize"
	"os"
	"path/filepath"
)

var (
	// TODO(ttacon): rename pkg to target or something equally uselessly generic
	pkg      = flag.String("pkg", "", "the package to print the number of lines for")
	fileOnly = flag.Bool("f", false, "the given target is a file not a package")

	goPath = os.Getenv("GOPATH")
)

func main() {
	flag.Parse()

	if !*fileOnly {
		if err := filesize.DirLines(filepath.Join(goPath, *pkg)); err != nil {
			fmt.Println(
				"Failed to determine the number of lines of the files in the package ",
				*pkg)
		}
		return
	}

	if err := filesize.FileLines(*pkg); err != nil {
		fmt.Println("failed to count lines for file, ", *pkg)
	}
}
