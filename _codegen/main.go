package main

import (
	"fmt"
	"go/ast"
	"go/build"
	"io"
	"os"
	"strings"

	"golang.org/x/tools/go/loader"
)

const (
	pkgName     = "unix"
	pkgFullName = "golang.org/x/sys/unix"
)

func main() {
	w := os.Stdout
	defer w.Close()
	if err := xmain(w); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func xmain(w io.Writer) error {
	prog, _, err := loadProgram()
	if err != nil {
		return err
	}
	if err = processProgram(w, prog); err != nil {
		return err
	}
	return nil
}

func loadProgram() (*loader.Program, []string, error) {
	var conf loader.Config
	conf.Build = &build.Default
	args := []string{pkgFullName}
	const includeTests = false
	rest, err := conf.FromArgs(args, includeTests)
	if err != nil {
		return nil, rest, err
	}
	prog, err := conf.Load()
	return prog, rest, err
}

func processProgram(w io.Writer, prog *loader.Program) error {
	fmt.Fprintf(w, "// Automatically generated file. DO NOT EDIT MANUALLY.\n")
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "package %s\n", pkgName)
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "import (\n")
	// TODO: Do not hard-code deps here; parse *loader.program
	fmt.Fprintf(w, "\t\"syscall\"\n")
	fmt.Fprintf(w, "\t\"time\"\n")
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "\t%s \"%s\"\n", pkgName, pkgFullName)
	fmt.Fprintf(w, ")\n")
	fmt.Fprintf(w, "\n")
	for _, pkgInfo := range prog.InitialPackages() {
		for _, file := range pkgInfo.Files {
			ast.FileExports(file)
			if err := processFile(w, file); err != nil {
				return err
			}
		}
	}
	return nil
}

func processFile(w io.Writer, file *ast.File) error {
	for _, xdecl := range file.Decls {
		switch decl := xdecl.(type) {
		case *ast.FuncDecl:
			if decl.Recv != nil {
				continue
			}
			if err := processFuncDecl(w, decl); err != nil {
				return err
			}
			fmt.Fprintf(w, "\n")
		}
	}
	return nil
}

func processFuncDecl(w io.Writer, decl *ast.FuncDecl) error {
	if decl.Recv != nil {
		panic("receiver is not expected here")
	}
	resNames, anonIdxTypeMap, resErrIdx := getFuncResultNames(decl.Type.Results, "_v")
	needsWrap := false
	if resErrIdx >= 0 {
		if decl.Name.Name == "Close" {
			// TODO: is this true for non-Linux?
			fmt.Fprintf(w, "// %s is an alias of %s.%s. Does NOT retry on EINTR. See close(2) for the reason.\n", decl.Name.Name, pkgFullName, decl.Name.Name)
		} else {
			fmt.Fprintf(w, "// %s is an alias of %s.%s, wrapped to automatically retry on EINTR.\n", decl.Name.Name, pkgFullName, decl.Name.Name)
			needsWrap = true
		}
	} else {
		fmt.Fprintf(w, "// %s is an alias of %s.%s.\n", decl.Name.Name, pkgFullName, decl.Name.Name)
	}
	fmt.Fprintf(w, "%s {\n", formatFuncDecl(decl)) // BEGIN function
	if needsWrap {
		for i, t := range anonIdxTypeMap {
			fmt.Fprintf(w, "\tvar (\n") // BEGIN var
			fmt.Fprintf(w, "\t\t%s %s\n", resNames[i], t)
			fmt.Fprintf(w, "\t)\n") // END var
		}
		fmt.Fprintf(w, "\tfor {\n") // BEGIN for
		fmt.Fprintf(w, "\t\t%s = %s.%s(%s)\n", strings.Join(resNames, ", "), pkgName, decl.Name.Name, formatFuncParams(decl.Type.Params, false))
		fmt.Fprintf(w, "\t\tif %s != syscall.EINTR {\n", resNames[resErrIdx])
		fmt.Fprintf(w, "\t\t\tbreak\n")
		fmt.Fprintf(w, "\t\t}\n")
		fmt.Fprintf(w, "\t}\n") // END for
		if len(anonIdxTypeMap) == 0 {
			fmt.Fprintf(w, "\treturn\n")
		} else {
			fmt.Fprintf(w, "\treturn %s\n", strings.Join(resNames, ", "))
		}
	} else {
		call := fmt.Sprintf("%s.%s(%s)", pkgName, decl.Name.Name, formatFuncParams(decl.Type.Params, false))
		if decl.Type.Results != nil {
			fmt.Fprintf(w, "\treturn %s\n", call)
		} else {
			fmt.Fprintf(w, "\t%s\n", call)
		}
	}
	fmt.Fprintf(w, "}\n") // END function
	return nil
}

func formatType(typ ast.Expr) string {
	switch t := typ.(type) {
	case nil:
		return ""
	case *ast.Ident:
		if t.IsExported() {
			return fmt.Sprintf("%s.%s", pkgName, t.Name)
		}
		return t.Name
	case *ast.SelectorExpr:
		return fmt.Sprintf("%s.%s", formatType(t.X), t.Sel.Name)
	case *ast.StarExpr:
		return fmt.Sprintf("*%s", formatType(t.X))
	case *ast.ArrayType:
		return fmt.Sprintf("[%s]%s", formatType(t.Len), formatType(t.Elt))
	case *ast.Ellipsis:
		return formatType(t.Elt)
	case *ast.FuncType:
		return fmt.Sprintf("func(%s)%s", formatFuncParams(t.Params, true), formatFuncResults(t.Results, true))
	case *ast.MapType:
		return fmt.Sprintf("map[%s]%s", formatType(t.Key), formatType(t.Value))
	case *ast.ChanType:
		// FIXME
		panic(fmt.Errorf("unsupported chan type %#v", t))
	case *ast.BasicLit:
		return t.Value
	default:
		panic(fmt.Errorf("unsupported type %#v", t))
	}
}

func formatFields(fields *ast.FieldList, withTypes bool) string {
	s := ""
	for i, field := range fields.List {
		for j, name := range field.Names {
			s += name.Name
			if j != len(field.Names)-1 {
				s += ", "
			}
		}
		if withTypes {
			s = strings.TrimSpace(s)
			s += " " + formatType(field.Type)
		}
		if i != len(fields.List)-1 {
			s += ", "
		}
	}
	return strings.TrimSpace(s)
}

func formatFuncParams(fields *ast.FieldList, withTypes bool) string {
	return formatFields(fields, withTypes)
}

func formatFuncResults(fields *ast.FieldList, withTypes bool) string {
	s := ""
	if fields != nil {
		s += " "
		needsPar := true
		if len(fields.List) == 0 {
			needsPar = false
		} else if len(fields.List) == 1 && len(fields.List[0].Names) == 0 {
			needsPar = false
		}
		if needsPar {
			s += "("
		}
		s += formatFields(fields, withTypes)
		if needsPar {
			s += ")"
		}
	}
	return s
}

func getFuncResultNames(fields *ast.FieldList, prefixForAnon string) (ss []string, anonIdxTypeMap map[int]string, errIdx int) {
	errIdx = -1
	if fields == nil {
		return
	}
	anonIdxTypeMap = make(map[int]string)
	i := 0
	for _, field := range fields.List {
		switch t := field.Type.(type) {
		case *ast.Ident:
			if t.Name == "error" {
				errIdx = i
			}
		case *ast.SelectorExpr:
			if formatType(t) == "syscall.Errno" {
				errIdx = i
			}
		}
		if len(field.Names) > 0 {
			for _, name := range field.Names {
				ss = append(ss, name.Name)
				i++
			}
		} else {
			ss = append(ss, fmt.Sprintf("%s%d", prefixForAnon, i))
			anonIdxTypeMap[i] = formatType(field.Type)
			i++
		}
	}
	return
}

func formatFuncDecl(decl *ast.FuncDecl) string {
	s := "func "
	if decl.Recv != nil {
		panic("receiver is not expected here")
	}
	const withTypes = true
	s += fmt.Sprintf("%s(%s)", decl.Name.Name, formatFuncParams(decl.Type.Params, withTypes))
	s += formatFuncResults(decl.Type.Results, withTypes)
	return s
}
