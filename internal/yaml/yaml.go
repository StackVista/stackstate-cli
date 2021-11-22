package yaml

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/mitchellh/copystructure"
	"gopkg.in/yaml.v3"
)

const (
	Indent = 2
)

type ErrNoSuchKey struct {
	key string
}

func (e ErrNoSuchKey) Error() string {
	return fmt.Sprintf("no such key '%s' in doc", e.key)
}

type ErrNotADoc struct {
	key   string
	value interface{}
}

func (e ErrNotADoc) Error() string {
	return fmt.Sprintf("'%v' is not a yaml (sub)document for finding key '%s'", e.value, e.key)
}

type YamlDoc map[string]interface{} //nolint:golint

func (y YamlDoc) Lookup(key string) (interface{}, error) {
	split := parsePath(key)

	return lookup(y, split)
}

func lookup(y YamlDoc, splitKey []string) (interface{}, error) {
	if v, ok := y[splitKey[0]]; ok {
		if len(splitKey) == 1 {
			return v, nil
		}

		switch val := v.(type) {
		case map[string]interface{}:
			return lookup(YamlDoc(val), splitKey[1:])
		case YamlDoc:
			return lookup(val, splitKey[1:])
		default:
			return nil, ErrNotADoc{joinPath(splitKey), val}
		}
	}

	return nil, ErrNoSuchKey{splitKey[0]}
}

func (y YamlDoc) SetValue(key string, value interface{}) {
	split := parsePath(key)

	var curDoc = y
	for i, k := range split {
		if len(split)-1 == i {
			curDoc[k] = value
			break
		}

		curDoc = curDoc.getOrCreateMap(k)
	}
}

// AppendArray appends given values to the array present at key in the YamlDoc
func (y YamlDoc) AppendArray(key string, values ...interface{}) error {
	split := parsePath(key)

	var curDoc = y
	for i, k := range split {
		if len(split)-1 == i {
			if v, ok := curDoc[k]; ok {
				if arr, ok := v.([]interface{}); ok {
					curDoc[k] = append(arr, values...)
				} else {
					return fmt.Errorf("value at '%s' was not an []interface{}, but %T", key, v)
				}
			} else {
				arr := values
				curDoc[k] = arr
			}
			break
		}

		curDoc = curDoc.getOrCreateMap(k)
	}

	return nil
}

func (y YamlDoc) Merge(x YamlDoc) (YamlDoc, error) {
	return mergeMaps(y, x)
}

func mergeMaps(left, right interface{}) (map[string]interface{}, error) {
	leftMap, err := asMap(left)
	if err != nil {
		return nil, err
	}

	rightMap, err := asMap(right)
	if err != nil {
		return nil, err
	}

	leftCopy, err := copyMap(leftMap)
	if err != nil {
		return nil, err
	}

	for k, v := range rightMap {
		// if k not present in leftCopy
		if _, ok := leftCopy[k]; !ok {
			vCopy, err := copystructure.Copy(v)
			if err != nil {
				return nil, err
			}
			leftCopy[k] = vCopy
		}

		switch rightV := v.(type) {
		case map[string]interface{}:
			m, err := mergeMaps(leftCopy[k], rightV)
			if err != nil {
				return nil, err
			}
			leftCopy[k] = m
		case YamlDoc:
			m, err := mergeMaps(leftCopy[k], rightV)
			if err != nil {
				return nil, err
			}
			leftCopy[k] = m
		case []interface{}:
			m, err := mergeSlices(leftCopy[k], rightV)
			if err != nil {
				return nil, err
			}
			leftCopy[k] = m
		default:
			rightCopy, err := copystructure.Copy(rightV)
			if err != nil {
				return nil, err
			}
			leftCopy[k] = rightCopy
		}
	}

	return leftCopy, nil
}

func asMap(i interface{}) (map[string]interface{}, error) {
	if m, ok := i.(map[string]interface{}); ok {
		return m, nil
	}

	if y, ok := i.(YamlDoc); ok {
		return y, nil
	}

	return nil, fmt.Errorf("cannot cast '%T' to map[string]interface{}", i)
}

func copyMap(m map[string]interface{}) (map[string]interface{}, error) {
	copy, err := copystructure.Copy(m)
	if err != nil {
		return nil, err
	}

	mCopy, ok := copy.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("copying did not deliver a map[string]interface{}")
	}

	return mCopy, nil
}

func mergeSlices(left, right interface{}) ([]interface{}, error) {
	leftArr, ok := left.([]interface{})
	if !ok {
		return nil, fmt.Errorf("left is not an []interface{}, but %T", left)
	}

	rightArr, ok := right.([]interface{})
	if !ok {
		return nil, fmt.Errorf("right is not an []interface{}, but %T", right)
	}

	leftArr, err := copySlice(leftArr)
	if err != nil {
		return nil, err
	}

	return append(leftArr, rightArr...), nil
}

func copySlice(i []interface{}) ([]interface{}, error) {
	copy, err := copystructure.Copy(i)
	if err != nil {
		return nil, err
	}

	iCopy, ok := copy.([]interface{})
	if !ok {
		return nil, fmt.Errorf("copying did not deliver an []interface{}")
	}

	return iCopy, nil
}

// EncodeToWriter encodes the YamlDoc to the given writer
func (y YamlDoc) EncodeToWriter(buffer io.Writer) error {
	encoder := yaml.NewEncoder(buffer)
	defer encoder.Close()

	encoder.SetIndent(Indent)

	if err := encoder.Encode(y); err != nil {
		return errors.New("could not encode YAML")
	}

	return nil
}

// Encode encodes the YamlDoc to a byte array
func (y YamlDoc) Encode() ([]byte, error) {
	byteBuffer := &bytes.Buffer{}

	if err := y.EncodeToWriter(byteBuffer); err != nil {
		return nil, err
	}
	return byteBuffer.Bytes(), nil
}

// EncodeToFile writes the encoded representation of this YamlDoc to the specified file
func (y YamlDoc) EncodeToFile(file string) error {
	c, err := y.Encode()
	if err != nil {
		return err
	}

	return ioutil.WriteFile(file, c, 0644) //nolint:gosec
}

// DecodeFromReader reads the encoded yaml representation from the reader and returns a decoded YamlDoc
func DecodeFromReader(r io.Reader) (YamlDoc, error) {
	var yamlContents map[string]interface{}

	if err := DecodeObjectFromReader(r, &yamlContents); err != nil {
		return nil, err
	}

	return yamlContents, nil
}

func DecodeObjectFromReader(r io.Reader, o interface{}) error {
	reader := bufio.NewReader(r)
	decoder := yaml.NewDecoder(reader)

	if err := decoder.Decode(o); err != nil {
		return err
	}

	return nil
}

func Decode(b []byte) (map[string]interface{}, error) {
	var yamlContents map[string]interface{}

	if err := yaml.Unmarshal(b, &yamlContents); err != nil {
		return nil, err
	}

	return yamlContents, nil
}

func (y YamlDoc) getOrCreateMap(key string) YamlDoc {
	if v, ok := y[key]; ok {
		return v.(YamlDoc)
	}

	y[key] = YamlDoc{}
	return y[key].(YamlDoc)
}

func parsePath(key string) []string {
	return strings.Split(key, ".")
}

func joinPath(keys []string) string {
	return strings.Join(keys, ".")
}
