// This fixture runs inside the candidate image and invokes the real CLI.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	if err := check(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func check() error {
	// Expectations are literals, independent of the fixture's timezone database.
	cases := []struct {
		zone           string
		epoch          int64
		timestamp, day string
	}{
		{"UTC", 0, "1970-01-01 00:00:00 UTC", "1970-01-01"},
		{"America/New_York", 0, "1969-12-31 19:00:00 EST", "1969-12-31"},
		{"America/New_York", 1593561600000, "2020-06-30 20:00:00 EDT", "2020-06-30"},
		{"Asia/Kathmandu", 1593561600000, "2020-07-01 05:45:00 +0545", "2020-07-01"},
		// DLA-4569-1: Vancouver must not fall back to UTC-08 in November 2026.
		{"America/Vancouver", 1793604600000, "2026-11-02 00:30:00 MST", "2026-11-02"},
		{"America/Vancouver", 1798788600000, "2027-01-01 00:30:00 MST", "2027-01-01"},
		// Historical winter timestamps must still use the former UTC-08 rule.
		{"America/Vancouver", 0, "1969-12-31 16:00:00 PST", "1969-12-31"},
	}
	for _, tc := range cases {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			switch strings.TrimPrefix(r.URL.Path, "/api") {
			case "/server/info":
				fmt.Fprint(w, `{"version":{"major":6,"minor":0,"patch":0,"diff":"","commit":"","isDev":false},"deploymentMode":"SelfHosted","applicationDomains":[]}`)
			case "/agents":
				fmt.Fprintf(w, `{"agents":[{"agentId":"tz-probe","lease":"Active","registeredEpochMs":%d,"leaseUntilEpochMs":%d,"nodeBudgetCount":1}]}`, tc.epoch, tc.epoch)
			case "/subscription":
				fmt.Fprintf(w, `{"_type":"LicensedSubscription","subscription":{"tenant":"fixture","plan":"test","expiryTimestampMs":%d}}`, tc.epoch)
			case "/security/tokens":
				fmt.Fprintf(w, `[{"id":1,"name":"fixture","expiration":%d,"roles":[]}]`, tc.epoch)
			default:
				http.NotFound(w, r)
			}
		}))
		for _, command := range []struct {
			args   []string
			column int
			want   string
		}{
			{[]string{"agent", "list"}, 2, tc.timestamp},
			{[]string{"agent", "list"}, 3, tc.timestamp},
			{[]string{"license", "show"}, 2, tc.day},
			{[]string{"service-token", "list"}, 2, tc.day},
		} {
			out, err := run(tc.zone, server.URL, command.args...)
			if err != nil {
				server.Close()
				return err
			}
			// Reassemble a wrapped table column at the CLI's default width.
			var column strings.Builder
			for _, line := range strings.Split(out, "\n") {
				cells := strings.Split(line, "|")
				if len(cells) > command.column {
					column.WriteString(strings.Join(strings.Fields(cells[command.column]), ""))
				}
			}
			if !strings.Contains(column.String(), strings.ReplaceAll(command.want, " ", "")) {
				server.Close()
				return fmt.Errorf("TZ=%s %v: expected %q in column %d\n%s", tc.zone, command.args, command.want, command.column, out)
			}
			fmt.Printf("PASS TZ=%s %v column=%d: %s\n", tc.zone, command.args, command.column, command.want)
		}
		out, err := run(tc.zone, server.URL, "agent", "list", "-o", "json")
		server.Close()
		if err != nil {
			return err
		}
		var result struct {
			Agents []struct {
				Registered int64 `json:"registeredEpochMs"`
				LeaseUntil int64 `json:"leaseUntilEpochMs"`
			} `json:"agents"`
		}
		if err := json.Unmarshal([]byte(out), &result); err != nil {
			return fmt.Errorf("agent JSON: %w\n%s", err, out)
		}
		if len(result.Agents) != 1 || result.Agents[0].Registered != tc.epoch || result.Agents[0].LeaseUntil != tc.epoch {
			return fmt.Errorf("TZ=%s: JSON timestamps changed: %s", tc.zone, out)
		}
		fmt.Printf("PASS TZ=%s JSON epoch milliseconds: %d\n", tc.zone, tc.epoch)
	}
	return nil
}

func run(zone, url string, args ...string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	args = append(args, "--url", url, "--api-token", "fixture")
	cmd := exec.CommandContext(ctx, "/usr/bin/sts", args...)
	cmd.Env = append(os.Environ(), "TZ="+zone, "NO_COLOR=1", "TERM=dumb", "XDG_CONFIG_HOME=/tmp/cli-timezone-test")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%v: %w\n%s", args[:2], err, out)
	}
	return string(out), nil
}
