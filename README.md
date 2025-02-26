# StackState CLI

## Documentation

Documentation of this CLI is still under development:

 * Development Branch - https://github.com/StackVista/stackstate-docs/tree/STAC-15294-cli2-docs
 * Gitbook Preview - https://stackstate.gitbook.io/stackstate-docs-development/v/stac-15294-cli2-docs/setup/cli/

### Getting development started

1. Install [`pre-commit`](https://pre-commit.com/) and install all the hooks:
```
pre-commit install
```
2. Install Nix package manager. See [Nix installation](#nix-installation)
3. Enable Nix Flakes support
```
mkdir -p ~/.config/nix && echo "experimental-features = nix-command flakes" >> ~/.config/nix/nix.conf
```
4. Enter development shell `nix develop`
5. Run `go run main.go`

It's also possible to use nix phases to test `nix develop --check` and build the application `nix develop --build`, as specified in the Makefile, without entering the shell.

6. Read the [command development guide](CMD_DEVELOPMENT.md).

## How to create a new release

 1. Check which version of the CLI is the latest by running `scripts/publish/print_latest_version.sh`
 2. Determine the semantic version number you want to use:
    - Breaking change? Major version! This should never happen after the initial release!
    - New commands and/or flags? Minor version.
    - Bugfixes/improvements of existing commands? Patch version.
 3. On a clean and committed and pushed commit run `scripts/create_release.sh -v VERSION` with your version.

 A tag will be pushed to Gitlab with the format: `v${VERSION}`. Gitlab's CI will then run the publishing pipeline.

## Nix installation

Install Nix package manager by following the [official installation instructions](https://nixos.org/download.html).

To install the package to the current profile you need to register the current repository

```sh
nix registry add stackstate-cli "git+ssh://git@github.com/stackvista/stackstate-cli"
nix profile install stackstate-cli#sts
```

Alternatively temporary shell environment with `sts` binary in $PATH can be started using

```
nix shell "git+ssh://git@github.com/stackvista/stackstate-cli?ref=master"
```

Where `ref=` can reference to any branch name

## CI dependencies

All the CI dependencies are also defined by flake.nix inside the `ciDeps` directive. It is possible to start bash session with all these dependencies in scope by calling `nix develop #.ci`.

To build the base image locally run

```
nix build .#ci-image & docker load < result

# Test
docker run -ti --rm stackstate-cli2-ci:latest go version
```

## Working with openapi

This repository pulls the stackstate-api spec from the [openapi repository](https://gitlab.com/stackvista/platform/stackstate-openapi).

### Bumping the openapi version
- Change the version/branch/commit sha in the `stackstate_openapi/openapi_version` file
- Run `nix develop -c ./scripts/generate_stackstate_api.sh`
- Commit the generated code

CI will check whether the requested api version and generated code are kep up to date.

### For development: Updating generated code without a OpenApi commit
- Run `nix develop -c ./scripts/run_openapi_local.sh {path to your open api folder}`
- example: `nix develop -c ./scripts/run_openapi_local.sh /Users/projects/stackstate-openapi`
