# StackState CLI 

### Getting development started

1. Install Nix package manager. See [Nix installation](#nix-installation)
2. Enable Nix Flakes support
```
mkdir -p ~/.config/nix && echo "experimental-features = nix-command flakes" >> ~/.config/nix/nix.conf
```
3. Enter development shell `nix develop`
4. Run `go run main.go`

It also possible to use nix phases to test `nix develop --check` and build the application `nix develop --build`, as specified in the Makefile, without entering the shell.

5. Read the [command development styling rules](STYLE.md).

## How to create a new release

 1. Determine the semantic version number you want to use:
    - Breaking change? Major version! This should never happen after the initial release!
    - New commands and/or flags? Minor version.
    - Bugfixes/improvements of existing commands? Patch version.
 2. Tag and push the commit you wish to release with a tag of the format "v${version}":

```sh
git tag v1.2.1 -m "v1.2.1"
git push origin v1.2.1
```
 
## Nix installation

Install Nix package manager by following the [official installation instructions](https://nixos.org/download.html).

To install the package to the current profile you need to register the current repository

```sh
nix registry add stackstate-cli "git+ssh://git@gitlab.com/stackvista/stackstate-cli2"
nix profile install stackstate-cli#sts
```

Alternatively temporary shell environment with `sts` binary in $PATH can be started using

```
nix shell "git+ssh://git@gitlab.com/stackvista/stackstate-cli2?ref=master"
```

Where `ref=` can reference to any branch name
