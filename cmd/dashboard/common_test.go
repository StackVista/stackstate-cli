package dashboard

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResolveDashboardIdOrUrn(t *testing.T) {
	tests := []struct {
		name           string
		id             int64
		identifier     string
		expectedResult string
		expectedError  string
	}{
		{
			name:           "Should resolve valid ID",
			id:             123,
			identifier:     "",
			expectedResult: "123",
			expectedError:  "",
		},
		{
			name:           "Should resolve valid identifier",
			id:             0,
			identifier:     "urn:custom:dashboard:test",
			expectedResult: "urn:custom:dashboard:test",
			expectedError:  "",
		},
		{
			name:           "Should prioritize ID when both are provided",
			id:             456,
			identifier:     "urn:custom:dashboard:test",
			expectedResult: "456",
			expectedError:  "",
		},
		{
			name:           "Should return error when neither is provided",
			id:             0,
			identifier:     "",
			expectedResult: "",
			expectedError:  "either --id or --identifier must be provided",
		},
		{
			name:           "Should resolve large ID",
			id:             9223372036854775807, // max int64
			identifier:     "",
			expectedResult: "9223372036854775807",
			expectedError:  "",
		},
		{
			name:           "Should resolve complex URN identifier",
			id:             0,
			identifier:     "urn:stackpack:kubernetes:dashboard:cluster-overview",
			expectedResult: "urn:stackpack:kubernetes:dashboard:cluster-overview",
			expectedError:  "",
		},
		{
			name:           "Should resolve identifier with special characters",
			id:             0,
			identifier:     "urn:custom:dashboard:test-name_with.special-chars",
			expectedResult: "urn:custom:dashboard:test-name_with.special-chars",
			expectedError:  "",
		},
		{
			name:           "Should handle empty string identifier as not provided",
			id:             0,
			identifier:     "",
			expectedResult: "",
			expectedError:  "either --id or --identifier must be provided",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ResolveDashboardIdOrUrn(tt.id, tt.identifier)

			assert.Equal(t, tt.expectedResult, result)

			if tt.expectedError != "" {
				assert.NotNil(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestResolveDashboardIdOrUrnPriority(t *testing.T) {
	// Test that ID takes priority over identifier when both are provided
	result, err := ResolveDashboardIdOrUrn(999, "urn:custom:dashboard:ignored")

	assert.Nil(t, err)
	assert.Equal(t, "999", result)
}

func TestResolveDashboardIdOrUrnEdgeCases(t *testing.T) {
	// Test negative ID (should still work as it's non-zero)
	result, err := ResolveDashboardIdOrUrn(-1, "")
	assert.Nil(t, err)
	assert.Equal(t, "-1", result)

	// Test with whitespace-only identifier (treated as empty)
	result, err = ResolveDashboardIdOrUrn(0, "   ")
	assert.Nil(t, err)
	assert.Equal(t, "   ", result) // The function doesn't trim whitespace
}
