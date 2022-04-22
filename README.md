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

TODO

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
