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
            go_1_18
            goreleaser
          ]);
        };

        packages = {
          sts = pkgs.buildGo118Module {
            pname = "sts";
            version = "0.0.1";

            src = ./.;

            # This hash locks the dependencies of this package.
            # Change it to the provided when the go dependencies change.
            # See https://www.tweag.io/blog/2021-03-04-gomod2nix/ for details
            vendorSha256 = "sha256-+jA7RJ6Eh9e6JjyTFDifcN+iVWE66w3tRzBwj2b8ijM=";

            postInstall = ''
              mv $out/bin/stackstate-cli2 $out/bin/sts
            '';
          };
        };

        defaultPackage = self.packages."${system}".sts;

        apps = {
          sts = {
            type = "app";
            program = "${self.packages."${system}".sts}/bin/sts";
          };
        };

        defaultApp = self.apps."${system}".sts;

      });
}
