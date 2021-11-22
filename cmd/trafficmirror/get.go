package trafficmirror

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/config"
)

func GetCommand(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Retrieve a list of trafficmirror endpoints",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := &http.Client{}

			req, err := http.NewRequestWithContext(cmd.Context(), http.MethodGet, cfg.TrafficMirror.URL, nil)
			if err != nil {
				return err
			}

			req.SetBasicAuth(cfg.TrafficMirror.User, cfg.TrafficMirror.Password)

			resp, err := client.Do(req)
			if err != nil {
				return err
			}

			defer resp.Body.Close()

			if resp.StatusCode != 200 {
				cmd.Printf("%v\n", resp.Status)
				return errors.New("non-success error returned from trafficmirror")
			}

			data, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			cmd.Printf("%v\n", string(data))
			return nil
		},
	}

	return cmd
}
