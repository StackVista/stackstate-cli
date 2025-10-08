package otelmapping_test

import (
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
)

const (
	TestMetricValue    = float64(23.5)
	TestComponentCount = int64(200)
	TestRelationCount  = int64(100)
)

var (
	TestError1     = stackstate_api.NewOtelMappingError(stackstate_api.MESSAGELEVEL_WARN, "Oops")
	TestError2     = stackstate_api.NewOtelMappingError(stackstate_api.MESSAGELEVEL_ERROR, "Oops2")
	TestSomeErrors = []stackstate_api.OtelMappingError{
		*TestError1,
		*TestError2,
	}
	TestMetricBucketValue = (func() *stackstate_api.MetricBucketValue {
		v := stackstate_api.NewMetricBucketValue()
		v.SetValue(float64(TestMetricValue))
		return v
	})()
	TestSomeMetrics = stackstate_api.OtelMappingMetrics{
		BucketSizeSeconds: 10,
		LatencySeconds:    []stackstate_api.MetricBucketValue{*TestMetricBucketValue, *TestMetricBucketValue, *TestMetricBucketValue},
	}
	TestSomeOtelMappingStatusItem = (func() *stackstate_api.OtelMappingStatusItem {
		v := stackstate_api.NewOtelMappingStatusItem("name", TestRelationCount, TestComponentCount)
		v.SetIdentifier("identifier")
		return v
	})()
	TestSomeOtelMappingStatus = (func() *stackstate_api.OtelMappingStatus {
		v := stackstate_api.NewOtelMappingStatus(*TestSomeOtelMappingStatusItem, TestSomeErrors)
		v.SetMetrics(TestSomeMetrics)
		return v
	})()
)

var (
	TestSomeOtelMappingItem = (func() *stackstate_api.OtelMappingItem {
		v := stackstate_api.NewOtelMappingItem("name")
		v.SetIdentifier("identifier")
		return v
	})()
	TestSomeOtelMappingItem2 = (func() *stackstate_api.OtelMappingItem {
		v := stackstate_api.NewOtelMappingItem("name2")
		v.SetIdentifier("identifier2")
		return v
	})()
	TestAllMappingItems = []stackstate_api.OtelMappingItem{*TestSomeOtelMappingItem, *TestSomeOtelMappingItem2}
)
