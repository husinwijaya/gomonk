package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"gomonk/josharianimpl"
	"log"
	"os"
	"path/filepath"
)

const implSuffix = "Impl"

func main() {
	if len(os.Args) < 4 {
		log.Fatal("missing argument: <go file that contains the interface> <interface name> <output file> [package name]")
	}
	fileSet := token.NewFileSet()
	targetInterface := os.Args[2]
	outputFile := os.Args[3]
	if _, err := os.Stat(filepath.Join(outputFile, "..")); os.IsNotExist(err) {
		os.MkdirAll(filepath.Join(outputFile, ".."), 0700)
	}
	output, err := os.Create(outputFile)
	if err != nil {
		log.Fatal("cannot create file: " + outputFile)
	}
	defer output.Close()
	sourceFile, err := parser.ParseFile(fileSet, os.Args[1], nil, parser.AllErrors)
	packageName := sourceFile.Name.Name
	if len(os.Args) > 4 {
		sourceFile.Name.Name = os.Args[4]
	}
	if err != nil {
		log.Fatal("error reading file", err)
	}
	ast.Inspect(sourceFile, interfaceInspector(targetInterface))
	printer.Fprint(output, fileSet, sourceFile)
	originalArgs := make([]string, len(os.Args))
	copy(originalArgs, os.Args)
	originalStdout := os.Stdout
	os.Stdout = output
	callJosharianimpl(packageName, targetInterface)
	os.Stdout = originalStdout
	output.Sync()
}

func callJosharianimpl(packageName, targetInterface string) {
	os.Args = []string{os.Args[0], "i *" + targetInterface + implSuffix, packageName + "." + targetInterface}
	fmt.Println()
	josharianimpl.Main()
}

func interfaceInspector(interfaceName string) func(n ast.Node) bool {
	return func(n ast.Node) bool {
		switch ts := n.(type) {
		// find variable declarations
		case *ast.TypeSpec:
			// which are public
			if ts.Name.IsExported() {
				switch ts.Type.(type) {
				// and are interfaces
				case *ast.InterfaceType:
					// check if interface name match as intended
					if ts.Name.Name == interfaceName {
						ts.Name.Name += implSuffix
						convertToStruct(ts)
						// stop traverse
						return false
					}
				}
			}
		}
		return true
	}
}

func convertToStruct(ts *ast.TypeSpec) {
	structType := new(ast.StructType)
	interfaceType := ts.Type.(*ast.InterfaceType)
	structType.Fields = interfaceType.Methods
	structType.Struct = interfaceType.Interface
	for _, field := range structType.Fields.List {
		field.Names[0].Name = josharianimpl.FuncPrefix + field.Names[0].Name
	}
	ts.Type = structType
}
