package values

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/yaml"
)

func TestShouldListProfiles(t *testing.T) {
	profiles, err := ListProfiles()
	assert.NoError(t, err)

	assert.Len(t, profiles, 4)
	assert.Contains(t, profiles, "ha")
	assert.Contains(t, profiles, "non-ha")
	assert.Contains(t, profiles, "development")
	assert.Contains(t, profiles, "development-split")
}

func TestShouldLoadProfile(t *testing.T) {
	yamlProfile, err := LoadProfile("development")
	assert.NoError(t, err)

	v, err := yamlProfile.Lookup("elasticsearch.replicas")
	assert.NoError(t, err)
	assert.Equal(t, 1, v)
	r, err := yamlProfile.Lookup("hbase.hbase.master.replicaCount")
	assert.NoError(t, err)
	assert.Equal(t, 1, r)
}

func TestShouldApplyProfile(t *testing.T) {
	baseYaml := `stackstate:
  admin:
    authentication:
      password: "foobar"
`

	base, err := yaml.Decode([]byte(baseYaml))
	assert.NoError(t, err)
	values, err := ApplyProfile(base, "development")
	assert.NoError(t, err)
	v, err := values.Lookup("stackstate.admin.authentication.password")
	assert.NoError(t, err)
	assert.Equal(t, v, "foobar")
	w, err := values.Lookup("hbase.hbase.master.replicaCount")
	assert.NoError(t, err)
	assert.Equal(t, w, 1)
}
