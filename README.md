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

 1. Noun - a noun is given on which the command operates, e.g. `missile`
 2. Verb - this is always followed by a verb, e.g. `lauch`
 3. Flags - parameters given to this required combination of noun and verb are controlled with flags, e.g. `--destination graceful`


### Rule: always use the Printer to print and the logger for verbosity

Please only use the `Printer` to print to stdout and stderr. This is to ensure consistency. If there is anything that you are missing, you will have to update the Printer.

If you want to print intermediate results of what the CLI is doing for troubleshooting purposes please use the logger. The logger is also used by the OpenAPI generated code and when `verbose` mode is enabled any requests made are printed by that code. 

### Rule: return CLIError from CLI commands when there is an error

Instead of printing errors directly using the Printer, please return errors from the CLI command's execution. This is to ensure error printing and handling is done consistently. 

When you have an error related to the input of the CLI you should use a `CLIError` that set `ShowUsage` to true. When the error is related to the API or due to technical issues `ShowUsage` should be false.

### Rule: each command must contain clear help text

`Use:` - a skeleton example of command usage:
* **noun** level command: Show only the command name, `[command]` will be appended in the CLI help. For example:
  ```text
  # The `Use:` text for the `script` command:
    Use:   "script"
  
  # will result in the following being printed in the CLI help:
    Usage:
      sts script [command]
  ```

* **noun verb** level command: Any mutually exclusive or required flags should be included in this description, other flags will be covered under the general group `[flags]`, which is automatically added to the end of the `Usage:` text when printed in the CLI help. For example:
    ```text
    # the `Use:` text for the `script execute` command
        Use:   "execute {-s SCRIPT | -f FILE}"
    
    # will result in the following being printed in the CLI help
      Usage:
        sts script execute {-s SCRIPT | -f FILE} [flags]
    ```
  * Use curly brackets `{}` to delimit REQUIRED flags (always show)
  * Use square brackets `[]` to delimit OPTIONAL flags (only show when mutually exclusive)
  * Use a pipe `|` to indicate mutually exclusive flags - only one of the options shown can be used

`Short:` - a concise description of the command:
* Start with a lower case letter
* Maximum one sentence in length
* No full stop (`.`) at the end
* Start with a verb and then describe what that verb works on. This should provide more info than the command.

`Example:` - copy/paste examples of command usage:
* Each example should start with a commented description of what it does.
  * Start with a lower case letter
  * Maximum one sentence in length
  * No full stop (`.`) at the end
  * Start with a verb and then describe what the command does
* Examples should be given that cover as many of the available options as possible.
* Try to stick to no more than 3 examples.
* For example:
  ```text
  # the Example: text for the `script execute` command
    Example: "# execute a script from file\n"+
      "sts execute --file \"path/to/my.script\"\n"+
      "\n"+
      "# execute a script with variables provided by an arguments-script\n"+
      "sts execute --script \"x+y\" --arguments-script \"[x: 1, y: 2]\"",
  
  # will result in the following being printed in the CLI help
    Examples:
    # execute a script from file
    sts execute --file "path/to/my.script"

    # execute a script with variables provided by an arguments-script
    sts execute --script "x+y" --arguments-script "[x: 1, y: 2]"  
  ```
