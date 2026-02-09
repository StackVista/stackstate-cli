package topology

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
)

type InspectArgs struct {
	ComponentType string
	Tags          []string
	Identifiers   []string
	Limit         int
}

func InspectCommand(cli *di.Deps) *cobra.Command {
	args := &InspectArgs{}

	cmd := &cobra.Command{
		Use:   "inspect",
		Short: "Inspect topology components",
		Long:  "Inspect topology components by type, tags, and identifiers. Displays component details and provides links to the UI.",
		Example: `# inspect components of a specific type
sts topology inspect --type "otel service instance"

# inspect with tag filtering
sts topology inspect --type "otel service instance" --tag "service.namespace:opentelemetry-demo-demo-dev"

# inspect with multiple tags (ANDed)
sts topology inspect --type "otel service instance" \
  --tag "service.namespace:opentelemetry-demo-demo-dev" \
  --tag "service.name:accountingservice"

# inspect with identifier filtering
sts topology inspect --type "otel service instance" --identifier "urn:opentelemetry:..."

# inspect with limit on number of results
sts topology inspect --type "otel service instance" --limit 10

# inspect and display as JSON (includes component links)
sts topology inspect --type "otel service instance" -o json`,
		RunE: cli.CmdRunEWithApi(func(cmd *cobra.Command, cli *di.Deps, api *stackstate_api.APIClient, serverInfo *stackstate_api.ServerInfo) common.CLIError {
			return RunInspectCommand(cmd, cli, api, serverInfo, args)
		}),
	}

	cmd.Flags().StringVar(&args.ComponentType, "type", "", "Component type (required)")
	cmd.MarkFlagRequired("type") //nolint:errcheck
	cmd.Flags().StringSliceVar(&args.Tags, "tag", []string{}, "Filter by tags in format 'tag-name:tag-value' (multiple allowed, ANDed together)")
	cmd.Flags().StringSliceVar(&args.Identifiers, "identifier", []string{}, "Filter by component identifiers (multiple allowed, ANDed together)")
	cmd.Flags().IntVar(&args.Limit, "limit", 0, "Maximum number of components to output (must be positive)")

	return cmd
}

func RunInspectCommand(
	_ *cobra.Command,
	cli *di.Deps,
	api *stackstate_api.APIClient,
	_ *stackstate_api.ServerInfo,
	args *InspectArgs,
) common.CLIError {
	if args.Limit < 0 {
		return common.NewExecutionError(fmt.Errorf("limit must be a positive number, got: %d", args.Limit))
	}

	query := buildSTQLQuery(args.ComponentType, args.Tags, args.Identifiers)

	metadata := stackstate_api.NewQueryMetadata(
		false,
		false,
		0,
		false,
		false,
		false,
		false,
		false,
		false,
		true,
	)

	request := stackstate_api.NewViewSnapshotRequest(
		"SnapshotRequest",
		query,
		"0.0.1",
		*metadata,
	)

	result, resp, err := api.SnapshotApi.QuerySnapshot(cli.Context).
		ViewSnapshotRequest(*request).
		Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	// The `CmdRunEWithApi` ensures the current context is loaded, so it should be safe to access context properties
	components, parseErr := parseSnapshotResponse(result, cli.CurrentContext.URL)
	if parseErr != nil {
		// Check if the error is a typed error from the response by examining the _type discriminator
		if typedErr := handleSnapshotError(result.ViewSnapshotResponse, resp); typedErr != nil {
			return typedErr
		}
		return common.NewExecutionError(parseErr)
	}

	// Apply limit if specified
	if args.Limit > 0 && len(components) > args.Limit {
		components = components[:args.Limit]
	}

	if cli.IsJson() {
		cli.Printer.PrintJson(map[string]interface{}{
			"components": components,
		})
		return nil
	} else {
		printTableOutput(cli, components)
	}

	return nil
}

func buildSTQLQuery(componentType string, tags []string, identifiers []string) string {
	query := fmt.Sprintf(`type = "%s"`, componentType)

	for _, tag := range tags {
		query += fmt.Sprintf(` AND label = "%s"`, tag)
	}

	// Add identifier filters (ANDed with IN clause)
	if len(identifiers) > 0 {
		quotedIds := make([]string, len(identifiers))
		for i, id := range identifiers {
			quotedIds[i] = fmt.Sprintf(`"%s"`, id)
		}
		query += fmt.Sprintf(` AND identifier IN (%s)`, strings.Join(quotedIds, ", "))
	}

	return query
}

type Component struct {
	ID          int64                  `json:"id"`
	Name        string                 `json:"name"`
	Type        string                 `json:"type"`
	Identifiers []string               `json:"identifiers"`
	Tags        []string               `json:"tags"`
	Properties  map[string]interface{} `json:"properties"`
	Layer       map[string]interface{} `json:"layer"`
	Domain      map[string]interface{} `json:"domain"`
	Environment map[string]interface{} `json:"environment,omitempty"`
	Link        string                 `json:"link"`
}

type ComponentMetadata struct {
	ComponentTypes map[int64]string
	Layers         map[int64]string
	Domains        map[int64]string
	Environments   map[int64]string
}

// handleSnapshotError checks if the response is a typed error by examining the _type discriminator
func handleSnapshotError(respMap map[string]interface{}, resp *http.Response) common.CLIError {
	if respMap == nil {
		return nil
	}

	responseType, ok := respMap["_type"].(string)
	if !ok {
		return nil
	}

	// Note: the `View` prefix on the error types comes from the existing (Scala) DTO types
	switch responseType {
	case "ViewSnapshotFetchTimeout":
		if usedTimeoutSeconds, ok := respMap["usedTimeoutSeconds"].(float64); ok {
			message := fmt.Sprintf(
				"Query timed out after %d seconds. Please refine your query to be more specific.",
				int64(usedTimeoutSeconds),
			)
			return common.NewResponseError(fmt.Errorf("%s", message), resp)
		}

	case "ViewSnapshotTooManyActiveQueries":
		message := "Too many active queries are running. Please try again later."
		return common.NewResponseError(fmt.Errorf("%s", message), resp)

	case "ViewSnapshotTopologySizeOverflow":
		if maxSize, ok := respMap["maxSize"].(float64); ok {
			message := fmt.Sprintf(
				"Query result exceeded maximum size of %d components. Please refine your query to be more specific.",
				int64(maxSize),
			)
			return common.NewResponseError(fmt.Errorf("%s", message), resp)
		}

	case "ViewSnapshotDataUnavailable":
		if availableSinceEpochMs, ok := respMap["availableSinceEpochMs"].(float64); ok {
			message := fmt.Sprintf(
				"Requested data is not available. Data is available from %d (epoch ms) onwards.",
				int64(availableSinceEpochMs),
			)
			return common.NewResponseError(fmt.Errorf("%s", message), resp)
		}
	}

	return nil
}

func parseSnapshotResponse(
	result *stackstate_api.QuerySnapshotResult,
	baseURL string,
) ([]Component, error) {
	// SnapshotResponse is already typed as map[string]interface{} from the generated API
	respMap := result.ViewSnapshotResponse
	if respMap == nil {
		return nil, fmt.Errorf("response data is nil")
	}

	// Check if this is a Snapshot (success) or an error type
	respType, ok := respMap["_type"].(string)
	if !ok {
		return nil, fmt.Errorf("response has no _type discriminator")
	}

	// Note: the `View` prefix on response type comes from the existing (Scala) DTO types
	// If it's not a ViewSnapshot, it's an error type - return empty components and let error handler deal with it
	if respType != "ViewSnapshot" {
		return nil, fmt.Errorf("response is an error type: %s", respType)
	}

	metadata := parseMetadata(respMap)

	// Parse components from the already-unmarshalled components field
	var components []Component
	if componentsSlice, ok := respMap["components"].([]interface{}); ok {
		for _, comp := range componentsSlice {
			if compMap, ok := comp.(map[string]interface{}); ok {
				components = append(components, parseComponentFromMap(compMap, metadata, baseURL))
			}
		}
	}

	return components, nil
}

// metadataFieldMapping maps response JSON field names to the corresponding
// ComponentMetadata struct field setter. Used by parseMetadata to extract
// all metadata categories in a single loop.
var metadataFieldMapping = []struct {
	field  string
	setter func(*ComponentMetadata) *map[int64]string
}{
	{"componentTypes", func(m *ComponentMetadata) *map[int64]string { return &m.ComponentTypes }},
	{"layers", func(m *ComponentMetadata) *map[int64]string { return &m.Layers }},
	{"domains", func(m *ComponentMetadata) *map[int64]string { return &m.Domains }},
	{"environments", func(m *ComponentMetadata) *map[int64]string { return &m.Environments }},
}

// parseMetadata extracts component type, layer, domain, and environment metadata
// from the opaque Snapshot response using a table-driven approach.
func parseMetadata(respMap map[string]interface{}) ComponentMetadata {
	metadata := ComponentMetadata{
		ComponentTypes: make(map[int64]string),
		Layers:         make(map[int64]string),
		Domains:        make(map[int64]string),
		Environments:   make(map[int64]string),
	}

	metadataMap, ok := respMap["metadata"].(map[string]interface{})
	if !ok {
		return metadata
	}

	for _, mapping := range metadataFieldMapping {
		if fieldValue, ok := metadataMap[mapping.field]; ok {
			*mapping.setter(&metadata) = parseMetadataField(fieldValue)
		}
	}

	return metadata
}

// parseMetadataField extracts id/name pairs from a metadata field.
// Each item in the slice should have "id" and "name" fields.
func parseMetadataField(metadataValue interface{}) map[int64]string {
	result := make(map[int64]string)

	if metadataValue == nil {
		return result
	}

	// The JSON decoder produces []interface{} for arrays
	items, ok := metadataValue.([]interface{})
	if !ok {
		return result
	}

	for _, item := range items {
		if itemMap, ok := item.(map[string]interface{}); ok {
			id, idOk := itemMap["id"].(float64)
			name, nameOk := itemMap["name"].(string)
			if idOk && nameOk {
				result[int64(id)] = name
			}
		}
	}

	return result
}

func parseComponentFromMap(compMap map[string]interface{}, metadata ComponentMetadata, baseURL string) Component {
	comp := Component{
		Identifiers: []string{},
		Tags:        []string{},
		Properties:  make(map[string]interface{}),
	}

	// Parse basic fields
	if id, ok := compMap["id"].(float64); ok {
		comp.ID = int64(id)
	}
	if name, ok := compMap["name"].(string); ok {
		comp.Name = name
	}

	// Parse type (first id and then lookup from component type metadata)
	if typeID, ok := compMap["type"].(float64); ok {
		if typeName, found := metadata.ComponentTypes[int64(typeID)]; found {
			comp.Type = typeName
		} else {
			comp.Type = fmt.Sprintf("Unknown (%d)", int64(typeID))
		}
	}

	// Parse identifiers
	if identifiersRaw, ok := compMap["identifiers"].([]interface{}); ok {
		for _, idRaw := range identifiersRaw {
			if id, ok := idRaw.(string); ok {
				comp.Identifiers = append(comp.Identifiers, id)
			}
		}
	}

	// Parse tags
	if tagsRaw, ok := compMap["tags"].([]interface{}); ok {
		for _, tagRaw := range tagsRaw {
			if tag, ok := tagRaw.(string); ok {
				comp.Tags = append(comp.Tags, tag)
			}
		}
	}

	// Parse properties
	if propertiesRaw, ok := compMap["properties"].(map[string]interface{}); ok {
		comp.Properties = propertiesRaw
	}

	// Parse layer, domain, and environment references
	comp.Layer = parseComponentReference(compMap, "layer", metadata.Layers)
	comp.Domain = parseComponentReference(compMap, "domain", metadata.Domains)
	comp.Environment = parseComponentReference(compMap, "environment", metadata.Environments)

	// Build link
	if len(comp.Identifiers) > 0 {
		comp.Link = buildComponentLink(baseURL, comp.Identifiers[0])
	}

	return comp
}

// parseComponentReference extracts a reference field (layer, domain, or environment)
// from a component and looks up its name in the provided metadata map.
// Returns a map with "id" and "name" keys, or nil if the field is not present.
func parseComponentReference(compMap map[string]interface{}, fieldName string, metadataMap map[int64]string) map[string]interface{} {
	if refID, ok := compMap[fieldName].(float64); ok {
		refIDInt := int64(refID)
		refName := "Unknown"
		if name, found := metadataMap[refIDInt]; found {
			refName = name
		}
		return map[string]interface{}{
			"id":   refIDInt,
			"name": refName,
		}
	}
	return nil
}

func buildComponentLink(baseURL, identifier string) string {
	encodedIdentifier := url.PathEscape(identifier)
	return fmt.Sprintf("%s/#/components/%s", baseURL, encodedIdentifier)
}

func printTableOutput(
	cli *di.Deps,
	components []Component,
) {
	var tableData [][]interface{}
	for _, comp := range components {
		identifiersStr := strings.Join(comp.Identifiers, ", ")
		tableData = append(tableData, []interface{}{
			comp.Name,
			comp.Type,
			identifiersStr,
		})
	}

	cli.Printer.Table(printer.TableData{
		Header:              []string{"Name", "Type", "Identifiers"},
		Data:                tableData,
		MissingTableDataMsg: printer.NotFoundMsg{Types: "components"},
	})
}
