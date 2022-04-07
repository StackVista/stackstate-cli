package common

type LoadingMsg int

const (
	AwaitingServer LoadingMsg = iota
)

func (l LoadingMsg) String() string {
	//nolint:gocritic
	switch l {
	case AwaitingServer:
		return "Awaiting server response.."
	}
	return "Loading..."
}
