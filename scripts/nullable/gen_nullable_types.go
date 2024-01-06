package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	String  = "string"
	Int     = "int"
	Int8    = "int8"
	Int16   = "int16"
	Int32   = "int32"
	Int64   = "int64"
	Uint    = "uint"
	Uint8   = "uint8"
	Uint16  = "uint16"
	Uint32  = "uint32"
	Uint64  = "uint64"
	Float32 = "float32"
	Float64 = "float64"
	Bool    = "bool"
	Byte    = "byte"
	Time    = "time"
)

var TYPES = []string{
	String,
	Int,
	Int8,
	Int16,
	Int32,
	Int64,
	Uint,
	Uint8,
	Uint16,
	Uint32,
	Uint64,
	Float32,
	Float64,
	Bool,
	Byte,
	Time,
}

var (
	PATH_PROJECT_ROOT  = GetProjectRootDir()
	PATH_MOD_FILE      = PATH_PROJECT_ROOT + "/go.mod"
	TEMPLATE_PATH      = PATH_PROJECT_ROOT + "/scripts/nullable/" + TEMPLATE_FILE
	TEMPLATE_TEST_PATH = PATH_PROJECT_ROOT + "/scripts/nullable/" + TEMPLATE_TEST_FILE
)

const (
	ProjectRootEnvVar      = "PROJECT_ROOT"
	CurrentDirectoryEnvVar = "PWD"
	TEMPLATE_FILE          = "nullable_types.gotoml"
	TEMPLATE_TEST_FILE     = "nullable_types_test.gotoml"
)

type Nullable struct {
	Type string
}

func (n Nullable) Short() string {
	if len(n.Type) > 0 {
		return n.Type[0:1]
	}

	return ""
}

func (n Nullable) NullableType() string {
	return strings.ToUpper(n.Short()) + n.Type[1:]
}

func GetProjectRootDir() string {
	if val, ok := os.LookupEnv(ProjectRootEnvVar); ok {
		return val
	} else {
		pwd := os.Getenv(CurrentDirectoryEnvVar)
		return pwd
	}
}

func generate_types() {
	nullableTemplate := template.Must(template.New(TEMPLATE_FILE).ParseFiles(TEMPLATE_PATH))
	nullableTestTemplate := template.Must(template.New(TEMPLATE_TEST_FILE).ParseFiles(TEMPLATE_TEST_PATH))
	nullableTypes := make([]Nullable, 0)
	for _, t := range TYPES {
		nullableTypes = append(nullableTypes, Nullable{Type: t})
	}

	for _, nt := range nullableTypes {
		file, err := os.Create(filepath.Join(PATH_PROJECT_ROOT, nt.Type+".go"))
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		testfile, err := os.Create(filepath.Join(PATH_PROJECT_ROOT, nt.Type+"_test.go"))
		if err != nil {
			log.Fatal(err)
		}
		defer testfile.Close()

		if err := nullableTemplate.Execute(file, nt); err != nil {
			log.Fatal(err)
		}

		if err := nullableTestTemplate.Execute(testfile, nt); err != nil {
			log.Fatal(err)
		}
	}


}

func main() {
	generate_types()
}
