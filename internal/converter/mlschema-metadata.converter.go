// Code generated by github.com/opendatahub-io/model-registry/gqlmapper. DO NOT EDIT.
// converter generated from mlschema-metadata.yaml
package converter

import (
    "github.com/opendatahub-io/model-registry/internal/model/db"
    "github.com/opendatahub-io/model-registry/internal/model/graph"
)
type MlschemaMetadataConverter struct {
    DelegateConverter MlschemaMetadataDelegateConverter
}
func (c MlschemaMetadataConverter) ConvertMlschemaModel(source *graph.MlschemaModel) (*db.Artifact, error) {
    return c.DelegateConverter.ConvertMlschemaModel(source)
}
func (c MlschemaMetadataConverter) ConvertToMlschemaModel(source *db.Artifact) (*graph.MlschemaModel, error) {
    return c.DelegateConverter.ConvertToMlschemaModel(source)
}
func (c MlschemaMetadataConverter) ConvertMlschemaModelEvaluation(source *graph.MlschemaModelEvaluation) (*db.Artifact, error) {
    return c.DelegateConverter.ConvertMlschemaModelEvaluation(source)
}
func (c MlschemaMetadataConverter) ConvertToMlschemaModelEvaluation(source *db.Artifact) (*graph.MlschemaModelEvaluation, error) {
    return c.DelegateConverter.ConvertToMlschemaModelEvaluation(source)
}
func (c MlschemaMetadataConverter) ConvertMlschemaEvaluationMeasure(source *graph.MlschemaEvaluationMeasure) (*db.Artifact, error) {
    return c.DelegateConverter.ConvertMlschemaEvaluationMeasure(source)
}
func (c MlschemaMetadataConverter) ConvertToMlschemaEvaluationMeasure(source *db.Artifact) (*graph.MlschemaEvaluationMeasure, error) {
    return c.DelegateConverter.ConvertToMlschemaEvaluationMeasure(source)
}
func (c MlschemaMetadataConverter) ConvertMlschemaEvaluationSpecification(source *graph.MlschemaEvaluationSpecification) (*db.Artifact, error) {
    return c.DelegateConverter.ConvertMlschemaEvaluationSpecification(source)
}
func (c MlschemaMetadataConverter) ConvertToMlschemaEvaluationSpecification(source *db.Artifact) (*graph.MlschemaEvaluationSpecification, error) {
    return c.DelegateConverter.ConvertToMlschemaEvaluationSpecification(source)
}
func (c MlschemaMetadataConverter) ConvertMlschemaEvaluationProcedure(source *graph.MlschemaEvaluationProcedure) (*db.Artifact, error) {
    return c.DelegateConverter.ConvertMlschemaEvaluationProcedure(source)
}
func (c MlschemaMetadataConverter) ConvertToMlschemaEvaluationProcedure(source *db.Artifact) (*graph.MlschemaEvaluationProcedure, error) {
    return c.DelegateConverter.ConvertToMlschemaEvaluationProcedure(source)
}
func (c MlschemaMetadataConverter) ConvertMlschemaTask(source *graph.MlschemaTask) (*db.Artifact, error) {
    return c.DelegateConverter.ConvertMlschemaTask(source)
}
func (c MlschemaMetadataConverter) ConvertToMlschemaTask(source *db.Artifact) (*graph.MlschemaTask, error) {
    return c.DelegateConverter.ConvertToMlschemaTask(source)
}
func (c MlschemaMetadataConverter) ConvertMlschemaData(source *graph.MlschemaData) (*db.Artifact, error) {
    return c.DelegateConverter.ConvertMlschemaData(source)
}
func (c MlschemaMetadataConverter) ConvertToMlschemaData(source *db.Artifact) (*graph.MlschemaData, error) {
    return c.DelegateConverter.ConvertToMlschemaData(source)
}
func (c MlschemaMetadataConverter) ConvertMlschemaDataSet(source *graph.MlschemaDataSet) (*db.Artifact, error) {
    return c.DelegateConverter.ConvertMlschemaDataSet(source)
}
func (c MlschemaMetadataConverter) ConvertToMlschemaDataSet(source *db.Artifact) (*graph.MlschemaDataSet, error) {
    return c.DelegateConverter.ConvertToMlschemaDataSet(source)
}
func (c MlschemaMetadataConverter) ConvertMlschemaFeature(source *graph.MlschemaFeature) (*db.Artifact, error) {
    return c.DelegateConverter.ConvertMlschemaFeature(source)
}
func (c MlschemaMetadataConverter) ConvertToMlschemaFeature(source *db.Artifact) (*graph.MlschemaFeature, error) {
    return c.DelegateConverter.ConvertToMlschemaFeature(source)
}
func (c MlschemaMetadataConverter) ConvertMlschemaAlgorithm(source *graph.MlschemaAlgorithm) (*db.Artifact, error) {
    return c.DelegateConverter.ConvertMlschemaAlgorithm(source)
}
func (c MlschemaMetadataConverter) ConvertToMlschemaAlgorithm(source *db.Artifact) (*graph.MlschemaAlgorithm, error) {
    return c.DelegateConverter.ConvertToMlschemaAlgorithm(source)
}
func (c MlschemaMetadataConverter) ConvertMlschemaSoftware(source *graph.MlschemaSoftware) (*db.Artifact, error) {
    return c.DelegateConverter.ConvertMlschemaSoftware(source)
}
func (c MlschemaMetadataConverter) ConvertToMlschemaSoftware(source *db.Artifact) (*graph.MlschemaSoftware, error) {
    return c.DelegateConverter.ConvertToMlschemaSoftware(source)
}
func (c MlschemaMetadataConverter) ConvertMlschemaImplementation(source *graph.MlschemaImplementation) (*db.Artifact, error) {
    return c.DelegateConverter.ConvertMlschemaImplementation(source)
}
func (c MlschemaMetadataConverter) ConvertToMlschemaImplementation(source *db.Artifact) (*graph.MlschemaImplementation, error) {
    return c.DelegateConverter.ConvertToMlschemaImplementation(source)
}
func (c MlschemaMetadataConverter) ConvertMlschemaHyperParameter(source *graph.MlschemaHyperParameter) (*db.Artifact, error) {
    return c.DelegateConverter.ConvertMlschemaHyperParameter(source)
}
func (c MlschemaMetadataConverter) ConvertToMlschemaHyperParameter(source *db.Artifact) (*graph.MlschemaHyperParameter, error) {
    return c.DelegateConverter.ConvertToMlschemaHyperParameter(source)
}
func (c MlschemaMetadataConverter) ConvertMlschemaHyperParameterSetting(source *graph.MlschemaHyperParameterSetting) (*db.Artifact, error) {
    return c.DelegateConverter.ConvertMlschemaHyperParameterSetting(source)
}
func (c MlschemaMetadataConverter) ConvertToMlschemaHyperParameterSetting(source *db.Artifact) (*graph.MlschemaHyperParameterSetting, error) {
    return c.DelegateConverter.ConvertToMlschemaHyperParameterSetting(source)
}
func (c MlschemaMetadataConverter) ConvertMlschemaImplementationCharacteristic(source *graph.MlschemaImplementationCharacteristic) (*db.Artifact, error) {
    return c.DelegateConverter.ConvertMlschemaImplementationCharacteristic(source)
}
func (c MlschemaMetadataConverter) ConvertToMlschemaImplementationCharacteristic(source *db.Artifact) (*graph.MlschemaImplementationCharacteristic, error) {
    return c.DelegateConverter.ConvertToMlschemaImplementationCharacteristic(source)
}
func (c MlschemaMetadataConverter) ConvertMlschemaDatasetCharacteristic(source *graph.MlschemaDatasetCharacteristic) (*db.Artifact, error) {
    return c.DelegateConverter.ConvertMlschemaDatasetCharacteristic(source)
}
func (c MlschemaMetadataConverter) ConvertToMlschemaDatasetCharacteristic(source *db.Artifact) (*graph.MlschemaDatasetCharacteristic, error) {
    return c.DelegateConverter.ConvertToMlschemaDatasetCharacteristic(source)
}
func (c MlschemaMetadataConverter) ConvertMlschemaFeatureCharacteristic(source *graph.MlschemaFeatureCharacteristic) (*db.Artifact, error) {
    return c.DelegateConverter.ConvertMlschemaFeatureCharacteristic(source)
}
func (c MlschemaMetadataConverter) ConvertToMlschemaFeatureCharacteristic(source *db.Artifact) (*graph.MlschemaFeatureCharacteristic, error) {
    return c.DelegateConverter.ConvertToMlschemaFeatureCharacteristic(source)
}
func (c MlschemaMetadataConverter) ConvertMlschemaRegisteredModel(source *graph.MlschemaRegisteredModel) (*db.Context, error) {
    return c.DelegateConverter.ConvertMlschemaRegisteredModel(source)
}
func (c MlschemaMetadataConverter) ConvertToMlschemaRegisteredModel(source *db.Context) (*graph.MlschemaRegisteredModel, error) {
    return c.DelegateConverter.ConvertToMlschemaRegisteredModel(source)
}
func (c MlschemaMetadataConverter) ConvertMlschemaModelVersion(source *graph.MlschemaModelVersion) (*db.Context, error) {
    return c.DelegateConverter.ConvertMlschemaModelVersion(source)
}
func (c MlschemaMetadataConverter) ConvertToMlschemaModelVersion(source *db.Context) (*graph.MlschemaModelVersion, error) {
    return c.DelegateConverter.ConvertToMlschemaModelVersion(source)
}
func (c MlschemaMetadataConverter) ConvertMlschemaDataCharacteristic(source *graph.MlschemaDataCharacteristic) (*db.Context, error) {
    return c.DelegateConverter.ConvertMlschemaDataCharacteristic(source)
}
func (c MlschemaMetadataConverter) ConvertToMlschemaDataCharacteristic(source *db.Context) (*graph.MlschemaDataCharacteristic, error) {
    return c.DelegateConverter.ConvertToMlschemaDataCharacteristic(source)
}
func (c MlschemaMetadataConverter) ConvertMlschemaStudy(source *graph.MlschemaStudy) (*db.Execution, error) {
    return c.DelegateConverter.ConvertMlschemaStudy(source)
}
func (c MlschemaMetadataConverter) ConvertToMlschemaStudy(source *db.Execution) (*graph.MlschemaStudy, error) {
    return c.DelegateConverter.ConvertToMlschemaStudy(source)
}
func (c MlschemaMetadataConverter) ConvertMlschemaExperiment(source *graph.MlschemaExperiment) (*db.Execution, error) {
    return c.DelegateConverter.ConvertMlschemaExperiment(source)
}
func (c MlschemaMetadataConverter) ConvertToMlschemaExperiment(source *db.Execution) (*graph.MlschemaExperiment, error) {
    return c.DelegateConverter.ConvertToMlschemaExperiment(source)
}
func (c MlschemaMetadataConverter) ConvertMlschemaRun(source *graph.MlschemaRun) (*db.Execution, error) {
    return c.DelegateConverter.ConvertMlschemaRun(source)
}
func (c MlschemaMetadataConverter) ConvertToMlschemaRun(source *db.Execution) (*graph.MlschemaRun, error) {
    return c.DelegateConverter.ConvertToMlschemaRun(source)
}

// goverter:converter
// goverter:wrapErrors
// goverter:extend ConvertArtifactState
// goverter:extend ConvertToArtifactState
// goverter:extend ConvertExecutionState
// goverter:extend ConvertToExecutionState
type MlschemaMetadataDelegateConverter interface {
    // goverter:map ID | ConvertArtifactID
    // goverter:map TypeID | ConvertTypeID
    // goverter:map Properties | ConvertArtifactProperties
    // goverter:ignore ArtifactType Attributions Events
    ConvertMlschemaModel(source *graph.MlschemaModel) (*db.Artifact, error)
    // goverter:map ID | ConvertToArtifactID
    // goverter:map TypeID | ConvertToTypeID
    // goverter:map Properties | ConvertToArtifactProperties
    // goverter:ignore Type Attributions Events
    ConvertToMlschemaModel(source *db.Artifact) (*graph.MlschemaModel, error)
    // goverter:map ID | ConvertArtifactID
    // goverter:map TypeID | ConvertTypeID
    // goverter:map Properties | ConvertArtifactProperties
    // goverter:ignore ArtifactType Attributions Events
    ConvertMlschemaModelEvaluation(source *graph.MlschemaModelEvaluation) (*db.Artifact, error)
    // goverter:map ID | ConvertToArtifactID
    // goverter:map TypeID | ConvertToTypeID
    // goverter:map Properties | ConvertToArtifactProperties
    // goverter:ignore Type Attributions Events
    ConvertToMlschemaModelEvaluation(source *db.Artifact) (*graph.MlschemaModelEvaluation, error)
    // goverter:map ID | ConvertArtifactID
    // goverter:map TypeID | ConvertTypeID
    // goverter:map Properties | ConvertArtifactProperties
    // goverter:ignore ArtifactType Attributions Events
    ConvertMlschemaEvaluationMeasure(source *graph.MlschemaEvaluationMeasure) (*db.Artifact, error)
    // goverter:map ID | ConvertToArtifactID
    // goverter:map TypeID | ConvertToTypeID
    // goverter:map Properties | ConvertToArtifactProperties
    // goverter:ignore Type Attributions Events
    ConvertToMlschemaEvaluationMeasure(source *db.Artifact) (*graph.MlschemaEvaluationMeasure, error)
    // goverter:map ID | ConvertArtifactID
    // goverter:map TypeID | ConvertTypeID
    // goverter:map Properties | ConvertArtifactProperties
    // goverter:ignore ArtifactType Attributions Events
    ConvertMlschemaEvaluationSpecification(source *graph.MlschemaEvaluationSpecification) (*db.Artifact, error)
    // goverter:map ID | ConvertToArtifactID
    // goverter:map TypeID | ConvertToTypeID
    // goverter:map Properties | ConvertToArtifactProperties
    // goverter:ignore Type Attributions Events
    ConvertToMlschemaEvaluationSpecification(source *db.Artifact) (*graph.MlschemaEvaluationSpecification, error)
    // goverter:map ID | ConvertArtifactID
    // goverter:map TypeID | ConvertTypeID
    // goverter:map Properties | ConvertArtifactProperties
    // goverter:ignore ArtifactType Attributions Events
    ConvertMlschemaEvaluationProcedure(source *graph.MlschemaEvaluationProcedure) (*db.Artifact, error)
    // goverter:map ID | ConvertToArtifactID
    // goverter:map TypeID | ConvertToTypeID
    // goverter:map Properties | ConvertToArtifactProperties
    // goverter:ignore Type Attributions Events
    ConvertToMlschemaEvaluationProcedure(source *db.Artifact) (*graph.MlschemaEvaluationProcedure, error)
    // goverter:map ID | ConvertArtifactID
    // goverter:map TypeID | ConvertTypeID
    // goverter:map Properties | ConvertArtifactProperties
    // goverter:ignore ArtifactType Attributions Events
    ConvertMlschemaTask(source *graph.MlschemaTask) (*db.Artifact, error)
    // goverter:map ID | ConvertToArtifactID
    // goverter:map TypeID | ConvertToTypeID
    // goverter:map Properties | ConvertToArtifactProperties
    // goverter:ignore Type Attributions Events
    ConvertToMlschemaTask(source *db.Artifact) (*graph.MlschemaTask, error)
    // goverter:map ID | ConvertArtifactID
    // goverter:map TypeID | ConvertTypeID
    // goverter:map Properties | ConvertArtifactProperties
    // goverter:ignore ArtifactType Attributions Events
    ConvertMlschemaData(source *graph.MlschemaData) (*db.Artifact, error)
    // goverter:map ID | ConvertToArtifactID
    // goverter:map TypeID | ConvertToTypeID
    // goverter:map Properties | ConvertToArtifactProperties
    // goverter:ignore Type Attributions Events
    ConvertToMlschemaData(source *db.Artifact) (*graph.MlschemaData, error)
    // goverter:map ID | ConvertArtifactID
    // goverter:map TypeID | ConvertTypeID
    // goverter:map Properties | ConvertArtifactProperties
    // goverter:ignore ArtifactType Attributions Events
    ConvertMlschemaDataSet(source *graph.MlschemaDataSet) (*db.Artifact, error)
    // goverter:map ID | ConvertToArtifactID
    // goverter:map TypeID | ConvertToTypeID
    // goverter:map Properties | ConvertToArtifactProperties
    // goverter:ignore Type Attributions Events
    ConvertToMlschemaDataSet(source *db.Artifact) (*graph.MlschemaDataSet, error)
    // goverter:map ID | ConvertArtifactID
    // goverter:map TypeID | ConvertTypeID
    // goverter:map Properties | ConvertArtifactProperties
    // goverter:ignore ArtifactType Attributions Events
    ConvertMlschemaFeature(source *graph.MlschemaFeature) (*db.Artifact, error)
    // goverter:map ID | ConvertToArtifactID
    // goverter:map TypeID | ConvertToTypeID
    // goverter:map Properties | ConvertToArtifactProperties
    // goverter:ignore Type Attributions Events
    ConvertToMlschemaFeature(source *db.Artifact) (*graph.MlschemaFeature, error)
    // goverter:map ID | ConvertArtifactID
    // goverter:map TypeID | ConvertTypeID
    // goverter:map Properties | ConvertArtifactProperties
    // goverter:ignore ArtifactType Attributions Events
    ConvertMlschemaAlgorithm(source *graph.MlschemaAlgorithm) (*db.Artifact, error)
    // goverter:map ID | ConvertToArtifactID
    // goverter:map TypeID | ConvertToTypeID
    // goverter:map Properties | ConvertToArtifactProperties
    // goverter:ignore Type Attributions Events
    ConvertToMlschemaAlgorithm(source *db.Artifact) (*graph.MlschemaAlgorithm, error)
    // goverter:map ID | ConvertArtifactID
    // goverter:map TypeID | ConvertTypeID
    // goverter:map Properties | ConvertArtifactProperties
    // goverter:ignore ArtifactType Attributions Events
    ConvertMlschemaSoftware(source *graph.MlschemaSoftware) (*db.Artifact, error)
    // goverter:map ID | ConvertToArtifactID
    // goverter:map TypeID | ConvertToTypeID
    // goverter:map Properties | ConvertToArtifactProperties
    // goverter:ignore Type Attributions Events
    ConvertToMlschemaSoftware(source *db.Artifact) (*graph.MlschemaSoftware, error)
    // goverter:map ID | ConvertArtifactID
    // goverter:map TypeID | ConvertTypeID
    // goverter:map Properties | ConvertArtifactProperties
    // goverter:ignore ArtifactType Attributions Events
    ConvertMlschemaImplementation(source *graph.MlschemaImplementation) (*db.Artifact, error)
    // goverter:map ID | ConvertToArtifactID
    // goverter:map TypeID | ConvertToTypeID
    // goverter:map Properties | ConvertToArtifactProperties
    // goverter:ignore Type Attributions Events
    ConvertToMlschemaImplementation(source *db.Artifact) (*graph.MlschemaImplementation, error)
    // goverter:map ID | ConvertArtifactID
    // goverter:map TypeID | ConvertTypeID
    // goverter:map Properties | ConvertArtifactProperties
    // goverter:ignore ArtifactType Attributions Events
    ConvertMlschemaHyperParameter(source *graph.MlschemaHyperParameter) (*db.Artifact, error)
    // goverter:map ID | ConvertToArtifactID
    // goverter:map TypeID | ConvertToTypeID
    // goverter:map Properties | ConvertToArtifactProperties
    // goverter:ignore Type Attributions Events
    ConvertToMlschemaHyperParameter(source *db.Artifact) (*graph.MlschemaHyperParameter, error)
    // goverter:map ID | ConvertArtifactID
    // goverter:map TypeID | ConvertTypeID
    // goverter:map Properties | ConvertArtifactProperties
    // goverter:ignore ArtifactType Attributions Events
    ConvertMlschemaHyperParameterSetting(source *graph.MlschemaHyperParameterSetting) (*db.Artifact, error)
    // goverter:map ID | ConvertToArtifactID
    // goverter:map TypeID | ConvertToTypeID
    // goverter:map Properties | ConvertToArtifactProperties
    // goverter:ignore Type Attributions Events
    ConvertToMlschemaHyperParameterSetting(source *db.Artifact) (*graph.MlschemaHyperParameterSetting, error)
    // goverter:map ID | ConvertArtifactID
    // goverter:map TypeID | ConvertTypeID
    // goverter:map Properties | ConvertArtifactProperties
    // goverter:ignore ArtifactType Attributions Events
    ConvertMlschemaImplementationCharacteristic(source *graph.MlschemaImplementationCharacteristic) (*db.Artifact, error)
    // goverter:map ID | ConvertToArtifactID
    // goverter:map TypeID | ConvertToTypeID
    // goverter:map Properties | ConvertToArtifactProperties
    // goverter:ignore Type Attributions Events
    ConvertToMlschemaImplementationCharacteristic(source *db.Artifact) (*graph.MlschemaImplementationCharacteristic, error)
    // goverter:map ID | ConvertArtifactID
    // goverter:map TypeID | ConvertTypeID
    // goverter:map Properties | ConvertArtifactProperties
    // goverter:ignore ArtifactType Attributions Events
    ConvertMlschemaDatasetCharacteristic(source *graph.MlschemaDatasetCharacteristic) (*db.Artifact, error)
    // goverter:map ID | ConvertToArtifactID
    // goverter:map TypeID | ConvertToTypeID
    // goverter:map Properties | ConvertToArtifactProperties
    // goverter:ignore Type Attributions Events
    ConvertToMlschemaDatasetCharacteristic(source *db.Artifact) (*graph.MlschemaDatasetCharacteristic, error)
    // goverter:map ID | ConvertArtifactID
    // goverter:map TypeID | ConvertTypeID
    // goverter:map Properties | ConvertArtifactProperties
    // goverter:ignore ArtifactType Attributions Events
    ConvertMlschemaFeatureCharacteristic(source *graph.MlschemaFeatureCharacteristic) (*db.Artifact, error)
    // goverter:map ID | ConvertToArtifactID
    // goverter:map TypeID | ConvertToTypeID
    // goverter:map Properties | ConvertToArtifactProperties
    // goverter:ignore Type Attributions Events
    ConvertToMlschemaFeatureCharacteristic(source *db.Artifact) (*graph.MlschemaFeatureCharacteristic, error)
    // goverter:map ID | ConvertContextID
    // goverter:map TypeID | ConvertTypeID
    // goverter:map Properties | ConvertContextProperties
    // goverter:ignore ContextType Attributions Associations Parents Children
    ConvertMlschemaRegisteredModel(source *graph.MlschemaRegisteredModel) (*db.Context, error)
    // goverter:map ID | ConvertToContextID
    // goverter:map TypeID | ConvertToTypeID
    // goverter:map Properties | ConvertToContextProperties
    // goverter:ignore Type Attributions Associations Parents Children
    ConvertToMlschemaRegisteredModel(source *db.Context) (*graph.MlschemaRegisteredModel, error)
    // goverter:map ID | ConvertContextID
    // goverter:map TypeID | ConvertTypeID
    // goverter:map Properties | ConvertContextProperties
    // goverter:ignore ContextType Attributions Associations Parents Children
    ConvertMlschemaModelVersion(source *graph.MlschemaModelVersion) (*db.Context, error)
    // goverter:map ID | ConvertToContextID
    // goverter:map TypeID | ConvertToTypeID
    // goverter:map Properties | ConvertToContextProperties
    // goverter:ignore Type Attributions Associations Parents Children
    ConvertToMlschemaModelVersion(source *db.Context) (*graph.MlschemaModelVersion, error)
    // goverter:map ID | ConvertContextID
    // goverter:map TypeID | ConvertTypeID
    // goverter:map Properties | ConvertContextProperties
    // goverter:ignore ContextType Attributions Associations Parents Children
    ConvertMlschemaDataCharacteristic(source *graph.MlschemaDataCharacteristic) (*db.Context, error)
    // goverter:map ID | ConvertToContextID
    // goverter:map TypeID | ConvertToTypeID
    // goverter:map Properties | ConvertToContextProperties
    // goverter:ignore Type Attributions Associations Parents Children
    ConvertToMlschemaDataCharacteristic(source *db.Context) (*graph.MlschemaDataCharacteristic, error)
    // goverter:map ID | ConvertExecutionID
    // goverter:map TypeID | ConvertTypeID
    // goverter:map Properties | ConvertExecutionProperties
    // goverter:ignore ExecutionType Associations Events
    ConvertMlschemaStudy(source *graph.MlschemaStudy) (*db.Execution, error)
    // goverter:map ID | ConvertToExecutionID
    // goverter:map TypeID | ConvertToTypeID
    // goverter:map Properties | ConvertToExecutionProperties
    // goverter:ignore Type Associations Events
    ConvertToMlschemaStudy(source *db.Execution) (*graph.MlschemaStudy, error)
    // goverter:map ID | ConvertExecutionID
    // goverter:map TypeID | ConvertTypeID
    // goverter:map Properties | ConvertExecutionProperties
    // goverter:ignore ExecutionType Associations Events
    ConvertMlschemaExperiment(source *graph.MlschemaExperiment) (*db.Execution, error)
    // goverter:map ID | ConvertToExecutionID
    // goverter:map TypeID | ConvertToTypeID
    // goverter:map Properties | ConvertToExecutionProperties
    // goverter:ignore Type Associations Events
    ConvertToMlschemaExperiment(source *db.Execution) (*graph.MlschemaExperiment, error)
    // goverter:map ID | ConvertExecutionID
    // goverter:map TypeID | ConvertTypeID
    // goverter:map Properties | ConvertExecutionProperties
    // goverter:ignore ExecutionType Associations Events
    ConvertMlschemaRun(source *graph.MlschemaRun) (*db.Execution, error)
    // goverter:map ID | ConvertToExecutionID
    // goverter:map TypeID | ConvertToTypeID
    // goverter:map Properties | ConvertToExecutionProperties
    // goverter:ignore Type Associations Events
    ConvertToMlschemaRun(source *db.Execution) (*graph.MlschemaRun, error)
}
