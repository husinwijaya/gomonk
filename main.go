package main

import (
	"github.com/dave/dst/decorator"
	"github.com/dave/jennifer/jen"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("missing argument: <go file that contains the interface> <interface name> <output file>")
	}
	inspector := InterfaceExtractor{name: os.Args[2], fileSet: token.NewFileSet()}
	sourceFile, err := parser.ParseFile(inspector.fileSet, os.Args[1], nil, parser.AllErrors)
	if err != nil {
		log.Fatal("error reading file", err)
	}
	ast.Inspect(sourceFile, inspector.inspect)
	genFile := jen.NewFile(sourceFile.Name.String())
	genFile.Type().Id(inspector.name + "Monk").Struct()
	inspector.generateFuncField()
	log.Println(genFile, inspector)
}

type InterfaceExtractor struct {
	name      string
	ifaceSpec *ast.TypeSpec
	ifaceType *ast.InterfaceType
	fileSet   *token.FileSet
}

func (this *InterfaceExtractor) inspect(n ast.Node) bool {
	switch ts := n.(type) {
	// find variable declarations
	case *ast.TypeSpec:
		// which are public
		if ts.Name.IsExported() {
			switch it := ts.Type.(type) {
			// and are interfaces
			case *ast.InterfaceType:
				// check if interface name match as intended
				if ts.Name.Name == this.name {
					this.ifaceSpec = ts
					this.ifaceType = it
					// stop traverse
					return false
				}
			}
		}
	}
	return true
}

func (this *InterfaceExtractor) generateFuncField() (code []jen.Code) {
	decorated, err := decorator.Decorate(this.fileSet, this.ifaceType)
	if err != nil {
		log.Fatal("error decorate interface", err)
	}
	log.Println(decorated)
	printer.Fprint(os.Stdout, this.fileSet, this.ifaceSpec)
	return nil
}
