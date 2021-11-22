package yaml

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = map[string]interface{}{
	"nested.key":             1,
	"nested.anotherkey":      7,
	"nested.deeper.sub":      2,
	"nested.arr":             []interface{}{3, 4},
	"nested.double.deep.sub": 6,
	"base":                   5,
}

func runTestSuite(t *testing.T, base string, d YamlDoc) {
	for k, expected := range tests {
		k := k
		expected := expected
		t.Run(strings.Join([]string{base, k}, "/"), func(t *testing.T) {
			v, err := d.Lookup(k)
			assert.NoError(t, err)
			assert.Equal(t, expected, v)
		})
	}
}

func TestLookup(t *testing.T) {
	d := YamlDoc{
		"nested": map[string]interface{}{
			"key":        1,
			"anotherkey": 7,
			"deeper": map[string]interface{}{
				"sub": 2,
			},
			"double": map[string]interface{}{
				"deep": map[string]interface{}{
					"sub": 6,
				},
			},
			"arr": []interface{}{
				3,
				4,
			},
		},
		"base": 5,
	}

	runTestSuite(t, "Lookup", d)
}

func TestSetValues(t *testing.T) {
	d := YamlDoc{}
	d.SetValue("nested.key", 1)
	d.SetValue("nested.anotherkey", 7)
	d.SetValue("nested.deeper.sub", 2)
	d.SetValue("nested.double.deep.sub", 6)
	d.SetValue("nested.arr", []interface{}{3, 4})
	d.SetValue("base", 5)

	runTestSuite(t, "SetValue", d)
}

func TestAppendArray(t *testing.T) {
	d := YamlDoc{}
	d.SetValue("nested.key", 1)
	d.SetValue("nested.anotherkey", 7)
	d.SetValue("nested.deeper.sub", 2)
	d.SetValue("nested.double.deep.sub", 6)
	assert.NoError(t, d.AppendArray("nested.arr", 3))
	assert.NoError(t, d.AppendArray("nested.arr", 4))
	d.SetValue("base", 5)

	runTestSuite(t, "AppendArray", d)
}

func TestDecode(t *testing.T) {
	b, err := ioutil.ReadFile("sample.yaml")
	assert.NoError(t, err)

	y, err := Decode(b)
	assert.NoError(t, err)

	runTestSuite(t, "Decode", y)
}

func TestMergeSimple(t *testing.T) {
	x := YamlDoc{}
	x.SetValue("nested.key", 1)
	x.SetValue("nested.anotherkey", 7)
	x.SetValue("nested.deeper.sub", 2)
	x.SetValue("nested.double.deep.sub", 6)
	x.SetValue("nested.arr", []interface{}{3, 4})

	y := YamlDoc{}
	y.SetValue("base", 5)

	z, err := x.Merge(y)
	assert.NoError(t, err)

	runTestSuite(t, "MergeSimple", z)
}

func TestMergeArray(t *testing.T) {
	x := YamlDoc{}
	x.SetValue("nested.key", 1)
	x.SetValue("nested.anotherkey", 7)
	x.SetValue("nested.deeper.sub", 2)
	x.SetValue("nested.double.deep.sub", 6)
	x.SetValue("nested.arr", []interface{}{3})
	x.SetValue("base", 5)

	y := YamlDoc{}
	y.SetValue("nested.arr", []interface{}{4})

	z, err := x.Merge(y)
	assert.NoError(t, err)

	runTestSuite(t, "MergeArray", z)
}

func TestMergeDeepMap(t *testing.T) {
	x := YamlDoc{}
	x.SetValue("nested.key", 1)
	x.SetValue("nested.anotherkey", 7)
	x.SetValue("nested.deeper.sub", 2)
	x.SetValue("nested.arr", []interface{}{3, 4})
	x.SetValue("base", 5)

	y := YamlDoc{}
	y.SetValue("nested.double.deep.sub", 6)

	z, err := x.Merge(y)
	assert.NoError(t, err)

	runTestSuite(t, "MergeArray", z)
}

func TestMergeMap(t *testing.T) {
	x := YamlDoc{}
	x.SetValue("nested.key", 1)
	x.SetValue("nested.anotherkey", "wrong")
	x.SetValue("nested.deeper.sub", 2)
	x.SetValue("nested.double.deep.sub", 6)
	x.SetValue("nested.arr", []interface{}{3, 4})
	x.SetValue("base", 5)

	y := YamlDoc{}
	y.SetValue("nested.anotherkey", 7)

	z, err := x.Merge(y)
	assert.NoError(t, err)

	runTestSuite(t, "MergeArray", z)
}
