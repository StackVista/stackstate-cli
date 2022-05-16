package time_flags

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
)

const (
	DateFormat = "2006-01-02"
)

func GetDateFlag(cmd *cobra.Command, flagName string) (*time.Time, common.CLIError) {
	strDate, err := cmd.Flags().GetString(flagName)
	if err != nil {
		return nil, common.NewCLIArgParseError(err)
	}

	if strDate == "" {
		return nil, nil
	}

	date, parseErr := ParseDate(strDate)
	if parseErr != nil {
		return nil, common.NewCLIArgParseError(parseErr)
	}

	return date, nil
}

func ParseDate(strDate string) (*time.Time, error) {
	date, parseErr := time.Parse(DateFormat, strDate)
	if parseErr != nil {
		return nil, newInvalidDateFormatError(strDate)
	}

	return &date, nil
}

func newInvalidDateFormatError(strDate string) error {
	return fmt.Errorf("invalid date format \"%s\". Expected an ISO8601 date-only string (e.g. 2016-09-13)", strDate)
}
