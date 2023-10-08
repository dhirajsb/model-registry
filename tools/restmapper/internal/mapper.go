package internal

import (
	"bufio"
	"fmt"
	"github.com/golang/glog"
	"github.com/opendatahub-io/model-registry/internal/model/library"
	"github.com/rogpeppe/go-internal/fmtsort"
	"os"
	"path"
	"reflect"
	"strings"
	"sync"
	"text/template"
)

var gqlGenFileName = "./gqlgen.yml"

type Mapper struct {
	SchemaTemplate    *template.Template
	ConverterTemplate *template.Template
	GqlgenTemplate    *template.Template
	Libraries         map[string]*library.MetadataLibrary

	wg     *sync.WaitGroup
	failed bool
}

func (m *Mapper) MapLibraries() {
	// set up wait jobs
	m.wg = &sync.WaitGroup{}
	// number of tasks to wait for
	// schema and converter generators + gqlgen config generator
	m.wg.Add(2*len(m.Libraries) + 1)

	for name, lib := range m.Libraries {

		_, libFile := path.Split(name)
		baseFileName := libFile[:len(libFile)-len(path.Ext(libFile))]

		// generate schemas in parallel
		go m.generateSchema(name, libFile, baseFileName, lib)

		// generate converters in parallel
		go m.generateConverter(name, libFile, baseFileName, lib)
	}

	// generate gqlgen model resolvers
	go m.generateGqlgenResolvers()

	m.wg.Wait()

	if m.failed {
		glog.Exit("error writing GraphQL schema files")
	}
}

func (m *Mapper) generateConverter(name string, libFile string, baseFileName string, lib *library.MetadataLibrary) {
	func(name string, lib *library.MetadataLibrary) {
		defer m.wg.Done()

		// get gql file name
		converterFile := baseFileName + ".converter.go"
		glog.Infof("writing %s from %s...", converterFile, libFile)

		outFile, err := os.Create(fmt.Sprintf("./internal/converter/%s", converterFile))
		if err != nil {
			glog.Errorf("error creating file %s: %v", converterFile, err)
			m.failed = true
			return
		}

		defer outFile.Close()
		writer := bufio.NewWriter(outFile)

		libGqlName, err := libraryGraphQLName(baseFileName)
		if err != nil {
			glog.Errorf("error creating converter type name from %s: %v", baseFileName, err)
			m.failed = true
			return
		}

		// execute template to generate converter
		err = m.ConverterTemplate.Execute(writer,
			&templateRequest{FileName: libFile, LibraryGraphQLName: libGqlName, Library: lib})
		if err != nil {
			glog.Errorf("error creating file %s: %v", converterFile, err)
			m.failed = true
			return
		}
		writer.Flush()

	}(name, lib)
}

func (m *Mapper) generateSchema(name string, libFile string, baseFileName string, lib *library.MetadataLibrary) {
	func(name string, lib *library.MetadataLibrary) {
		defer m.wg.Done()

		// get gql file name
		gqlFile := baseFileName + ".graphqls"
		glog.Infof("writing %s from %s...", gqlFile, libFile)

		outFile, err := os.Create(fmt.Sprintf("./api/graphql/%s", gqlFile))
		if err != nil {
			glog.Errorf("error creating file %s: %v", gqlFile, err)
			m.failed = true
			return
		}

		defer outFile.Close()
		writer := bufio.NewWriter(outFile)

		// execute template to generate gql schema
		err = m.SchemaTemplate.Execute(writer, &templateRequest{FileName: libFile, Library: lib})
		if err != nil {
			glog.Errorf("error creating file %s: %v", gqlFile, err)
			m.failed = true
			return
		}
		writer.Flush()

	}(name, lib)
}

func (m *Mapper) generateGqlgenResolvers() {
	defer m.wg.Done()

	// create sorted keys map of all library types
	modelMap, err := m.createModelMap(m.Libraries)
	if err != nil {
		glog.Errorf("error generating model resolvers: %w", err)
		m.failed = true
		return
	}
	glog.Infof("found %d library models", len(*modelMap))

	// read gqlgen.yml and load models
	data, err := os.ReadFile(gqlGenFileName)
	if err != nil {
		glog.Errorf("error reading %s: %w", gqlGenFileName, err)
		m.failed = true
		return
	}
	gqlgenYml := string(data)

	reader := strings.NewReader(gqlgenYml)
	builder := strings.Builder{}
	currentModels, err := ScanModelMapping(reader, &builder)
	glog.Infof("found %d current models", len(*currentModels))

	// replace the library models in existing models map
	for name, model := range *modelMap {
		(*currentModels)[name] = model
	}

	// write updated sorted models back to gqlgen.yml
	sortedMap := fmtsort.Sort(reflect.ValueOf(*currentModels))
	glog.Infof("found %d models after merging", len(sortedMap.Value))

	for _, s := range sortedMap.Value {
		builder.WriteString(s.String())
	}
	newGqlgenYml := builder.String()

	err = os.WriteFile(gqlGenFileName, []byte(newGqlgenYml), os.ModeExclusive)
	if err != nil {
		glog.Errorf("error writing %s: %w", gqlGenFileName, err)
		m.failed = true
		return
	}
	glog.Infof("successfully wrote %s", gqlGenFileName)
}

func (m *Mapper) createModelMap(libraries map[string]*library.MetadataLibrary) (*map[string]string, error) {
	result := make(map[string]string)
	for _, lib := range libraries {
		for _, t := range lib.ArtifactTypes {
			name, err := t.GraphQLName()
			if err != nil {
				return nil, err
			}
			model, err := m.generateModelResolvers(&t, "artifact-resolvers")
			if err != nil {
				return nil, err
			}
			result[name] = model
		}
		for _, t := range lib.ContextTypes {
			name, err := t.GraphQLName()
			if err != nil {
				return nil, err
			}
			model, err := m.generateModelResolvers(&t, "context-resolvers")
			if err != nil {
				return nil, err
			}
			result[name] = model
		}
		for _, t := range lib.ExecutionTypes {
			name, err := t.GraphQLName()
			if err != nil {
				return nil, err
			}
			model, err := m.generateModelResolvers(&t, "execution-resolvers")
			if err != nil {
				return nil, err
			}
			result[name] = model
		}
	}
	return &result, nil
}

func (m *Mapper) generateModelResolvers(t any, templateName string) (string, error) {
	builder := strings.Builder{}
	err := m.GqlgenTemplate.ExecuteTemplate(&builder, templateName, &t)
	if err != nil {
		return "", err
	}
	return builder.String(), nil
}

func libraryGraphQLName(baseFileName string) (string, error) {
	return library.ToGraphQLIdentifier(baseFileName, true)
}

type templateRequest struct {
	FileName           string
	LibraryGraphQLName string
	Library            *library.MetadataLibrary
}
