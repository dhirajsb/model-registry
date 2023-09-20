package converter

import (
	"fmt"
	"github.com/opendatahub-io/model-registry/internal/ml_metadata/proto"
	"github.com/opendatahub-io/model-registry/internal/model/db"
	"github.com/opendatahub-io/model-registry/internal/model/graph"
	"strconv"
)

func ConvertEventPath(path []graph.EventStep) ([]db.EventPath, error) {
	var result []db.EventPath
	for _, step := range path {
		ep := db.EventPath{}
		switch s := step.(type) {
		case graph.EventStepIndex:
			index := s.Index
			ep.StepIndex = &index
		case graph.EventStepKey:
			key := s.Key
			ep.StepKey = &key
		}
		result = append(result, ep)
	}
	return result, nil
}

func ConvertToEventPath(pathSteps []db.EventPath) ([]graph.EventStep, error) {
	var result []graph.EventStep
	for _, step := range pathSteps {
		if step.StepIndex != nil {
			result = append(result, graph.EventStepIndex{Index: *step.StepIndex})
		}
		if step.StepKey != nil {
			result = append(result, graph.EventStepKey{Key: *step.StepKey})
		}
	}
	return result, nil
}

func ConvertToArtifactPropertyValue(prop *db.ArtifactProperty) (graph.Value, error) {
	return convertToPropertyValue(prop.Name, prop.IntValue, prop.DoubleValue,
		prop.StringValue, prop.ByteValue, prop.ProtoValue, prop.BoolValue, prop.TypeURL)
}

func ConvertToContextPropertyValue(prop *db.ContextProperty) (graph.Value, error) {
	return convertToPropertyValue(prop.Name, prop.IntValue, prop.DoubleValue,
		prop.StringValue, prop.ByteValue, prop.ProtoValue, prop.BoolValue, prop.TypeURL)
}

func ConvertToExecutionPropertyValue(prop *db.ExecutionProperty) (graph.Value, error) {
	return convertToPropertyValue(prop.Name, prop.IntValue, prop.DoubleValue,
		prop.StringValue, prop.ByteValue, prop.ProtoValue, prop.BoolValue, prop.TypeURL)
}

func ConvertArtifactProperties(props []*graph.InstanceProperty) ([]db.ArtifactProperty, error) {
	result := make([]db.ArtifactProperty, 0)
	for _, prop := range props {
		value := db.ArtifactProperty{
			// TODO handle copying artifact id externally, perhaps in the DB service
			//ExecutionID:      id,
			Name:             prop.Name,
			IsCustomProperty: prop.IsCustomProperty,
		}
		convertArtifactProperty(&value, prop.PropertyValue)
		result = append(result, value)
	}
	return result, nil
}

func ConvertToArtifactProperties(props []db.ArtifactProperty) ([]*graph.InstanceProperty, error) {
	result := make([]*graph.InstanceProperty, 0)
	for _, prop := range props {
		value := graph.InstanceProperty{
			Name:             prop.Name,
			IsCustomProperty: prop.IsCustomProperty,
		}
		propValue, err := ConvertToArtifactPropertyValue(&prop)
		if err != nil {
			return nil, err
		}
		value.PropertyValue = propValue
		result = append(result, &value)
	}
	return result, nil
}

func ConvertContextProperties(props []*graph.InstanceProperty) ([]db.ContextProperty, error) {
	result := make([]db.ContextProperty, 0)
	for _, prop := range props {
		value := db.ContextProperty{
			// TODO handle copying context id externally, perhaps in the DB service
			//ExecutionID:      id,
			Name:             prop.Name,
			IsCustomProperty: prop.IsCustomProperty,
		}
		convertContextProperty(&value, prop.PropertyValue)
		result = append(result, value)
	}
	return result, nil
}

func ConvertToContextProperties(props []db.ContextProperty) ([]*graph.InstanceProperty, error) {
	result := make([]*graph.InstanceProperty, 0)
	for _, prop := range props {
		value := graph.InstanceProperty{
			Name:             prop.Name,
			IsCustomProperty: prop.IsCustomProperty,
		}
		propValue, err := ConvertToContextPropertyValue(&prop)
		if err != nil {
			return nil, err
		}
		value.PropertyValue = propValue
		result = append(result, &value)
	}
	return result, nil
}

func ConvertExecutionProperties(props []*graph.InstanceProperty) ([]db.ExecutionProperty, error) {
	result := make([]db.ExecutionProperty, 0)
	for _, prop := range props {
		value := db.ExecutionProperty{
			// TODO handle copying exection id externally, perhaps in the DB service
			//ExecutionID:      id,
			Name:             prop.Name,
			IsCustomProperty: prop.IsCustomProperty,
		}
		convertExecutionProperty(&value, prop.PropertyValue)
		result = append(result, value)
	}
	return result, nil
}

func ConvertToExecutionProperties(props []db.ExecutionProperty) ([]*graph.InstanceProperty, error) {
	result := make([]*graph.InstanceProperty, 0)
	for _, prop := range props {
		value := graph.InstanceProperty{
			Name:             prop.Name,
			IsCustomProperty: prop.IsCustomProperty,
		}
		propValue, err := ConvertToExecutionPropertyValue(&prop)
		if err != nil {
			return nil, err
		}
		value.PropertyValue = propValue
		result = append(result, &value)
	}
	return result, nil
}

func convertToPropertyValue(name string, intValue *int64, doubleValue *float64,
	stringValue *string, byteValue *[]byte, protoValue *[]byte,
	boolValue *bool, typeURL *string) (graph.Value, error) {
	if intValue != nil {
		return graph.IntValue{Value: *intValue}, nil
	}
	if doubleValue != nil {
		return graph.DoubleValue{Value: *doubleValue}, nil
	}
	if stringValue != nil {
		return graph.StringValue{Value: *stringValue}, nil
	}
	if byteValue != nil {
		// TODO use proto Struct to unmarshal ByteValue and convert to StructValue
		return graph.StructValue{Value: nil}, nil
	}
	if protoValue != nil {
		return graph.ProtoValue{Value: &graph.ProtoTypeValue{
			TypeURL: typeURL,
			Value:   *protoValue,
		}}, nil
	}
	if boolValue != nil {
		return graph.BoolValue{Value: *boolValue}, nil
	}
	return nil, fmt.Errorf("unknown value type %s", name)
}

func ConvertArtifactState(state graph.ArtifactState) int8 {
	result, ok := proto.Artifact_State_value[state.String()]
	if !ok {
		result = int32(proto.Artifact_UNKNOWN.Number())
	}
	return int8(result)
}

func ConvertToArtifactState(state int8) graph.ArtifactState {
	result, ok := proto.Artifact_State_name[int32(state)]
	if !ok {
		result = string(graph.ArtifactStateUnknown)
	}
	return graph.ArtifactState(result)
}

func ConvertExecutionState(state graph.ExecutionState) int8 {
	result, ok := proto.Execution_State_value[state.String()]
	if !ok {
		result = int32(proto.Execution_UNKNOWN.Number())
	}
	return int8(result)
}

func ConvertToExecutionState(state int8) graph.ExecutionState {
	result, ok := proto.Execution_State_name[int32(state)]
	if !ok {
		result = string(graph.ExecutionStateUnknown)
	}
	return graph.ExecutionState(result)
}

func ConvertEventType(state graph.EventType) int8 {
	result, ok := proto.Event_Type_value[state.String()]
	if !ok {
		result = int32(proto.Event_UNKNOWN.Number())
	}
	return int8(result)
}

func ConvertToEventType(state int8) graph.EventType {
	result, ok := proto.Event_Type_name[int32(state)]
	if !ok {
		result = string(graph.EventTypeUnknown)
	}
	return graph.EventType(result)
}

func convertArtifactProperty(target *db.ArtifactProperty, source graph.Value) {
	switch v := source.(type) {
	case graph.IntValue:
		target.IntValue = &v.Value
	case graph.DoubleValue:
		target.DoubleValue = &v.Value
	case graph.StringValue:
		target.StringValue = &v.Value
	case graph.StructValue:
		target.ByteValue = ConvertStructToBytes(v.Value)
	case graph.ProtoValue:
		target.TypeURL = v.Value.TypeURL
		target.ByteValue = &v.Value.Value
	case graph.BoolValue:
		target.BoolValue = &v.Value
	}
}

func convertContextProperty(target *db.ContextProperty, source graph.Value) {
	switch v := source.(type) {
	case graph.IntValue:
		target.IntValue = &v.Value
	case graph.DoubleValue:
		target.DoubleValue = &v.Value
	case graph.StringValue:
		target.StringValue = &v.Value
	case graph.StructValue:
		target.ByteValue = ConvertStructToBytes(v.Value)
	case graph.ProtoValue:
		target.TypeURL = v.Value.TypeURL
		target.ByteValue = &v.Value.Value
	case graph.BoolValue:
		target.BoolValue = &v.Value
	}
}

func convertExecutionProperty(target *db.ExecutionProperty, source graph.Value) {
	switch v := source.(type) {
	case graph.IntValue:
		target.IntValue = &v.Value
	case graph.DoubleValue:
		target.DoubleValue = &v.Value
	case graph.StringValue:
		target.StringValue = &v.Value
	case graph.StructValue:
		target.ByteValue = ConvertStructToBytes(v.Value)
	case graph.ProtoValue:
		target.TypeURL = v.Value.TypeURL
		target.ByteValue = &v.Value.Value
	case graph.BoolValue:
		target.BoolValue = &v.Value
	}
}

func ConvertStructToBytes(value []*graph.StructTuple) *[]byte {
	// TODO
	return nil
}

const (
	IdPrefixLength    = 2
	ArtifactPrefix    = "AT"
	ContextPrefix     = "CT"
	ExecutionPrefix   = "EN"
	EventPrefix       = "ET"
	AttributionPrefix = "AR"
	AssociationPrefix = "AS"
	TypePrefix        = "TE"
)

func ConvertToArtifactID(id int64) string {
	// add type prefix
	return ConvertToGraphQLID(ArtifactPrefix, id)
}

func ConvertToContextID(id int64) string {
	// add type prefix
	return ConvertToGraphQLID(ContextPrefix, id)
}

func ConvertToExecutionID(id int64) string {
	// add type prefix
	return ConvertToGraphQLID(ExecutionPrefix, id)
}

func ConvertToEventID(id int64) string {
	// add type prefix
	return ConvertToGraphQLID(EventPrefix, id)
}

func ConvertToAttributionID(id int64) string {
	// add type prefix
	return ConvertToGraphQLID(AttributionPrefix, id)
}

func ConvertToAssociationID(id int64) string {
	// add type prefix
	return ConvertToGraphQLID(AssociationPrefix, id)
}

func ConvertToTypeID(id int64) string {
	// add type prefix
	return ConvertToGraphQLID(TypePrefix, id)
}

func ConvertArtifactID(id string) (int64, error) {
	if len(id) == 0 {
		return 0, nil
	}
	// check type prefix
	prefix := id[:IdPrefixLength]
	if prefix != ArtifactPrefix {
		return -1, fmt.Errorf("invalid artifact id: %s", prefix)
	}
	return strconv.ParseInt(id[IdPrefixLength:], 10, 64)
}

func ConvertContextID(id string) (int64, error) {
	if len(id) == 0 {
		return 0, nil
	}
	// check type prefix
	prefix := id[:IdPrefixLength]
	if prefix != ContextPrefix {
		return -1, fmt.Errorf("invalid context id: %s", prefix)
	}
	return strconv.ParseInt(id[IdPrefixLength:], 10, 64)
}

func ConvertExecutionID(id string) (int64, error) {
	if len(id) == 0 {
		return 0, nil
	}
	// check type prefix
	prefix := id[:IdPrefixLength]
	if prefix != ExecutionPrefix {
		return -1, fmt.Errorf("invalid execution id: %s", prefix)
	}
	return strconv.ParseInt(id[IdPrefixLength:], 10, 64)
}

func ConvertEventID(id string) (int64, error) {
	if len(id) == 0 {
		return 0, nil
	}
	// check type prefix
	prefix := id[:IdPrefixLength]
	if prefix != EventPrefix {
		return -1, fmt.Errorf("invalid event id: %s", prefix)
	}
	return strconv.ParseInt(id[IdPrefixLength:], 10, 64)
}

func ConvertAttributionID(id string) (int64, error) {
	if len(id) == 0 {
		return 0, nil
	}
	// check type prefix
	prefix := id[:IdPrefixLength]
	if prefix != AttributionPrefix {
		return -1, fmt.Errorf("invalid attribution id: %s", prefix)
	}
	return strconv.ParseInt(id[IdPrefixLength:], 10, 64)
}

func ConvertAssociationID(id string) (int64, error) {
	if len(id) == 0 {
		return 0, nil
	}
	// check type prefix
	prefix := id[:IdPrefixLength]
	if prefix != AssociationPrefix {
		return -1, fmt.Errorf("invalid association id: %s", prefix)
	}
	return strconv.ParseInt(id[IdPrefixLength:], 10, 64)
}

func ConvertTypeID(id string) (int64, error) {
	if len(id) == 0 {
		return 0, nil
	}
	// check type prefix
	prefix := id[:IdPrefixLength]
	if prefix != TypePrefix {
		return -1, fmt.Errorf("invalid type id: %s", prefix)
	}
	return strconv.ParseInt(id[IdPrefixLength:], 10, 64)
}

func ConvertToGraphQLID(typePrefix string, id int64) string {
	// remove type prefix
	return fmt.Sprintf("%s%v", typePrefix, id)
}

// FilterGraphQLProperties modifies props to remove fields from fieldsMap
func FilterGraphQLProperties(props *[]*graph.InstanceProperty, fieldsMap map[string]struct{}) {
	var filteredProps []*graph.InstanceProperty
	for _, prop := range *props {
		if _, ok := fieldsMap[prop.Name]; !ok {
			filteredProps = append(filteredProps, prop)
		}
	}
	*props = filteredProps
}

// FilterTypeProperties modifies props to remove type properties from fieldsMap, and returns them
func FilterTypeProperties(props *[]*graph.InstanceProperty, fieldsMap map[string]struct{}) map[string]*graph.InstanceProperty {
	result := make(map[string]*graph.InstanceProperty, 0)
	var filteredProps []*graph.InstanceProperty
	for _, prop := range *props {
		if _, ok := fieldsMap[prop.Name]; ok {
			result[prop.Name] = prop
		} else {
			filteredProps = append(filteredProps, prop)
		}
	}
	*props = filteredProps
	return result
}
