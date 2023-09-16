package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

var flagArity = flag.Int("arity", 1, "arity")

func main() {
	flag.Parse()

	returnTypes := []string{"None", "Error", "Result", "Value"}
	maxArity := *flagArity

	for _, ctx := range []bool{false, true} {
		for _, returnType := range returnTypes {
			for arity := 1; arity <= maxArity; arity++ {
				err := GenerateFile(ctx, returnType, arity)
				if err != nil {
					panic(err)
				}
			}
		}
	}

}

func GenerateFile(ctx bool, returnType string, arity int) error {
	var path string
	if ctx {
		path += "ctx_"
	}
	path += "func"
	switch returnType {
	case "None":
		path += ".go"
	case "Error":
		path += "_error.go"
	case "Result":
		path += "_result.go"
	case "Value":
		path += "_value.go"
	default:
		return fmt.Errorf("invalid return type: %s", returnType)
	}

	fd, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fd.Close()

	b, err := io.ReadAll(fd)
	if err != nil {
		return err
	}

	var arityCall strings.Builder
	var arityDecl strings.Builder
	var arityType strings.Builder
	for i := 0; i < arity; i++ {
		if i > 0 {
			arityCall.WriteString(", ")
			arityDecl.WriteString(", ")
			arityType.WriteString(", ")
		}
		arityCall.WriteString(fmt.Sprintf("p%d", i))
		arityDecl.WriteString(fmt.Sprintf("p%d P%d", i, i))
		arityType.WriteString(fmt.Sprintf("P%d", i))
	}

	augmented := regexp.MustCompile("Func([^t])").ReplaceAll(b, []byte(fmt.Sprintf("Func%d$1", arity)))
	if ctx {
		augmented = regexp.MustCompile(`f\(ctx\)`).ReplaceAll(augmented, []byte(fmt.Sprintf("f(ctx, %s)", arityCall.String())))
		augmented = regexp.MustCompile(`\(ctx context.Context\)`).ReplaceAll(augmented, []byte(fmt.Sprintf("(ctx context.Context, %s)", arityDecl.String())))
	} else {
		augmented = regexp.MustCompile(`f\(\)`).ReplaceAll(augmented, []byte("f("+arityCall.String()+")"))
		augmented = regexp.MustCompile(`(return func|Exec)\(\)`).ReplaceAll(augmented, []byte("${1}("+arityDecl.String()+")"))
		augmented = regexp.MustCompile(`(type.*) func\(\)`).ReplaceAll(augmented, []byte("$1 func("+arityDecl.String()+")"))
	}

	switch returnType {
	case "None", "Error":
		augmented = regexp.MustCompile(`(f|\)) (Ctx|)Func([0-9]+)(Error|)`).ReplaceAll(augmented, []byte("${1} ${2}Func${3}${4}["+arityType.String()+"]"))
		augmented = regexp.MustCompile(`type (Ctx|)Func([0-9]+)(Error|)`).ReplaceAll(augmented, []byte("type ${1}Func${2}${3}["+arityType.String()+" any]"))
	case "Value", "Result":
		augmented = regexp.MustCompile(`\[(R|T)\]`).ReplaceAll(augmented, []byte(fmt.Sprintf("[$1, %s]", arityType.String())))
		augmented = regexp.MustCompile(`\[(R|T) any\]`).ReplaceAll(augmented, []byte(fmt.Sprintf("[$1, %s any]", arityType.String())))
	}

	augmented = addCurrying(augmented, ctx, returnType, arity)

	newPath := fmt.Sprintf("%d_%s", arity, path)
	err = os.WriteFile(newPath, augmented, 0644)
	if err != nil {
		panic(err)
	}

	return nil
}

func addCurrying(b []byte, ctx bool, returnType string, arity int) []byte {
	for i := 0; i < arity; i++ {
		b = append(b, curryFromTo(ctx, returnType, arity, i)...)
	}
	return b
}

func curryFromTo(ctx bool, returnType string, from, to int) []byte {
	var arityCall strings.Builder
	var arityDeclCurry strings.Builder
	var arityTypeFrom strings.Builder

	if returnType == "Result" || returnType == "Value" {
		arityTypeFrom.WriteString("R, ")
	}

	if ctx {
		arityCall.WriteString("ctx, ")
	}

	for i := 0; i < from-to; i++ {
		if i > 0 {
			arityCall.WriteString(", ")
			arityDeclCurry.WriteString(", ")
			arityTypeFrom.WriteString(", ")
		}
		arityCall.WriteString(fmt.Sprintf("p%d", i))
		arityDeclCurry.WriteString(fmt.Sprintf("p%d P%d", i, i))
		arityTypeFrom.WriteString(fmt.Sprintf("P%d", i))
	}

	var arityDeclTo strings.Builder
	var arityTypeTo strings.Builder

	if returnType == "Result" || returnType == "Value" {
		arityTypeTo.WriteString("R")
		if to > 0 {
			arityTypeTo.WriteString(", ")
		}
	}

	if ctx {
		arityDeclTo.WriteString("ctx context.Context")
		if to > 0 {
			arityDeclTo.WriteString(", ")
		}
	}

	for i := from - to; i < from; i++ {
		arityCall.WriteString(", ")
		arityTypeFrom.WriteString(", ")
		if i > from-to {
			arityDeclTo.WriteString(", ")
			arityTypeTo.WriteString(", ")
		}
		arityCall.WriteString(fmt.Sprintf("p%d", i))
		arityDeclTo.WriteString(fmt.Sprintf("p%d P%d", i, i))
		arityTypeFrom.WriteString(fmt.Sprintf("P%d", i))
		arityTypeTo.WriteString(fmt.Sprintf("P%d", i))
	}

	var ctxPrefix string
	if ctx {
		ctxPrefix = "Ctx"
	}
	returnPrefix := returnType
	if returnType == "None" {
		returnPrefix = ""
	}

	var toArityPrefix string
	var toArityType string
	if to > 0 {
		toArityPrefix = fmt.Sprintf("%d", to)
		toArityType = fmt.Sprintf("[%s]", arityTypeTo.String())
	} else if returnType == "Result" || returnType == "Value" {
		toArityType = fmt.Sprintf("[%s]", arityTypeTo.String())
	}

	var returnDecl string
	switch returnType {
	case "None":
		returnDecl = ""
	case "Error":
		returnDecl = "error"
	case "Value":
		returnDecl = "R"
	case "Result":
		returnDecl = "(R, error)"
	}

	var call string
	switch returnType {
	case "None":
		call = fmt.Sprintf("f(%s)", arityCall.String())
	case "Error", "Value", "Result":
		call = fmt.Sprintf("return f(%s)", arityCall.String())
	}

	tmpl := fmt.Sprintf(`

func (f %sFunc%d%s[%s]) Curry%d(%s) %sFunc%s%s%s {
	return func(%s) %s {
		%s
	}
}
	`,
		ctxPrefix, from, returnPrefix, arityTypeFrom.String(), from-to, arityDeclCurry.String(), ctxPrefix, toArityPrefix, returnPrefix, toArityType,
		arityDeclTo.String(), returnDecl,
		call,
	)

	return []byte(tmpl)
}
