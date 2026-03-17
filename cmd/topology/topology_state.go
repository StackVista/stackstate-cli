package topology

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	stscobra "github.com/stackvista/stackstate-cli/internal/cobra"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
)

type StateArgs struct {
	ComponentType string
	Tags          []string
	Identifiers   []string
	Limit         int
	STQL          string
}

func StateCommand(cli *di.Deps) *cobra.Command {
	args := &StateArgs{}

	cmd := &cobra.Command{
		Use:   "state",
		Short: "Show the health state of topology components",
		Long:  "Show the health state of topology components by type, tags, and identifiers. Displays the health state for each matching component.",
		Example: `# show state of components of a specific type
sts topology state --type "otel service instance"

# show state with tag filtering
sts topology state --type "otel service instance" --tag "service.namespace:opentelemetry-demo-demo-dev"

# show state with multiple tags (ANDed)
sts topology state --type "otel service instance" \
  --tag "service.namespace:opentelemetry-demo-demo-dev" \
  --tag "service.name:accountingservice"

# show state with identifier filtering
sts topology state --type "otel service instance" --identifier "urn:opentelemetry:..."

# show state with limit on number of results
sts topology state --type "otel service instance" --limit 10

# show state and display as JSON
sts topology state --type "otel service instance" -o json

# show state using a custom STQL query
sts topology state --stql 'type = "otel service instance" AND healthState = "CRITICAL"'`,
		RunE: cli.CmdRunEWithApi(func(cmd *cobra.Command, cli *di.Deps, api *stackstate_api.APIClient, serverInfo *stackstate_api.ServerInfo) common.CLIError {
			return RunStateCommand(cmd, cli, api, serverInfo, args)
		}),
	}

	cmd.Flags().StringVar(&args.ComponentType, "type", "", "Component type")
	cmd.Flags().StringSliceVar(&args.Tags, "tag", []string{}, "Filter by tags in format 'tag-name:tag-value' (multiple allowed, ANDed together)")
	cmd.Flags().StringSliceVar(&args.Identifiers, "identifier", []string{}, "Filter by component identifiers (multiple allowed, ANDed together)")
	cmd.Flags().IntVar(&args.Limit, "limit", 0, "Maximum number of components to output (must be positive)")
	cmd.Flags().StringVar(&args.STQL, "stql", "", "STQL query to select components (mutually exclusive with --type, --tag, --identifier)")
	stscobra.MarkMutexFlags(cmd, []string{"type", "stql"}, "query", true)

	return cmd
}

func RunStateCommand(
	_ *cobra.Command,
	cli *di.Deps,
	api *stackstate_api.APIClient,
	_ *stackstate_api.ServerInfo,
	args *StateArgs,
) common.CLIError {
	if args.Limit < 0 {
		return common.NewExecutionError(fmt.Errorf("limit must be a positive number, got: %d", args.Limit))
	}

	var query string
	if args.STQL != "" {
		query = args.STQL
	} else {
		query = buildSTQLQuery(args.ComponentType, args.Tags, args.Identifiers)
	}

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

	componentStates, parseErr := parseStateResponse(result)
	if parseErr != nil {
		if typedErr := handleSnapshotError(result.ViewSnapshotResponse, resp); typedErr != nil {
			return typedErr
		}
		return common.NewExecutionError(parseErr)
	}

	// Apply limit if specified
	if args.Limit > 0 && len(componentStates) > args.Limit {
		componentStates = componentStates[:args.Limit]
	}

	if cli.IsJson() {
		cli.Printer.PrintJson(map[string]interface{}{
			"components": componentStates,
		})
		return nil
	} else {
		printStateTableOutput(cli, componentStates)
	}

	return nil
}

// ComponentState holds the component information along with its health state.
type ComponentState struct {
	ID          int64    `json:"id"`
	Name        string   `json:"name"`
	Type        string   `json:"type"`
	Identifiers []string `json:"identifiers"`
	HealthState string   `json:"healthState"`
}

func parseStateResponse(result *stackstate_api.QuerySnapshotResult) ([]ComponentState, error) {
	respMap := result.ViewSnapshotResponse
	if respMap == nil {
		return nil, fmt.Errorf("response data is nil")
	}

	respType, ok := respMap["_type"].(string)
	if !ok {
		return nil, fmt.Errorf("response has no _type discriminator")
	}

	if respType != "ViewSnapshot" {
		return nil, fmt.Errorf("response is an error type: %s", respType)
	}

	metadata := parseMetadata(respMap)

	var componentStates []ComponentState
	if componentsSlice, ok := respMap["components"].([]interface{}); ok {
		for _, comp := range componentsSlice {
			if compMap, ok := comp.(map[string]interface{}); ok {
				componentStates = append(componentStates, parseComponentStateFromMap(compMap, metadata))
			}
		}
	}

	return componentStates, nil
}

func parseComponentStateFromMap(compMap map[string]interface{}, metadata ComponentMetadata) ComponentState {
	cs := ComponentState{
		Identifiers: []string{},
		HealthState: "UNKNOWN",
	}

	// Parse basic fields
	if id, ok := compMap["id"].(float64); ok {
		cs.ID = int64(id)
	}
	if name, ok := compMap["name"].(string); ok {
		cs.Name = name
	}

	// Parse type
	if typeID, ok := compMap["type"].(float64); ok {
		if typeName, found := metadata.ComponentTypes[int64(typeID)]; found {
			cs.Type = typeName
		} else {
			cs.Type = fmt.Sprintf("Unknown (%d)", int64(typeID))
		}
	}

	// Parse identifiers
	if identifiersRaw, ok := compMap["identifiers"].([]interface{}); ok {
		for _, idRaw := range identifiersRaw {
			if id, ok := idRaw.(string); ok {
				cs.Identifiers = append(cs.Identifiers, id)
			}
		}
	}

	// Parse health state from state.healthState
	if stateMap, ok := compMap["state"].(map[string]interface{}); ok {
		if healthState, ok := stateMap["healthState"].(string); ok {
			cs.HealthState = healthState
		}
	}

	return cs
}

func printStateTableOutput(cli *di.Deps, componentStates []ComponentState) {
	var tableData [][]interface{}
	for _, cs := range componentStates {
		identifiersStr := strings.Join(cs.Identifiers, ", ")
		tableData = append(tableData, []interface{}{
			cs.Name,
			cs.Type,
			cs.HealthState,
			identifiersStr,
		})
	}

	cli.Printer.Table(printer.TableData{
		Header:              []string{"Name", "Type", "Health State", "Identifiers"},
		Data:                tableData,
		MissingTableDataMsg: printer.NotFoundMsg{Types: "components"},
	})
}
