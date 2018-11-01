package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	var typeName string
	var pkgPath string
	var outpath string
	gopath := os.Getenv("GOPATH")
	flag.StringVar(&typeName, "type", "byte", "type name for code generation")
	flag.StringVar(&pkgPath, "pkgpath", "builtin", "pkg path to import for getting your type, defaults to builtin")
	flag.StringVar(&outpath, "outpath", gopath+"/src/github.com/dennisfrancis/mta/stas/", "path where to write the output file")
	flag.Parse()

	infile := gopath + "/src/github.com/dennisfrancis/mta/stas/template.go"
	fp, err := os.Open(infile)
	if err != nil {
		fmt.Println("Error opening ", infile)
		return
	}
	defer fp.Close()
	parts := strings.Split(pkgPath, "/")
	pkgName := parts[len(parts)-1]
	importReq := true
	if pkgPath == "builtin" {
		importReq = false
	}
	typeNameTitle := strings.ToUpper(typeName[0:1]) + typeName[1:]
	outfile := outpath + strings.ToLower(typeName) + "_array.go"
	fout, err1 := os.Create(outfile)
	if err1 != nil {
		fmt.Println("Error creating ", outfile)
		return
	}
	defer fout.Close()
	importDone := false
	if importReq {
		typeName = pkgName + "." + typeName
	}
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		if !importDone {
			idx := strings.Index(line, "//\"typepackage\"")
			if idx >= 0 {
				importDone = true
				if !importReq {
					continue
				}
				line = line[:idx] + "\"" + pkgPath + "\""
			}
		}
		if strings.Contains(line, "go run") {
			continue
		}

		if strings.HasPrefix(line, "type ElementType struct{}") {
			continue
		}
		line = strings.Replace(line, "ElementTypeArray", typeNameTitle+"Array", -1)
		line = strings.Replace(line, "[]ElementType", "[]"+typeName, -1)
		fout.WriteString(line + "\n")
	}
	if err2 := scanner.Err(); err2 != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err2)
	}
}
