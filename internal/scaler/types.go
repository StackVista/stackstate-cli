package scaler

import (
	"strconv"
	"time"

	"gopkg.in/yaml.v3"
)

const (
	ChecksumDataKey     = "checksum"
	DateDataKey         = "date"
	ScaledDownDataKey   = "scaledDown"
	OriginalReplicasKey = "originalReplicas"
)

type ScaleStatus struct {
	Checksum         string
	Date             time.Time
	ScaledDown       bool
	OriginalReplicas map[string]ReplicaCount
}

func EmptyStatus() *ScaleStatus {
	return &ScaleStatus{
		Checksum:         "",
		Date:             time.Now(),
		ScaledDown:       false,
		OriginalReplicas: map[string]ReplicaCount{},
	}
}

func StatusFromConfigMapData(data map[string]string) (*ScaleStatus, error) {
	date, err := time.Parse(time.RFC3339, data[DateDataKey])
	if err != nil {
		return nil, err
	}

	scaled, err := strconv.ParseBool(data[ScaledDownDataKey])
	if err != nil {
		return nil, err
	}

	m := map[string]ReplicaCount{}
	if err := yaml.Unmarshal([]byte(data[OriginalReplicasKey]), &m); err != nil {
		return nil, err
	}

	return &ScaleStatus{
		Checksum:         data[ChecksumDataKey],
		Date:             date,
		ScaledDown:       scaled,
		OriginalReplicas: m,
	}, nil
}

func StatusToConfigMapData(status *ScaleStatus, data map[string]string) (map[string]string, error) {
	b, err := yaml.Marshal(status.OriginalReplicas)
	if err != nil {
		return nil, err
	}

	data[ChecksumDataKey] = status.Checksum
	data[DateDataKey] = status.Date.Format(time.RFC3339)
	data[OriginalReplicasKey] = string(b)
	data[ScaledDownDataKey] = strconv.FormatBool(status.ScaledDown)

	return data, nil
}
