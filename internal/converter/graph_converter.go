package converter

import (
	"github.com/opendatahub-io/model-registry/internal/model/db"
	"github.com/opendatahub-io/model-registry/internal/model/graph"
)

type ConvertArtifact func(source *graph.ArtifactInterface) (*db.Artifact, error)
type ConvertToArtifact func(source *db.Artifact) (*graph.ArtifactInterface, error)
type ConvertContext func(source *graph.ContextInterface) (*db.Context, error)
type ConvertToContext func(source *db.Context) (*graph.ContextInterface, error)
type ConvertExecution func(source *graph.ExecutionInterface) (*db.Execution, error)
type ConvertToExecution func(source *db.Execution) (*graph.ExecutionInterface, error)

// goverter:converter
// goverter:wrapErrors
// goverter:extend ConvertArtifactState
// goverter:extend ConvertToArtifactState
// goverter:extend ConvertExecutionState
// goverter:extend ConvertToExecutionState
type GraphConverter interface {
	// goverter:map ID | ConvertArtifactID
	// goverter:map TypeID | ConvertTypeID
	// goverter:map Properties | ConvertArtifactProperties
	// goverter:ignore ArtifactType Attributions Events
	ConvertArtifact(source *graph.Artifact) (*db.Artifact, error)
	// goverter:map ID | ConvertToArtifactID
	// goverter:map TypeID | ConvertToTypeID
	// goverter:map Properties | ConvertToArtifactProperties
	// goverter:ignore Type Attributions Events
	ConvertToArtifact(source *db.Artifact) (*graph.Artifact, error)

	// goverter:map ID | ConvertContextID
	// goverter:map TypeID | ConvertTypeID
	// goverter:map Properties | ConvertContextProperties
	// goverter:ignore ContextType Attributions Associations Parents Children
	ConvertContext(source *graph.Context) (*db.Context, error)
	// goverter:map ID | ConvertToContextID
	// goverter:map TypeID | ConvertToTypeID
	// goverter:map Properties | ConvertToContextProperties
	// goverter:ignore Type Attributions Associations Parents Children
	ConvertToContext(source *db.Context) (*graph.Context, error)

	// goverter:map ID | ConvertExecutionID
	// goverter:map TypeID | ConvertTypeID
	// goverter:map Properties | ConvertExecutionProperties
	// goverter:ignore ExecutionType Associations Events
	ConvertExecution(source *graph.Execution) (*db.Execution, error)
	// goverter:map ID | ConvertToExecutionID
	// goverter:map TypeID | ConvertToTypeID
	// goverter:map Properties | ConvertToExecutionProperties
	// goverter:ignore Type Associations Events
	ConvertToExecution(source *db.Execution) (*graph.Execution, error)

	// goverter:map ID | ConvertEventID
	// goverter:map Type | ConvertEventType
	// goverter:map ArtifactID | ConvertArtifactID
	// goverter:map ExecutionID | ConvertExecutionID
	// goverter:map Path PathSteps | ConvertEventPath
	// goverter:ignore Artifact Execution
	ConvertEvent(source *graph.Event) (*db.Event, error)
	// goverter:map ID | ConvertToEventID
	// goverter:map Type | ConvertToEventType
	// goverter:map ArtifactID | ConvertToArtifactID
	// goverter:map ExecutionID | ConvertToExecutionID
	// goverter:map PathSteps Path | ConvertToEventPath
	// goverter:ignore Artifact Execution
	ConvertToEvent(source *db.Event) (*graph.Event, error)

	// goverter:map ID | ConvertAttributionID
	// goverter:map ContextID | ConvertContextID
	// goverter:map ArtifactID | ConvertArtifactID
	// goverter:ignore Context Artifact
	ConvertAttribution(source *graph.Attribution) (*db.Attribution, error)
	// goverter:map ID | ConvertToAttributionID
	// goverter:map ContextID | ConvertToContextID
	// goverter:map ArtifactID | ConvertToArtifactID
	// goverter:ignore Context Artifact
	ConvertToAttribution(source *db.Attribution) (*graph.Attribution, error)

	// goverter:map ID | ConvertAssociationID
	// goverter:map ContextID | ConvertContextID
	// goverter:map ExecutionID | ConvertExecutionID
	// goverter:ignore Context Execution
	ConvertAssociation(source *graph.Association) (*db.Association, error)
	// goverter:map ID | ConvertToAssociationID
	// goverter:map ContextID | ConvertToContextID
	// goverter:map ExecutionID | ConvertToExecutionID
	// goverter:ignore Context Execution
	ConvertToAssociation(source *db.Association) (*graph.Association, error)
}
