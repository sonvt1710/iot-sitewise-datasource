package resource

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go/service/iotsitewise"
	"github.com/patrickmn/go-cache"
)

// cacheDuration is a constant that defines how long to keep cached elements before they are refreshed
const cacheDuration = time.Minute * 5

// cacheCleanupInterval is the interval at which the internal cache is cleaned / garbage collected
const cacheCleanupInterval = time.Minute * 10

type cachingProvider struct {
	resources *SitewiseResources
	cache     *cache.Cache
}

func NewCachingProvider(resources *SitewiseResources) cachingProvider {
	return cachingProvider{
		resources: resources,
		cache:     cache.New(cacheDuration, cacheCleanupInterval),
	}
}

func (cp *cachingProvider) Asset(ctx context.Context, assetId string) (*iotsitewise.DescribeAssetOutput, error) {
	val, ok := cp.cache.Get(assetId)
	if ok {
		a, ok := val.(iotsitewise.DescribeAssetOutput)
		if ok {
			return &a, nil
		}
	}

	a, err := cp.resources.Asset(ctx, assetId)
	if err != nil {
		return nil, err
	}
	cp.cache.Set(assetId, *a, -1)
	return a, nil
}

func (cp *cachingProvider) Property(ctx context.Context, assetId string, propertyId string) (*iotsitewise.DescribeAssetPropertyOutput, error) {
	key := assetId + "/" + propertyId
	val, ok := cp.cache.Get(key)
	if ok {
		a, ok := val.(iotsitewise.DescribeAssetPropertyOutput)
		if ok {
			return &a, nil
		}
	}

	a, err := cp.resources.Property(ctx, assetId, propertyId)
	if err != nil {
		return nil, err
	}
	cp.cache.Set(key, *a, -1)
	return a, nil
}

func (cp *cachingProvider) AssetModel(ctx context.Context, modelId string) (*iotsitewise.DescribeAssetModelOutput, error) {
	val, ok := cp.cache.Get(modelId)
	if ok {
		a, ok := val.(iotsitewise.DescribeAssetModelOutput)
		if ok {
			return &a, nil
		}
	}

	a, err := cp.resources.AssetModel(ctx, modelId)
	if err != nil {
		return nil, err
	}
	cp.cache.Set(modelId, *a, -1)
	return a, nil
}
