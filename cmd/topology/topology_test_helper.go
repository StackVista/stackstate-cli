package topology

import (
	sts "github.com/stackvista/stackstate-cli/generated/stackstate_api"
)

//nolint:mnd
func mockSnapshotResponse() sts.QuerySnapshotResult {
	return sts.QuerySnapshotResult{
		Type: "QuerySnapshotResult",
		ViewSnapshotResponse: map[string]interface{}{
			"_type": "ViewSnapshot",
			"components": []interface{}{
				map[string]interface{}{
					"id":               float64(229404307680647),
					"name":             "test-component",
					"typeIdentifier":   "urn:test:component-type:test",
					"layerIdentifier":  "urn:test:layer:test",
					"domainIdentifier": "urn:test:domain:test",
					"identifiers":      []interface{}{"urn:test:component:1"},
					"tags":             []interface{}{"service.namespace:test"},
					"state": map[string]interface{}{
						"healthState": "CRITICAL",
					},
				},
			},
			"metadata": map[string]interface{}{
				"componentTypes": []interface{}{
					map[string]interface{}{
						"identifier": "urn:test:component-type:test",
						"name":       "test type",
					},
				},
				"layers": []interface{}{
					map[string]interface{}{
						"identifier": "urn:test:layer:test",
						"name":       "Test Layer",
					},
				},
				"domains": []interface{}{
					map[string]interface{}{
						"identifier": "urn:test:domain:test",
						"name":       "Test Domain",
					},
				},
			},
		},
	}
}
