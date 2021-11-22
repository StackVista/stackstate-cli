package trafficmirror

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/config"
)

func DeleteCommand(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete [<url to be removed>]",
		Args:  cobra.ExactArgs(1),
		Short: "Remove an endpoint from the list of trafficmirror endpoints",
		RunE: func(cmd *cobra.Command, args []string) error {
			data := url.Values{}
			data.Set("url", args[0])

			url := cfg.TrafficMirror.URL + "?" + data.Encode()

			client := &http.Client{}

			req, err := http.NewRequestWithContext(cmd.Context(), http.MethodDelete, url, nil)
			if err != nil {
				return err
			}

			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
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

			respData, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			cmd.Printf("ok%v\n", string(respData))
			return nil
		},
	}
	return cmd
}
