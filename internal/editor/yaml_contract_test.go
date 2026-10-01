package editor

import (
	"fmt"
	"github.com/go-openapi/swag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.yaml.in/yaml/v3"
	"strings"
	"testing"
)

func TestOpenAPIOrderedJSONRoundTrip(t *testing.T) {
	const fixture = `{"z":null,"a":{"enabled":true,"number":12,"string":"0123"},"text":"first\nsecond\n"}`
	var value swag.JSONMapSlice
	require.NoError(t, swag.ReadJSON([]byte(fixture), &value))
	actual, err := swag.WriteJSON(value)
	require.NoError(t, err)
	assert.Equal(t, fixture, string(actual))
}

func TestOpenAPIYAMLNodeConversion(t *testing.T) {
	const fixture = `z: null
a: &values
  enabled: true
  number: 12
  string: "0123"
  legacy: yes
copy: *values
text: |
  first
  second
`
	var node yaml.Node
	require.NoError(t, yaml.Unmarshal([]byte(fixture), &node))
	actual, err := swag.YAMLToJSON(node.Content[0])
	require.NoError(t, err)
	const expected = `{"z":null,"a":{"enabled":true,"number":12,"string":"0123","legacy":"yes"},"copy":{"enabled":true,"number":12,"string":"0123","legacy":"yes"},"text":"first\nsecond\n"}`
	assert.Equal(t, expected, string(actual))
}

func TestOpenAPIRejectsExcessiveJSONDepth(t *testing.T) {
	const depth = 20000
	payload := `{"a":` + strings.Repeat("[", depth) + strings.Repeat("]", depth) + `}`
	var value swag.JSONMapSlice
	require.Error(t, swag.ReadJSON([]byte(payload), &value))
}

func TestOpenAPIRejectsExcessiveYAMLAliases(t *testing.T) {
	var fixture strings.Builder
	fixture.WriteString("a0: &a0 \"value\"\n")
	for i := 1; i <= 20; i++ {
		fmt.Fprintf(&fixture, "a%d: &a%d [*a%d, *a%d]\n", i, i, i-1, i-1)
	}
	doc, err := swag.BytesToYAMLDoc([]byte(fixture.String()))
	require.NoError(t, err)
	_, err = swag.YAMLToJSON(doc)
	require.ErrorContains(t, err, "excessive aliasing")
}
