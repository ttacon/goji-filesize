package filesize

import (
	"fmt"
	"go/parser"
	"go/token"
	"path"
)

// TODO(ttacon): refactor to return formatted lines to print,
// because only printing some (in case of DirsLines) is gross
func DirsLines(dirs []string) error {
	for _, dir := range dirs {
		if err := DirLines(dir); err != nil {
			return err
		}
	}
	return nil
}

func DirLines(dir string) error {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, nil, 0)
	if err != nil {
		return err
	}

	for _, pkg := range pkgs {
		for fName, file := range pkg.Files {
			fmt.Printf(
				"%s: %d\n",
				path.Base(fName),
				fset.File(file.Pos()).LineCount())
		}
	}
	return nil
}

func FileLines(fName string) error {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, fName, nil, 0)
	if err != nil {
		return err
	}
	fmt.Printf("%s: %d\n", path.Base(fName), fset.File(file.Pos()).LineCount())
	return nil
}
