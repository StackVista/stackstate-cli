{
  description = "StackState CLI";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in {

        devShell = pkgs.mkShell {
          buildInputs = (with pkgs; [
            go
            goreleaser
          ]);
        };

        packages = {
          sts = pkgs.buildGoModule {
            pname = "sts";
            version = "0.0.1";

            src = ./.;

            # This hash locks the dependencies of this package.
            # Change it to the provided when the go dependencies change.
            # See https://www.tweag.io/blog/2021-03-04-gomod2nix/ for details
            vendorSha256 = "sha256-vf0zhp2Y8MUJRvyBhJDLpmoooM8I8cboTFxR3JKs2Ns=";

            postInstall = ''
            mv $out/bin/stackstate-cli2 $out/bin/sts
          '';
          };
        };

        defaultPackage = self.packages."${system}".sts;
      });
}
