package library

import (
	"fmt"
	"github.com/gertd/go-pluralize"
	"github.com/golang/glog"
	"gopkg.in/yaml.v3"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"unicode"
	"unicode/utf8"
)

//go:generate go-enum -type=PropertyValueType

type PropertyValueType int32

const (
	UNKNOWN PropertyValueType = iota
	INT
	DOUBLE
	STRING
	STRUCT
	PROTO
	BOOLEAN
)

type MetadataType struct {
	Name             *string        `yaml:"name,omitempty"`
	Version          *string        `yaml:"version,omitempty"`
	Description      *string        `yaml:"description,omitempty"`
	ExternalId       *string        `yaml:"external_id,omitempty"`
	Properties       []PropertyType `yaml:"properties,omitempty"`
	graphQLName      *string        `yaml:"-"`
	graphQLQueryName *string        `yaml:"-"`
}

type PropertyType struct {
	Name           string            `yaml:"name"`
	Type           PropertyValueType `yaml:"type"`
	Description    *string           `yaml:"description,omitempty"`
	graphQLName    *string           `yaml:"-"`
	goName         *string           `yaml:"-"`
	graphQLType    *string           `yaml:"-"`
	graphQLWrapper *string           `yaml:"-"`
}

type ArtifactType struct {
	MetadataType `yaml:",inline"`
	// TODO add support for base type enum
	//BaseType *ArtifactType_SystemDefinedBaseType `yaml:"base_type,omitempty"`
}

type ContextType struct {
	MetadataType `yaml:",inline"`
}

type ExecutionType struct {
	MetadataType `yaml:",inline"`
	//InputType  *ArtifactStructType                  `yaml:"input_type,omitempty"`
	//OutputType *ArtifactStructType                  `yaml:"output_type,omitempty"`
	//BaseType   *ExecutionType_SystemDefinedBaseType `yaml:"base_type,omitempty"`
}

type MetadataLibrary struct {
	ArtifactTypes  []ArtifactType  `yaml:"artifact-types,omitempty"`
	ContextTypes   []ContextType   `yaml:"context-types,omitempty"`
	ExecutionTypes []ExecutionType `yaml:"execution-types,omitempty"`
}

func LoadLibraries(dirs []string) (map[string]*MetadataLibrary, error) {
	result := make(map[string]*MetadataLibrary)
	for _, dir := range dirs {
		abs, err := filepath.Abs(dir)
		if err != nil {
			return nil, fmt.Errorf("error getting absolute library path for %s: %w", dir, err)
		}
		_, err = os.Stat(abs)
		if err != nil {
			return nil, fmt.Errorf("error opening library path for %s: %w", abs, err)
		}
		err = filepath.WalkDir(abs, func(path string, entry fs.DirEntry, err error) error {
			if err != nil {
				glog.Warningf("error reading library path %s: %v", path, err)
				return filepath.SkipDir
			}
			if entry.IsDir() || !isYamlFile(path) {
				return nil
			}

			bytes, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("failed to read library file %s: %w", path, err)
			}
			lib := &MetadataLibrary{}
			err = yaml.Unmarshal(bytes, lib)
			if err != nil {
				return fmt.Errorf("failed to parse library file %s: %w", path, err)
			}
			result[path] = lib
			return nil
		})
		if err != nil {
			return nil, fmt.Errorf("failed to read library directory %s: %w", abs, err)
		}
	}
	return result, nil
}

func isYamlFile(path string) bool {
	lowerPath := strings.ToLower(filepath.Ext(path))
	return strings.HasSuffix(lowerPath, ".yaml") || strings.HasSuffix(lowerPath, ".yml")
}
func (t *MetadataType) GraphQLName() (string, error) {
	if t.graphQLName == nil {
		name, err := ToGraphQLIdentifier(*t.Name, true)
		if err != nil {
			return name, err
		}
		t.graphQLName = &name
	}
	return *t.graphQLName, nil
}

func ToGraphQLIdentifier(name string, capitalFirst bool) (string, error) {
	r, size := utf8.DecodeRuneInString(name)
	if r == utf8.RuneError {
		return "", fmt.Errorf("unable to convert %s to GraphQL identifier", name)
	}
	result := strings.Builder{}
	if capitalFirst {
		result.WriteRune(unicode.ToUpper(r))
	} else {
		result.WriteRune(unicode.ToLower(r))
	}
	wordBreak := false
	// copy the rest, skipping any non alphanumerics and '_'
	_ = strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || r == '_' || unicode.IsDigit(r) {
			if wordBreak {
				r = unicode.ToUpper(r)
				wordBreak = false
			}
			result.WriteRune(r)
		} else {
			// word break at non alphanumeric and '_'
			wordBreak = true
		}
		return r
	}, (name)[size:])
	return result.String(), nil
}

var client = pluralize.NewClient()

func (t *MetadataType) GraphQLQueryName() (string, error) {
	if t.graphQLQueryName == nil {
		name, err := t.GraphQLName()
		if err != nil {
			return name, err
		}
		r, size := utf8.DecodeRuneInString(name)
		name = string(unicode.ToLower(r)) + name[size:]
		str := client.Plural(name)
		t.graphQLQueryName = &str
	}
	return *t.graphQLQueryName, nil
}

func (p *PropertyType) GraphQLName() (name string, err error) {
	if p.graphQLName == nil {
		name, err = ToGraphQLIdentifier(p.Name, false)
		if err != nil {
			return name, err
		}
		p.graphQLName = &name
	}
	return *p.graphQLName, nil
}

func (p *PropertyType) GoName() (name string, err error) {
	if p.goName == nil {
		name, err = ToGraphQLIdentifier(p.Name, true)
		if err != nil {
			return name, err
		}
		p.goName = &name
	}
	return *p.goName, nil
}

// FieldGraphQLType maps fields to GraphQL scalar types when possible
func (p *PropertyType) GraphQLType() (string, error) {
	if p.graphQLType == nil {
		result := "String"
		switch p.Type {
		case INT:
			result = "Int64"
		case DOUBLE:
			result = "Float"
		case STRING:
			result = "String"
		case STRUCT:
			result = "[StructTuple!]"
		case PROTO:
			result = "ProtoTypeValue"
		case BOOLEAN:
			result = "Boolean"
		}
		p.graphQLType = &result
	}
	return *p.graphQLType, nil
}

// FieldGraphQLWrapper maps fields to GraphQL wrapper types
func (p *PropertyType) GraphQLWrapper() (string, error) {
	if p.graphQLWrapper == nil {
		result := "StringValue"
		switch p.Type {
		case INT:
			result = "IntValue"
		case DOUBLE:
			result = "DoubleValue"
		case STRING:
			result = "StringValue"
		case STRUCT:
			result = "StructValue"
		case PROTO:
			result = "ProtoValue"
		case BOOLEAN:
			result = "BoolValue"
		}
		p.graphQLWrapper = &result
	}
	return *p.graphQLWrapper, nil
}
