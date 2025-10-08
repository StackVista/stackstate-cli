package otelmapping

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type StatusArgs struct {
	Identifier string
}

func FormatOtelMappingStatusTable(otelmappings []stackstate_api.OtelMappingStatusItem) printer.TableData {
	data := make([][]interface{}, len(otelmappings))

	for i, otelmapping := range otelmappings {
		identifier := "-"
		if otelmapping.HasIdentifier() {
			identifier = otelmapping.GetIdentifier()
		}
		data[i] = []interface{}{
			otelmapping.Name,
			identifier,
			otelmapping.ComponentCount,
			otelmapping.RelationCount,
		}
	}
	return printer.TableData{
		Header:              []string{"Name", "Identifier", "Components", "Relations"},
		Data:                data,
		MissingTableDataMsg: printer.NotFoundMsg{Types: "otel mappings"},
	}
}

func FormatOtelMappingTable(otelmappings []stackstate_api.OtelMappingItem) printer.TableData {
	data := make([][]interface{}, len(otelmappings))

	for i, otelmapping := range otelmappings {
		identifier := "-"
		if otelmapping.HasIdentifier() {
			identifier = otelmapping.GetIdentifier()
		}
		data[i] = []interface{}{
			otelmapping.Name,
			identifier,
		}
	}
	return printer.TableData{
		Header:              []string{"Name", "Identifier"},
		Data:                data,
		MissingTableDataMsg: printer.NotFoundMsg{Types: "otel mappings"},
	}
}

type StatusFetcher func(cli *di.Deps, api *stackstate_api.APIClient, identifier string) (*stackstate_api.OtelMappingStatus, *http.Response, error)

func FetchComponentStatus(cli *di.Deps, api *stackstate_api.APIClient, identifier string) (*stackstate_api.OtelMappingStatus, *http.Response, error) {
	return api.OtelMappingApi.GetOtelComponentMappingStatus(cli.Context, identifier).Execute()
}

func FetchRelationStatus(cli *di.Deps, api *stackstate_api.APIClient, identifier string) (*stackstate_api.OtelMappingStatus, *http.Response, error) {
	return api.OtelMappingApi.GetOtelRelationMappingStatus(cli.Context, identifier).Execute()
}

func RunStatus(args *StatusArgs, mappingType string, fetch StatusFetcher) di.CmdWithApiFn {
	return func(cmd *cobra.Command, cli *di.Deps, api *stackstate_api.APIClient, serverInfo *stackstate_api.ServerInfo) common.CLIError {
		mappingStatus, resp, err := fetch(cli, api, args.Identifier)
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		jsonKey := fmt.Sprintf("otel-%s-mapping", mappingType)
		title := cases.Title(language.English).String(mappingType)
		mappingTitle := fmt.Sprintf("Otel %s Mapping:", title)
		metricsTitle := fmt.Sprintf("Otel %s Mapping Metrics:", title)
		errorsTitle := fmt.Sprintf("Otel %s Mapping Errors:", title)
		errorsNotFoundMsg := fmt.Sprintf("otel %s mapping errors", mappingType)

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				jsonKey: mappingStatus,
			})
		} else {
			cli.Printer.PrintLn("\n")
			cli.Printer.PrintLn(mappingTitle)
			cli.Printer.Table(FormatOtelMappingStatusTable([]stackstate_api.OtelMappingStatusItem{
				mappingStatus.Item,
			}))

			if mappingStatus.HasMetrics() {
				cli.Printer.PrintLn("\n")
				cli.Printer.PrintLn(metricsTitle)
				size := mappingStatus.Metrics.BucketSizeSeconds
				cli.Printer.Table(printer.TableData{
					Header: []string{"Metric", fmt.Sprintf("%ds ago", size), fmt.Sprintf("%d-%ds ago", size, 2*size), fmt.Sprintf("%d-%ds ago", 2*size, 3*size)}, //nolint:mnd
					Data: [][]interface{}{
						printer.MetricBucketToRow("latency seconds", mappingStatus.Metrics.LatencySeconds),
					},
					MissingTableDataMsg: printer.NotFoundMsg{Types: "metrics"},
				})
			}

			data := make([][]interface{}, len(mappingStatus.ErrorDetails))
			for i, error := range mappingStatus.ErrorDetails {
				id := "-"
				if error.HasIssueId() {
					id = *error.IssueId
				}
				data[i] = []interface{}{id, error.Level, error.Message}
			}

			cli.Printer.PrintLn("\n")
			cli.Printer.PrintLn(errorsTitle)
			cli.Printer.Table(printer.TableData{
				Header:              []string{"Issue Id", "Level", "Message"},
				Data:                data,
				MissingTableDataMsg: printer.NotFoundMsg{Types: errorsNotFoundMsg},
			})
		}

		return nil
	}
}
