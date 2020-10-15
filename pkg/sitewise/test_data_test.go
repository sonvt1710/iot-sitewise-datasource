package sitewise

import (
	"context"
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/backend"

	"github.com/grafana/iot-sitewise-datasource/pkg/testutil"

	"github.com/grafana/iot-sitewise-datasource/pkg/sitewise/client"

	"github.com/grafana/iot-sitewise-datasource/pkg/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iotsitewise"
)

type testDataFunc func(t *testing.T, client client.Client) interface{}

// How to run tests:
//
// Use shell environment variables. Ex:
// export AWS_ACCESS_KEY_ID="<key id>"
// export AWS_SECRET_ACCESS_KEY="<secret key>"
// export AWS_SESSION_TOKEN="<session token>"
//
func TestGenerateTestData(t *testing.T) {

	t.Skip("Integration Test") // comment line to run this

	m := make(map[string]testDataFunc)

	m["property-history-values.json"] = func(t *testing.T, client client.Client) interface{} {
		ctx := context.Background()

		// hard coded values from my account
		query := models.AssetPropertyValueQuery{}
		query.AssetId = testutil.TestAssetId
		query.PropertyId = testutil.TestPropIdAvgWind
		query.TimeRange = backend.TimeRange{
			From: time.Now().Add(time.Hour * -3), // return 3 hours of data. 60*3/5 = 36 points
			To:   time.Now(),
		}

		resp, err := GetAssetPropertyValues(ctx, client, query)
		if err != nil {
			t.Fatal(err)
		}

		return resp
	}
	m["property-value.json"] = func(t *testing.T, client client.Client) interface{} {
		ctx := context.Background()

		query := models.AssetPropertyValueQuery{}
		query.AssetId = testutil.TestAssetId
		query.PropertyId = testutil.TestPropIdAvgWind

		resp, err := GetAssetPropertyValue(ctx, client, query)
		if err != nil {
			t.Fatal(err)
		}

		return resp
	}
	m["property-aggregate-values.json"] = func(t *testing.T, client client.Client) interface{} {
		ctx := context.Background()

		query := models.AssetPropertyValueQuery{}
		query.Resolution = "1m"
		query.AggregateTypes = []string{"AVERAGE", "MAXIMUM", "MINIMUM"}
		query.AssetId = testutil.TestAssetId
		query.PropertyId = testutil.TestPropIdRawWin
		query.TimeRange = backend.TimeRange{
			From: time.Now().Add(time.Hour * -3), // return 3 hours of data. 60*3/5 = 36 points
			To:   time.Now(),
		}

		resp, err := GetAssetPropertyAggregates(ctx, client, query)
		if err != nil {
			t.Fatal(err)
		}

		return resp
	}
	m["describe-asset.json"] = func(t *testing.T, client client.Client) interface{} {
		ctx := context.Background()
		resp, err := GetAssetDescription(ctx, client, models.DescribeAssetQuery{AssetId: testutil.TestAssetId})
		if err != nil {
			t.Fatal(err)
		}
		return resp
	}
	m["describe-asset-property-avg-wind.json"] = func(t *testing.T, client client.Client) interface{} {
		ctx := context.Background()
		resp, err := GetAssetPropertyDescription(ctx, client, models.DescribeAssetPropertyQuery{
			AssetId:    testutil.TestAssetId,
			PropertyId: testutil.TestPropIdAvgWind,
		})
		if err != nil {
			t.Fatal(err)
		}
		return resp
	}

	m["describe-asset-property-raw-wind.json"] = func(t *testing.T, client client.Client) interface{} {
		ctx := context.Background()
		resp, err := GetAssetPropertyDescription(ctx, client, models.DescribeAssetPropertyQuery{
			AssetId:    testutil.TestAssetId,
			PropertyId: testutil.TestPropIdRawWin,
		})
		if err != nil {
			t.Fatal(err)
		}
		return resp
	}

	sesh := session.Must(session.NewSession())
	sw := iotsitewise.New(sesh, aws.NewConfig().WithRegion("us-east-1"))

	for k, v := range m {
		writeTestData(k, v, sw, t)
	}
}

func writeTestData(filename string, tf testDataFunc, client client.Client, t *testing.T) {

	resp := tf(t, client)

	js, err := json.MarshalIndent(resp, "", "    ")

	if err != nil {
		t.Fatal(err)
	}

	f, err := os.Create("../testdata/" + filename)

	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		cerr := f.Close()
		if err == nil {
			err = cerr
		}
	}()

	_, err = f.Write(js)
	if err != nil {
		t.Fatal(err)
	}

}
