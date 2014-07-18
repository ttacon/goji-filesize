package filesize

import (
	"fmt"
	"go/parser"
	"go/token"
)

func DirLines(dir string) error {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, nil, 0)
	if err != nil {
		return err
	}

	for _, pkg := range pkgs {
		for fName, file := range pkg.Files {
			fmt.Printf("%s: %d\n", fName, fset.File(file.Pos()).LineCount())
		}
	}
	return nil
}

func FileLines(fName string) error {
	return nil
}
