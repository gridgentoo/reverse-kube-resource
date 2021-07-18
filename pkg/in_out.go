package pkg

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/tools/go/ast/astutil"
)

func checkFatal(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%w", err)
		os.Exit(1)
	}
}

func printImports(pkg string, obj []Import, buf *bytes.Buffer) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", fmt.Sprintf("package %v", pkg), 0)
	checkFatal(err)
	for _, o := range obj {
		astutil.AddNamedImport(fset, f, o.name, o.path)
	}

	ast.SortImports(fset, f)
	printer.Fprint(buf, fset, f)
}

func printLines(rawVars []RawVar, buf *bytes.Buffer) {
	if len(rawVars) == 0 {
		return
	}
	single := false
	if len(rawVars) == 1 && len(rawVars[0].helpers) == 0 {
		single = true
		fmt.Fprintf(buf, "var ")
	} else {
		fmt.Fprintln(buf, "var (")
	}
	for _, o := range rawVars {
		if !single {
			fmt.Fprintf(buf, "// %v %q\n", o.kind, o.name)
		}
		for _, l := range o.helpers {
			fmt.Fprintln(buf, l)
		}
		if len(o.helpers) != 0 {
			fmt.Fprintln(buf, "")
		}
		for _, l := range o.lines {
			fmt.Fprintln(buf, l)
		}
		fmt.Fprintln(buf, "")
	}
	if !single {
		fmt.Fprintln(buf, ")")
	}
}

func Print(pkg, boilerplate string, imports []Import, vars []RawVar) {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "// Code generated. DO NOT EDIT!\n\n")
	printImports(pkg, imports, &buf)
	printLines(vars, &buf)
	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Printf("%v", string(buf.Bytes()))
		checkFatal(err)
	}
	fmt.Printf("%v", string(formatted))
}

func readFile(path string) (all [][]byte) {
	data, err := os.ReadFile(path)
	checkFatal(err)
	split := strings.Split(string(data), "---")
	for i, _ := range split {
		if len(split[i]) != 0 {
			all = append(all, []byte(split[i]))
		}
	}
	return
}

func read(path string) (all [][]byte) {
	fi, err := os.Stat(path)
	checkFatal(err)
	if fi.IsDir() {
		err := filepath.Walk(path, func(p string, i fs.FileInfo, err error) error {
			pl := strings.ToLower(p)
			if !(strings.HasSuffix(pl, ".yaml") || strings.HasSuffix(pl, ".yml")) || i.IsDir() {
				return nil
			}
			all = append(all, readFile(p)...)
			return nil
		})
		checkFatal(err)
	} else {
		all = readFile(path)
	}
	return all
}

func ReadInput(path string) (objs []object) {
	d := read(path)
	for _, data := range d {
		objs = append(objs, object{
			rt: getRuntimeObject(data),
			un: getUnstructuredObject(data),
		})
	}
	return
}
