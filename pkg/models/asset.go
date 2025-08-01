package models

import (
	"encoding/json"

	iotsitewisetypes "github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
)

type DescribeAssetQuery struct {
	BaseQuery
}

type DescribeAssetPropertyQuery struct {
	BaseQuery
}

type ListAssetsQuery struct {
	BaseQuery
	ModelId string                            `json:"modelId,omitempty"`
	Filter  iotsitewisetypes.ListAssetsFilter `json:"filter,omitempty"`
}

type ListTimeSeriesQuery struct {
	BaseQuery
	TimeSeriesType iotsitewisetypes.ListTimeSeriesType `json:"timeSeriesType,omitempty"`
	AssetId        string                              `json:"assetId,omitempty"`
	AliasPrefix    string                              `json:"aliasPrefix,omitempty"`
}

type ListAssociatedAssetsQuery struct {
	BaseQuery
	HierarchyId     string `json:"hierarchyId,omitempty"`
	LoadAllChildren bool   `json:"loadAllChildren,omitempty"`
	// TraversalDirection is implied from the existence of HierarchyId
}

func GetDescribeAssetQuery(dq *backend.DataQuery) (*DescribeAssetQuery, error) {
	query := &DescribeAssetQuery{}
	if err := json.Unmarshal(dq.JSON, query); err != nil {
		return nil, err
	}

	// AssetId <--> AssetIds backward compatibility
	query.MigrateAssetProperty()

	// add on the DataQuery params
	query.QueryType = dq.QueryType

	return query, nil
}

func GetListAssetPropertiesQuery(dq *backend.DataQuery) (*ListAssetPropertiesQuery, error) {
	query := &ListAssetPropertiesQuery{}
	if err := json.Unmarshal(dq.JSON, query); err != nil {
		return nil, err
	}

	// AssetId <--> AssetIds backward compatibility
	query.MigrateAssetProperty()

	query.QueryType = dq.QueryType
	return query, nil
}

func GetListAssetsQuery(dq *backend.DataQuery) (*ListAssetsQuery, error) {
	query := &ListAssetsQuery{}
	if err := json.Unmarshal(dq.JSON, query); err != nil {
		return nil, err
	}

	// AssetId <--> AssetIds backward compatibility
	query.MigrateAssetProperty()

	// add on the DataQuery params
	query.MaxDataPoints = int32(dq.MaxDataPoints)
	query.QueryType = dq.QueryType
	return query, nil
}

func GetListTimeSeriesQuery(dq *backend.DataQuery) (*ListTimeSeriesQuery, error) {
	query := &ListTimeSeriesQuery{}
	if err := json.Unmarshal(dq.JSON, query); err != nil {
		return nil, err
	}

	// add on the DataQuery params
	query.MaxDataPoints = int32(dq.MaxDataPoints)
	query.QueryType = dq.QueryType
	return query, nil
}

func GetListAssociatedAssetsQuery(dq *backend.DataQuery) (*ListAssociatedAssetsQuery, error) {
	query := &ListAssociatedAssetsQuery{}
	if err := json.Unmarshal(dq.JSON, query); err != nil {
		return nil, err
	}

	// AssetId <--> AssetIds backward compatibility
	query.MigrateAssetProperty()

	// add on the DataQuery params
	query.MaxDataPoints = int32(dq.MaxDataPoints)
	query.QueryType = dq.QueryType
	return query, nil
}
