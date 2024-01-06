package main

import (
	"log"
	"os"
	"path/filepath"
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
}

var (
	PATH_PROJECT_ROOT = GetProjectRootDir()
	PATH_MOD_FILE     = PATH_PROJECT_ROOT + "/go.mod"
	TEMPLATE_PATH     = PATH_PROJECT_ROOT + "/scripts/nullable/" + TEMPLATE_FILE
)

const (
	ProjectRootEnvVar      = "PROJECT_ROOT"
	CurrentDirectoryEnvVar = "PWD"
	TEMPLATE_FILE          = "nullable_types.gotmpl"
)

type Nullable struct {
	NullableType string
}

func GetProjectRootDir() string {
	if val, ok := os.LookupEnv(ProjectRootEnvVar); ok {
		return val
	} else {
		pwd := os.Getenv(CurrentDirectoryEnvVar)
		return pwd
	}
}

func main() {
	nullableTemplate := template.Must(template.New(TEMPLATE_FILE).ParseFiles(TEMPLATE_PATH))
	nullableTypes := make([]Nullable, 0)
	for _, t := range TYPES {
		nullableTypes = append(nullableTypes, Nullable{NullableType: t})
	}

	for _, nt := range nullableTypes {
		file, err := os.Create(filepath.Join(PATH_PROJECT_ROOT, "nullable", nt.NullableType+".go"))
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()

		if err := nullableTemplate.Execute(file, nt); err != nil {
			log.Fatal(err)
		}
	}

}
