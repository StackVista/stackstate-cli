package common

type Output int64

const (
	TextOutput Output = iota
	JSONOutput
	YAMLOutput

	DefaultOutput = TextOutput
)

var OutputMapping = map[Output]string{
	TextOutput: "text",
	JSONOutput: "json",
	YAMLOutput: "yaml",
}

func (o Output) String() string {
	return OutputMapping[o]
}

func ToOutput(s string) Output {
	switch s {
	case "text":
		return TextOutput
	case "json":
		return JSONOutput
	case "yaml":
		return YAMLOutput
	default:
		return DefaultOutput
	}
}
