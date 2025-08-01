package fields

import (
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise"
	iotsitewisetypes "github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"

	"github.com/grafana/grafana-plugin-sdk-go/data"
	"github.com/grafana/iot-sitewise-datasource/pkg/models"
	"github.com/grafana/iot-sitewise-datasource/pkg/util"
)

func NewFieldWithName(name string, fieldType data.FieldType, length int) *data.Field {
	field := data.NewFieldFromFieldType(fieldType, length)
	field.Name = name
	return field
}

func NameField(length int) *data.Field {
	return NewFieldWithName(Name, data.FieldTypeString, length)
}

func IdField(length int) *data.Field {
	return NewFieldWithName(Id, data.FieldTypeString, length)
}

func ArnField(length int) *data.Field {
	return NewFieldWithName(Arn, data.FieldTypeString, length)
}

func ModelIdField(length int) *data.Field {
	return NewFieldWithName(ModelId, data.FieldTypeString, length)
}

// Description fields are optional
func DescriptionField(length int) *data.Field {
	return NewFieldWithName(Description, data.FieldTypeNullableString, length)
}

func StatusErrorField(length int) *data.Field {
	return NewFieldWithName(StatusError, data.FieldTypeNullableString, length)
}

func StatusStateField(length int) *data.Field {
	return NewFieldWithName(StatusState, data.FieldTypeString, length)
}

func HierarchiesField(length int) *data.Field {
	return NewFieldWithName(Hierarchies, data.FieldTypeString, length)
}

func CreationDateField(length int) *data.Field {
	return NewFieldWithName(CreationDate, data.FieldTypeTime, length)
}

func LastUpdateField(length int) *data.Field {
	return NewFieldWithName(LastUpdate, data.FieldTypeTime, length)
}

func TimeField(length int) *data.Field {
	return NewFieldWithName(Time, data.FieldTypeTime, length)
}

func QualityField(length int) *data.Field {
	return NewFieldWithName(Quality, data.FieldTypeString, length)
}

func PropertiesField(length int) *data.Field {
	return NewFieldWithName(Properties, data.FieldTypeString, length)
}

func CompositeModelsField(length int) *data.Field {
	return NewFieldWithName(CompositeModels, data.FieldTypeString, length)
}

func PropertyValueField(property *iotsitewise.DescribeAssetPropertyOutput, length int) *data.Field {
	propertyName := util.GetPropertyName(property)

	return PropertyValueFieldNamed(propertyName, property, length)
}

func PropertyValueFieldForQuery(query models.AssetPropertyValueQuery, property *iotsitewise.DescribeAssetPropertyOutput, length int) *data.Field {
	if models.QueryTypePropertyAggregate == query.QueryType {
		return PropertyValueFieldNamed("raw", property, length)
	} else {
		return PropertyValueField(property, length)
	}
}

func PropertyValueFieldNamed(name string, property *iotsitewise.DescribeAssetPropertyOutput, length int) *data.Field {
	unit := util.GetPropertyUnit(property)

	valueField := NewFieldWithName(name, FieldTypeForPropertyValue(property), length)
	valueField.Config = &data.FieldConfig{
		Unit: ToGrafanaUnit(&unit),
	}

	return valueField
}

func DatumField(length int, col iotsitewisetypes.ColumnInfo) *data.Field {
	return NewFieldWithName(*col.Name, FieldTypeForQueryResult(col), length)
}

func AggregationField(length int, name string) *data.Field {
	return NewFieldWithName(name, data.FieldTypeFloat64, length)
}

func AnomalyScoreField(length int) *data.Field {
	return NewFieldWithName(AnomalyScore, data.FieldTypeFloat64, length)
}

func PredictionReasonField(length int) *data.Field {
	return NewFieldWithName(PredictionReason, data.FieldTypeString, length)
}

func DiagnosticField(length int, assetId string) *data.Field {
	return NewFieldWithName(assetId, data.FieldTypeFloat64, length)
}

// for time series

func AliasField(length int) *data.Field {
	return NewFieldWithName(Alias, data.FieldTypeString, length)
}

func AssetIdField(length int) *data.Field {
	return NewFieldWithName(AssetId, data.FieldTypeString, length)
}

func DataTypeField(length int) *data.Field {
	return NewFieldWithName(DataType, data.FieldTypeString, length)
}

func DataTypeSpecField(length int) *data.Field {
	return NewFieldWithName(DataTypeSpec, data.FieldTypeString, length)
}

func PropertyIdField(length int) *data.Field {
	return NewFieldWithName(PropertyId, data.FieldTypeString, length)
}

func TimeSeriesArnField(length int) *data.Field {
	return NewFieldWithName(TimeSeriesArn, data.FieldTypeString, length)
}

func TimeSeriesIdField(length int) *data.Field {
	return NewFieldWithName(TimeSeriesId, data.FieldTypeString, length)
}

func TimeSeriesCreationDateField(length int) *data.Field {
	return NewFieldWithName(TimeSeriesCreationDate, data.FieldTypeTime, length)
}

func TimeSeriesLastUpdateDateField(length int) *data.Field {
	return NewFieldWithName(TimeSeriesLastUpdateDate, data.FieldTypeTime, length)
}
