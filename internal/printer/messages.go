package printer

import (
	"fmt"
)

type LoadingMsg int

const (
	AwaitingServer        LoadingMsg = iota
	NoTableDataDefaultMsg            = "No data to display."
)

func (l LoadingMsg) String() string {
	//nolint:gocritic
	switch l {
	case AwaitingServer:
		return "Awaiting server response.."
	}
	return "Loading..."
}

type MissingTableDataMsg interface {
	String() string
}

type NotFoundMsg struct {
	Types string
}

func (m NotFoundMsg) String() string {
	return fmt.Sprintf("No %s found.", m.Types)
}
