package stackpack

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
)

const (
	Name      = "name"
	NameUsage = "StackPack name"
)

type DescribeArgs struct {
	Name string
}

func StackpackDescribeCommand(cli *di.Deps) *cobra.Command {
	args := &DescribeArgs{}
	cmd := &cobra.Command{
		Use:   "describe",
		Short: "Describe a StackPack",
		Long:  "Describe a StackPack.",
		RunE:  cli.CmdRunEWithApi(RunStackpackDescribeCommand(args)),
	}
	cmd.Flags().StringVar(&args.Name, Name, "", NameUsage)
	cmd.MarkFlagRequired(Name) //nolint:errcheck

	return cmd
}

func RunStackpackDescribeCommand(args *DescribeArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		stackPackList, err := fetchAllStackPacks(cli, api)
		if err != nil {
			return err
		}

		for _, stackpack := range stackPackList {
			if stackpack.Name == args.Name {
				if cli.IsJson() {
					cli.Printer.PrintJson(map[string]interface{}{
						"stackpack": stackpack,
					})
				} else {
					cli.Printer.PrintLn(fmt.Sprintf("StackPack %s v%s (latest available is v%s)", stackpack.DisplayName, stackpack.Version, stackpack.LatestVersion.Version))

					if len(stackpack.Categories) > 0 {
						cli.Printer.PrintLn(fmt.Sprintf("Categories: %s", strings.Join(stackpack.Categories, ", ")))
					}

					if len(stackpack.Integrations) > 0 {
						integrations := make([]string, 0)
						for _, i := range stackpack.Integrations {
							integrations = append(integrations, i.DisplayName)
						}
						cli.Printer.PrintLn(fmt.Sprintf("Provided integrations: %s", strings.Join(integrations, ",")))
					}
					if len(stackpack.Steps) > 0 {
						cli.Printer.PrintLn("\nInstance configuration parameters:")
						cli.Printer.Table(printer.TableData{
							Header:              []string{"name", "display name", "type"},
							Data:                formatParametersTable(stackpack),
							MissingTableDataMsg: printer.NotFoundMsg{Types: "instance parameters"},
						})
					}
					cli.Printer.PrintLn("\nInstalled instances:")
					cli.Printer.Table(printer.TableData{
						Header:              []string{"id", "status", "version", "last updated"},
						Data:                formatInstancesTable(stackpack),
						MissingTableDataMsg: printer.NotFoundMsg{Types: "installed StackPack instances"},
					})
				}
				return nil
			}
		}

		notFound := fmt.Errorf("StackPack '%s' not found.", args.Name)
		if cli.IsJson() {
			cli.Printer.PrintErrJson(notFound)
		} else {
			cli.Printer.PrintErr(notFound)
		}
		return nil
	}
}
