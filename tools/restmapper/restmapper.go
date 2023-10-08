package main

import (
	_ "embed"
	"flag"
	"github.com/golang/glog"
	"github.com/opendatahub-io/model-registry/internal/model/library"
	"github.com/opendatahub-io/model-registry/tools/restmapper/internal"
	"text/template"
)

// OK to hardcode path here since this is an internal tool
// the generate reference MUST be in main.go
var libraryDirs = []string{"./config/metadata-library"}

//go:embed internal/library-template.gotmpl
var schemaTemplateStr string

//go:embed internal/converter-template.gotmpl
var converterTemplateStr string

//go:embed internal/gqlgen-template.gotmpl
var gqlgenTemplateStr string

func main() {
	// default to logging to stderr
	flag.Set("logtostderr", "true")
	defer glog.Flush()

	libs, err := library.LoadLibraries(libraryDirs)
	if err != nil {
		glog.Exitf("error loading libraries from %s: %v", libraryDirs, err)
	}

	m := &internal.Mapper{}
	m.Libraries = libs

	// load templates
	m.SchemaTemplate, err = template.New("schemaMapper").
		Parse(schemaTemplateStr)
	if err != nil {
		glog.Exitf("error parsing schema template: %v", err)
	}
	m.ConverterTemplate, err = template.New("converterMapper").
		Parse(converterTemplateStr)
	if err != nil {
		glog.Exitf("error parsing converter template: %v", err)
	}
	m.GqlgenTemplate, err = template.New("resolverMapper").
		Parse(gqlgenTemplateStr)
	if err != nil {
		glog.Exitf("error parsing gqlgen template: %v", err)
	}

	m.MapLibraries()
}
