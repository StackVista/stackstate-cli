package printer

import (
	"fmt"

	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
)

func metricValueOrDash(bucket []stackstate_api.MetricBucketValue, index int) interface{} {
	if index < len(bucket) && bucket[index].HasValue() {
		return bucket[index].GetValue()
	}
	return "-"
}

func MetricBucketToJson(name string, bucket []stackstate_api.MetricBucketValue, size int32) map[string]interface{} {
	return map[string]interface{}{
		"name":                               name,
		fmt.Sprintf("now-%d", size):          metricValueOrDash(bucket, 0),
		fmt.Sprintf("%d-%d", size, 2*size):   metricValueOrDash(bucket, 1), //nolint:mnd
		fmt.Sprintf("%d-%d", 2*size, 3*size): metricValueOrDash(bucket, 2), //nolint:mnd
	}
}

func MetricBucketToRow(name string, bucket []stackstate_api.MetricBucketValue) []interface{} {
	return []interface{}{
		name,
		metricValueOrDash(bucket, 0),
		metricValueOrDash(bucket, 1),
		metricValueOrDash(bucket, 2), //nolint:mnd
	}
}
