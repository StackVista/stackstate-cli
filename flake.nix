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
            golangci-lint
            openapi-generator-cli
          ]);
        };

        packages = {
          sts = pkgs.buildGo118Module {
            pname = "sts";
            version = "0.0.1";

            src = ./.;

            # This hash locks the dependencies of this package.
            # Change it to the provided when the go dependencies change.
            # See https://www.tweag.io/blog/2021-03-04-gomod2nix/ for details.
            #
            # NOTE In case if your build fails due to incosistency in vendor modules
            # Comment out the real hash and uncomment the fake one then on next `nix build .` run
            # you will get a new real hash which can be used here.
            #
            # vendorSha256 = pkgs.lib.fakeSha256;
            vendorSha256 = "sha256-ogEbNzB78RGtrVM427ADwTcWcbY+ofIn0jxSWOdjplQ=";

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
