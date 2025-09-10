package config

import (
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"
	"strings"

	"github.com/mcuadros/go-defaults"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/util"
)

type Config struct {
	Contexts       []*NamedContext `yaml:"contexts" json:"contexts"`
	CurrentContext string          `yaml:"current-context" json:"current-context"`
}

type NamedContext struct {
	Name    string      `yaml:"name" json:"name"`
	Context *StsContext `yaml:"context" json:"context"`
}

type StsContext struct {
	URL              string `yaml:"url" json:"url"`
	APIToken         string `yaml:"api-token,omitempty" json:"api-token,omitempty"`
	ServiceToken     string `yaml:"service-token,omitempty" json:"service-token,omitempty"`
	K8sSAToken       string `yaml:"-" json:"-"` // This should only be passed from command line or env variables
	K8sSATokenPath   string `yaml:"-" json:"-"` // This should only be passed from command line or env variables
	APIPath          string `yaml:"api-path" default:"/api" json:"api-path"`
	AdminAPIPath     string `yaml:"admin-api-path" default:"/admin" json:"admin-api-path"`
	SkipSSL          bool   `yaml:"skip-ssl" default:"false" json:"skip-ssl"`
	CaCertPath       string `yaml:"-" json:"-"` // This should only be passed from command line
	CaCertBase64Data string `yaml:"ca-cert-base64-data,omitempty" json:"ca-cert-base64-data,omitempty"`
}

func EmptyConfig() *Config {
	return &Config{
		Contexts: []*NamedContext{},
	}
}

func (c *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	m := map[string]interface{}{}
	type cfg Config

	if err := unmarshal(&m); err != nil {
		return err
	}

	// Check whether this is in the new format
	if _, ok := m["contexts"]; ok {
		if err := unmarshal((*cfg)(c)); err != nil {
			return err
		}
	} else {
		// Fallback to the old format
		ctx := &StsContext{}
		if err := unmarshal(ctx); err != nil {
			return err
		}
		c.Contexts = []*NamedContext{
			{
				Name:    "default",
				Context: ctx,
			},
		}
		c.CurrentContext = "default"
	}

	return nil
}

// GetContext gets the context with the given name
func (c *Config) GetContext(name string) (*NamedContext, common.CLIError) {
	for _, context := range c.Contexts {
		if context.Name == name {
			return context, nil
		}
	}

	return nil, common.NewNotFoundError(fmt.Errorf("context with name '%s' not found", name))
}

// UnmarshalYAML unmarshals the StsContext YAML part into a struct, ensuring that any defaults are set.
func (c *StsContext) UnmarshalYAML(unmarshal func(interface{}) error) error {
	defaults.SetDefaults(c)

	type cfg StsContext

	if err := unmarshal((*cfg)(c)); err != nil {
		return err
	}

	return nil
}

// Merges the StsContext with a fallback object.
func (c *StsContext) Merge(fallback *StsContext) *StsContext {
	newCtx := &StsContext{
		URL:              util.DefaultIfEmpty(c.URL, fallback.URL),
		APIPath:          util.DefaultIfEmpty(util.DefaultIfEmpty(c.APIPath, fallback.APIPath), "/api"),
		AdminAPIPath:     util.DefaultIfEmpty(util.DefaultIfEmpty(c.AdminAPIPath, fallback.AdminAPIPath), "/admin"),
		K8sSATokenPath:   util.DefaultIfEmpty(c.K8sSATokenPath, fallback.K8sSATokenPath),
		SkipSSL:          c.SkipSSL || fallback.SkipSSL,
		CaCertBase64Data: util.DefaultIfEmpty(c.CaCertBase64Data, fallback.CaCertBase64Data),
		CaCertPath:       util.DefaultIfEmpty(c.CaCertPath, fallback.CaCertPath),
	}

	if !c.HasAuthenticationTokenSet() {
		newCtx.APIToken = fallback.APIToken
		newCtx.ServiceToken = fallback.ServiceToken
		newCtx.K8sSAToken = fallback.K8sSAToken
	} else {
		newCtx.APIToken = c.APIToken
		newCtx.ServiceToken = c.ServiceToken
		newCtx.K8sSAToken = c.K8sSAToken
	}
	return newCtx
}

func (c *StsContext) HasAuthenticationTokenSet() bool {
	return len(util.RemoveEmpty([]string{c.APIToken, c.ServiceToken, c.K8sSAToken})) > 0
}

func (c *StsContext) HasCaCertificateSet() bool {
	return c.CaCertBase64Data != "" || c.CaCertPath != ""
}

func (c *StsContext) HasCaCertificateFromFileSet() bool {
	return c.CaCertPath != ""
}

func (c *StsContext) HasCaCertificateFromArgSet() bool {
	return c.CaCertBase64Data != ""
}

func (c *StsContext) Validate(contextName string) common.CLIError {
	errors := []error{}

	switch {
	case c.URL == "":
		errors = append(errors, MissingFieldError{FieldName: "url"})
	case !strings.HasPrefix(c.URL, "http://") && !strings.HasPrefix(c.URL, "https://"):
		errors = append(errors, fmt.Errorf("URL %s must start with \"https://\" or \"http://\"", c.URL))
	}

	if c.APIToken == "" && c.ServiceToken == "" && c.K8sSAToken == "" {
		errors = append(errors, MissingFieldError{FieldName: "{api-token | service-token | k8s-sa-token}"})
	}

	authenticationTokens := util.RemoveEmpty([]string{c.APIToken, c.ServiceToken, c.K8sSAToken})
	if len(authenticationTokens) > 1 {
		errors = append(errors, fmt.Errorf("can only specify one of {api-token | service-token | k8s-sa-token}"))
	}

	if c.HasCaCertificateFromArgSet() {
		caCertData, err := base64.StdEncoding.DecodeString(c.CaCertBase64Data)
		if err != nil {
			return common.NewAPIClientCreateError(fmt.Sprintf("%s is not a valid base64 encoded string", common.CaCertBase64DataFlag))
		}
		if err := validateX509Certificate(caCertData); err != nil {
			return common.NewAPIClientCreateError(fmt.Sprintf("%s is not a valid X509 certificate: %v", common.CaCertBase64DataFlag, err))
		}
	}

	if c.HasCaCertificateFromFileSet() {
		caCertData, serr := os.ReadFile(c.CaCertPath)
		if serr != nil {
			return common.NewReadFileError(serr, c.CaCertPath)
		}
		if err := validateX509Certificate(caCertData); err != nil {
			return common.NewAPIClientCreateError(fmt.Sprintf("%s is not a valid X509 certificate: %v", common.CaCertPathFlag, err))
		}
	}

	if len(errors) > 0 {
		return ValidateContextError{
			ContextName:      contextName,
			ValidationErrors: errors,
		}
	}

	return nil
}

func validateX509Certificate(caCertData []byte) error {
	block, _ := pem.Decode(caCertData)
	if block != nil {
		if block.Type != "CERTIFICATE" {
			return fmt.Errorf("expected PEM block type CERTIFICATE, got %s", block.Type)
		}
		caCertData = block.Bytes
	}
	if _, err := x509.ParseCertificate(caCertData); err != nil {
		return err
	}
	return nil
}
