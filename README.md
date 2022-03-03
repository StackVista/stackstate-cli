# StackState CLI 

## Getting development started

1. Install the latest version of Golang
2. Run `go run main.go` 

Please read the command development rules. 

## Command Development Rules

To ensure a superb command line experience for our users development of new CLI commands must adhere to strict rules. 

### Rule: familiarize yourself with CLI development

CLI development has a big UX component to it. Please familourize yourself with CLI development by reading the following two resources:
 1. [Command Line Interface Guidelines](https://clig.dev)
 2. [12 Factor CLI Apps](https://medium.com/@jdxcode/12-factor-cli-apps-dd3c227a0e46)


### Rule: noun, verb, flags

Every CLI command must follow this structure:

 1. Noun - a noun is given on which the command operates on, e.g. `missile`
 2. Verb - this is always followed by a verb, e.g. `lauch`
 3. Flags - optional parameters given to this required combination of noun and verb are controlled with flags, e.g. `--destination graceful`


### Rule: always use the Printer to print and the logger for verbosity

Please only use the `Printer` to print to stdout and stderr. This is to ensure consistency. If there is anything that you are missing, you will have to update the Printer.

If you want to print intermediate results of what the CLI is doing for troubleshooting purposes please use the logger. The logger is also used by the OpenAPI generated code and when `verbose` mode is enabled any requests made are printed by that code. 

## Rule: return CLIError from CLI commands when there is an error

Instead of printing errors directly using the Printer, please return errors from the CLI command's execution. This is to ensure error printing and handling is done consistently. 

When you have an error related to the input of the CLI you should use a `CLIError` that set `ShowUsage` to true. When the error is related to the API or due to technical issues `ShowUsage` should be false.