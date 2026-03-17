package topology

import (
	sts "github.com/stackvista/stackstate-cli/generated/stackstate_api"
)

func mockSnapshotResponse() sts.QuerySnapshotResult {
	return sts.QuerySnapshotResult{
		Type: "QuerySnapshotResult",
		ViewSnapshotResponse: map[string]interface{}{
			"_type": "ViewSnapshot",
			"components": []interface{}{
				map[string]interface{}{
					"id":          float64(229404307680647),
					"name":        "test-component",
					"type":        float64(239975151751041),
					"layer":       float64(186771622698247),
					"domain":      float64(209616858431909),
					"identifiers": []interface{}{"urn:test:component:1"},
					"tags":        []interface{}{"service.namespace:test"},
					"state": map[string]interface{}{
						"healthState": "CRITICAL",
					},
				},
			},
			"metadata": map[string]interface{}{
				"componentTypes": []interface{}{
					map[string]interface{}{
						"id":   float64(239975151751041),
						"name": "test type",
					},
				},
				"layers": []interface{}{
					map[string]interface{}{
						"id":   float64(186771622698247),
						"name": "Test Layer",
					},
				},
				"domains": []interface{}{
					map[string]interface{}{
						"id":   float64(209616858431909),
						"name": "Test Domain",
					},
				},
			},
		},
	}
}
