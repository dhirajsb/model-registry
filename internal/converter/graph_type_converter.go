package converter

import (
	"github.com/opendatahub-io/model-registry/internal/ml_metadata/proto"
	"github.com/opendatahub-io/model-registry/internal/model/db"
	"github.com/opendatahub-io/model-registry/internal/model/graph"
)

// goverter:converter
// goverter:wrapErrors
type TypeConverter interface {
	// goverter:map ID | ConvertTypeID
	// goverter:map TypeKind | ConvertTypeKind
	// goverter:map Properties | ConvertTypeProperties
	// goverter:ignore InputType OutputType
	ConvertArtifactType(source *graph.ArtifactType) (*db.Type, error)
	// goverter:map ID | ConvertToTypeID
	// goverter:map TypeKind | ConvertToTypeKind
	// goverter:map Properties | ConvertToTypeProperties
	ConvertToArtifactType(source *db.Type) (*graph.ArtifactType, error)

	// goverter:map ID | ConvertTypeID
	// goverter:map TypeKind | ConvertTypeKind
	// goverter:map Properties | ConvertTypeProperties
	// goverter:ignore InputType OutputType
	ConvertContextType(source *graph.ContextType) (*db.Type, error)
	// goverter:map ID | ConvertToTypeID
	// goverter:map TypeKind | ConvertToTypeKind
	// goverter:map Properties | ConvertToTypeProperties
	ConvertToContextType(source *db.Type) (*graph.ContextType, error)

	// goverter:map ID | ConvertTypeID
	// goverter:map TypeKind | ConvertTypeKind
	// goverter:map Properties | ConvertTypeProperties
	// goverter:ignore InputType OutputType
	ConvertExecutionType(source *graph.ExecutionType) (*db.Type, error)
	// goverter:map ID | ConvertToTypeID
	// goverter:map TypeKind | ConvertToTypeKind
	// goverter:map Properties | ConvertToTypeProperties
	ConvertToExecutionType(source *db.Type) (*graph.ExecutionType, error)
}

func ConvertTypeProperties(props []*graph.TypeProperty) ([]db.TypeProperty, error) {
	var result []db.TypeProperty
	for _, prop := range props {
		result = append(result, db.TypeProperty{
			Name:     prop.Name,
			DataType: proto.PropertyType_value[prop.DataType.String()],
		})
	}
	return result, nil
}

func ConvertToTypeProperties(props []db.TypeProperty) ([]*graph.TypeProperty, error) {
	var result []*graph.TypeProperty
	for _, prop := range props {
		result = append(result, &graph.TypeProperty{
			Name:     prop.Name,
			DataType: graph.DataType(proto.PropertyType_name[prop.DataType]),
		})
	}
	return result, nil
}

var typeKind_values = map[graph.TypeKind]int8{graph.TypeKindExecutionType: 0, graph.TypeKindArtifactType: 1, graph.TypeKindContextType: 2}
var typeKind_names = map[int8]graph.TypeKind{0: graph.TypeKindExecutionType, 1: graph.TypeKindArtifactType, 2: graph.TypeKindContextType}

func ConvertTypeKind(source graph.TypeKind) int8 {
	return typeKind_values[source]
}

func ConvertToTypeKind(source int8) graph.TypeKind {
	return typeKind_names[source]
}
