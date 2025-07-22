package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	//nolint:lll
	selfSignedBase64Cert = "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNKekNDQWRHZ0F3SUJBZ0lVVi9hSmoxZkVjQ2dOVTJGYWZZMHVSTHF5N21Bd0RRWUpLb1pJaHZjTkFRRUwKQlFBd0tURW5NQ1VHQTFVRUF3d2VkbWxzYVdGcmIzWXVjMkZ1WkdKdmVDNXpkR0ZqYTNOMFlYUmxMbWx2TUNBWApEVEkxTURjeU1USXdORFUxTmxvWUR6SXhNalV3TmpJM01qQTBOVFUyV2pBcE1TY3dKUVlEVlFRRERCNTJhV3hwCllXdHZkaTV6WVc1a1ltOTRMbk4wWVdOcmMzUmhkR1V1YVc4d1hEQU5CZ2txaGtpRzl3MEJBUUVGQUFOTEFEQkkKQWtFQW9wUXVPSmZJa0xDV0pLVDcwaGdiSEpwVWtFQitaYTJwOXVBMUlOUktNNEFyN2RjVjltdXhOS09jSloycwpWdCtiK1lTS1c4cnRteE5QUVh1RTJENHRlUUlEQVFBQm80SE9NSUhMTUIwR0ExVWREZ1FXQkJRVTBPTFZRRzEyCndNb0VLSGdxSG1aeVhTelozekFmQmdOVkhTTUVHREFXZ0JRVTBPTFZRRzEyd01vRUtIZ3FIbVp5WFN6WjN6QVAKQmdOVkhSTUJBZjhFQlRBREFRSC9NSGdHQTFVZEVRUnhNRytDSG5acGJHbGhhMjkyTG5OaGJtUmliM2d1YzNSaApZMnR6ZEdGMFpTNXBiNElqYjNSc2NDMTJhV3hwWVd0dmRpNXpZVzVrWW05NExuTjBZV05yYzNSaGRHVXVhVytDCktHOTBiSEF0YUhSMGNDMTJhV3hwWVd0dmRpNXpZVzVrWW05NExuTjBZV05yYzNSaGRHVXVhVzh3RFFZSktvWkkKaHZjTkFRRUxCUUFEUVFBZllBVk1lTVJHbFcrR1prellPeGRIaVhYNEFISHA5SWxvWlBMbUJHNExtdlpDODBoVgpLNGNSVUVHSGtSeGdrMGgwYzl3RDhOZFZSM1FuRTBubjZXUEUKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
)

func TestShouldUnmarshalConfig(t *testing.T) {
	config := fmt.Sprintf(`
contexts:
- name: default
  context:
    url: http://localhost:8080
    api-token: foo
    ca-cert-base64-data: %s
- name: prod
  context:
    url: http://prod:8080
    service-token: foo
    api-path: /hidden/api
    admin-api-path: /admin/api
current-context: prod
`, selfSignedBase64Cert)

	cfg, err := unmarshalYAMLConfig([]byte(config))
	assert.NoError(t, err)
	assert.Len(t, cfg.Contexts, 2)
	assert.Equal(t, "prod", cfg.CurrentContext)
	assert.Equal(t, "http://localhost:8080", cfg.Contexts[0].Context.URL)
	assert.Equal(t, "foo", cfg.Contexts[0].Context.APIToken)
	assert.Empty(t, cfg.Contexts[0].Context.ServiceToken)
	assert.Equal(t, "/api", cfg.Contexts[0].Context.APIPath)
	assert.Equal(t, "/admin", cfg.Contexts[0].Context.AdminAPIPath)
	assert.Equal(t, selfSignedBase64Cert, cfg.Contexts[0].Context.CaCertBase64Data)
	assert.Equal(t, "http://prod:8080", cfg.Contexts[1].Context.URL)
	assert.Equal(t, "foo", cfg.Contexts[1].Context.ServiceToken)
	assert.Equal(t, "/hidden/api", cfg.Contexts[1].Context.APIPath)
	assert.Equal(t, "/admin/api", cfg.Contexts[1].Context.AdminAPIPath)
	assert.Empty(t, cfg.Contexts[1].Context.APIToken)
}

func TestShouldNotReadK8sSaTokenFromConfig(t *testing.T) {
	config := `
contexts:
- name: nightly
  context:
    url: http://nightly:8080
    k8s-sa-token: foobar
    api-path: /hidden/api
current-context: nightly
`

	cfg, err := unmarshalYAMLConfig([]byte(config))
	assert.NoError(t, err)
	assert.Len(t, cfg.Contexts, 1)
	assert.Equal(t, "nightly", cfg.CurrentContext)
	assert.Equal(t, "http://nightly:8080", cfg.Contexts[0].Context.URL)
	assert.Equal(t, "", cfg.Contexts[0].Context.K8sSAToken)
}

func TestShouldUnmarshalOldConfigFormat(t *testing.T) {
	config := `
url: http://localhost:8080
api-token: foo
`
	cfg, err := unmarshalYAMLConfig([]byte(config))
	assert.NoError(t, err)
	assert.Len(t, cfg.Contexts, 1)
	assert.Equal(t, "default", cfg.CurrentContext)
	assert.Equal(t, "http://localhost:8080", cfg.Contexts[0].Context.URL)
	assert.Equal(t, "foo", cfg.Contexts[0].Context.APIToken)
	assert.Empty(t, cfg.Contexts[0].Context.ServiceToken)
	assert.Equal(t, "/api", cfg.Contexts[0].Context.APIPath)
	assert.Equal(t, "/admin", cfg.Contexts[0].Context.AdminAPIPath)
}

func TestValidateValidStsContext(t *testing.T) {
	config := fmt.Sprintf(`
contexts:
- name: default
  context:
    url: http://localhost:8080
    api-token: foo
    ca-cert-base64-data: %s
current-context: default
`, selfSignedBase64Cert)
	c, err := unmarshalYAMLConfig([]byte(config))
	assert.NoError(t, err)
	assert.NoError(t, c.Contexts[0].Context.Validate(c.Contexts[0].Name))
}

func TestValidateStsContextWithMissingTokens(t *testing.T) {
	config := `
contexts:
- name: default
  context:
    url: http://localhost:8080
current-context: default
`
	c, err := unmarshalYAMLConfig([]byte(config))
	assert.NoError(t, err)
	assert.ErrorContains(t, c.Contexts[0].Context.Validate(c.Contexts[0].Name), "Failed to validate the 'default' context:\n* Missing field '{api-token | service-token | k8s-sa-token}'")
}

func TestValidateStsContextWithMissingURL(t *testing.T) {
	config := `
contexts:
- name: default
  context:
    api-token: foo
current-context: default
`
	c, err := unmarshalYAMLConfig([]byte(config))
	assert.NoError(t, err)
	assert.ErrorContains(t, c.Contexts[0].Context.Validate(c.Contexts[0].Name), "Failed to validate the 'default' context:\n* Missing field 'url'")
}

func TestValidateInvalidCertBase64Data(t *testing.T) {
	config := `
contexts:
- name: default
  context:
    url: http://localhost:8080
    api-token: foo
    ca-cert-base64-data: not-a-valid-base64
current-context: default
`
	c, err := unmarshalYAMLConfig([]byte(config))
	assert.NoError(t, err)
	assert.ErrorContains(t, c.Contexts[0].Context.Validate(c.Contexts[0].Name), "ca-cert-base64-data is not a valid base64 encoded string")
}

func TestValidateInvalidCertificate(t *testing.T) {
	config := `
contexts:
- name: default
  context:
    url: http://localhost:8080
    api-token: foo
    ca-cert-base64-data: bm90IGEgdmFsaWQgYmFzZTY0Cg==
current-context: default
`
	c, err := unmarshalYAMLConfig([]byte(config))
	assert.NoError(t, err)
	assert.ErrorContains(t, c.Contexts[0].Context.Validate(c.Contexts[0].Name), "ca-cert-base64-data is not a valid X509 certificate")
}

func TestValidateStsContextWithMalformedURL(t *testing.T) {
	config := `
contexts:
- name: default
  context:
    url: localhost:8080
    api-token: foo
current-context: default
`
	c, err := unmarshalYAMLConfig([]byte(config))
	assert.NoError(t, err)
	assert.ErrorContains(t, c.Contexts[0].Context.Validate(c.Contexts[0].Name), "Failed to validate the 'default' context:\n* URL localhost:8080 must start with \"https://\" or \"http://\"")
}

func TestValidateWithMultipleErrors(t *testing.T) {
	config := `
contexts:
- name: default
  context:
    url: localhost:8080
    api-token: foo
    service-token: bar
current-context: default
`
	c, err := unmarshalYAMLConfig([]byte(config))
	assert.NoError(t, err)
	assert.ErrorContains(t, c.Contexts[0].Context.Validate(c.Contexts[0].Name), `Failed to validate the 'default' context:
* URL localhost:8080 must start with "https://" or "http://"
* Can only specify one of {api-token | service-token | k8s-sa-token}`)
}

func TestMergeWithNoTokenOverride(t *testing.T) {
	c := &StsContext{
		URL: "http://localhost:8080",
	}
	f := &StsContext{
		URL:      "http://other:8080",
		APIToken: "foo",
	}

	n := c.Merge(f)
	assert.Equal(t, "http://localhost:8080", n.URL)
	assert.Equal(t, "foo", n.APIToken)
}

func TestMergeWithSameTokenOverride(t *testing.T) {
	c := &StsContext{
		URL:      "http://localhost:8080",
		APIToken: "bar",
	}
	f := &StsContext{
		URL:      "http://other:8080",
		APIToken: "foo",
	}

	n := c.Merge(f)
	assert.Equal(t, "http://localhost:8080", n.URL)
	assert.Equal(t, "bar", n.APIToken)
}

func TestMergeWithOtherTokenOverride(t *testing.T) {
	c := &StsContext{
		URL:          "http://localhost:8080",
		ServiceToken: "bar",
	}
	f := &StsContext{
		URL:      "http://other:8080",
		APIToken: "foo",
	}

	n := c.Merge(f)
	assert.Equal(t, "http://localhost:8080", n.URL)
	assert.Equal(t, "bar", n.ServiceToken)
	assert.Equal(t, "", n.APIToken)
}

type ValidateX509CertificateTest struct {
	name        string
	certData    []byte
	expectError bool
	errorMsg    string
}

func TestValidateX509Certificate(t *testing.T) {
	validPEMCert := []byte(`-----BEGIN CERTIFICATE-----
MIICJzCCAdGgAwIBAgIUV/aJj1fEcCgNU2FafY0uRLqy7mAwDQYJKoZIhvcNAQEL
BQAwKTEnMCUGA1UEAwwedmlsaWFrb3Yuc2FuZGJveC5zdGFja3N0YXRlLmlvMCAX
DTI1MDcyMTIwNDU1NloYDzIxMjUwNjI3MjA0NTU2WjApMScwJQYDVQQDDB52aWxp
YWtvdi5zYW5kYm94LnN0YWNrc3RhdGUuaW8wXDANBgkqhkiG9w0BAQEFAANLADBI
AkEAopQuOJfIkLCWJKT70hgbHJpUkEB+Za2p9uA1INRKM4Ar7dcV9muxNKOcJZ2s
Vt+b+YSKW8rtmxNPQXuE2D4teQIDAQABo4HOMIHLMB0GA1UdDgQWBBQU0OLVQG12
wMoEKHgqHmZyXSzZ3zAfBgNVHSMEGDAWgBQU0OLVQG12wMoEKHgqHmZyXSzZ3zAP
BgNVHRMBAf8EBTADAQH/MHgGA1UdEQRxMG+CHnZpbGlha292LnNhbmRib3guc3Rh
Y2tzdGF0ZS5pb4Ijb3RscC12aWxpYWtvdi5zYW5kYm94LnN0YWNrc3RhdGUuaW+C
KG90bHAtaHR0cC12aWxpYWtvdi5zYW5kYm94LnN0YWNrc3RhdGUuaW8wDQYJKoZI
hvcNAQELBQADQQAfYAVMeMRGlW+GZkzYOxdHiXX4AHHp9IloZPLmBG4LmvZC80hV
K4cRUEGHkRxgk0h0c9wD8NdVR3QnE0nn6WPE
-----END CERTIFICATE-----`)

	invalidPEMWrongType := []byte(`-----BEGIN TEST KEY-----
MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQg7S8j1SWx4gXGKVhR
Q0W6ixfaWXFOQ5Xk7p9sX2BxE3FoRANCAAQ5YK1G2P3+nRjwKwVCT/ixkNXwlPuK
rAHi2zCsHwKV+1gF7NqJEGbO6UBq0o4n9wGVoGkrRK5vHlL3HyFlxqSP
-----END TEST KEY-----`)

	invalidPEMData := []byte(`-----BEGIN CERTIFICATE-----
invalid-cert-data
-----END CERTIFICATE-----`)

	tests := []ValidateX509CertificateTest{
		{
			name:        "valid PEM certificate",
			certData:    validPEMCert,
			expectError: false,
		},
		{
			name:        "empty certificate data",
			certData:    []byte{},
			expectError: true,
			errorMsg:    "x509: malformed certificate",
		},
		{
			name:        "nil certificate data",
			certData:    nil,
			expectError: true,
			errorMsg:    "x509: malformed certificate",
		},
		{
			name:        "invalid PEM block type",
			certData:    invalidPEMWrongType,
			expectError: true,
			errorMsg:    "expected PEM block type CERTIFICATE, got TEST KEY",
		},
		{
			name:        "invalid PEM certificate data",
			certData:    invalidPEMData,
			expectError: true,
			errorMsg:    "x509: malformed certificate",
		},
		{
			name:        "plain text data",
			certData:    []byte("not a certificate"),
			expectError: true,
			errorMsg:    "x509: malformed certificate",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := validateX509Certificate(test.certData)

			if test.expectError {
				require.Error(t, err)
				assert.Contains(t, err.Error(), test.errorMsg)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
