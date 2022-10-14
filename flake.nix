{
  description = "StackState CLI";

  nixConfig.bash-prompt = "STS CLI 2 $ ";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let  
        pkgs = import nixpkgs { inherit system; overlays = [ ]; };
        pkgs-linux = import nixpkgs { system = "x86_64-linux"; overlays = [ ]; };

        # Dependencies used for both development and CI/CD
        sharedDeps = pkgs: (with pkgs; [
          bash
          go_1_19
          gotools
          diffutils # Required for golangci-lint
          golangci-lint
          openapi-generator-cli
        ]);

        # Dependencies used only by CI/CD
        ciDeps = pkgs: (with pkgs; [
          git
          cacert
          gcc
          coreutils-full
          goreleaser
          awscli
          docker
        ]);

        darwinDevShellExtraDeps = pkgs: pkgs.lib.optionals pkgs.stdenv.isDarwin (with pkgs.darwin.apple_sdk_11_0; [
          Libsystem 
          IOKit
        ]);
      in {

        devShells = {
          dev = pkgs.mkShell {
            buildInputs = sharedDeps(pkgs) ++ darwinDevShellExtraDeps(pkgs);
          };

          ci = pkgs.mkShell {
            buildInputs = sharedDeps(pkgs) ++ ciDeps(pkgs);
          };
        };

        devShell = self.devShells."${system}".dev;

        packages = {
          sts = pkgs.buildGo119Module {
            pname = "sts";
            version = "2.0.0";

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
            vendorSha256 = "sha256-aXTDHT1N+4Qpkuxb8vvBvP2VPyS5ofCgX6XFhJ5smUQ=";

            postInstall = ''
              mv $out/bin/stackstate-cli2 $out/bin/sts
            '';
          };

          ci-image = pkgs.dockerTools.buildImage {
            name = "stackstate-cli2-ci";
            tag = "latest";
            created = "now";

            contents = sharedDeps(pkgs-linux) ++ ciDeps(pkgs-linux);

            config = {
              Env = [
                "GIT_SSL_CAINFO=/etc/ssl/certs/ca-bundle.crt"
                "SSL_CERT_FILE=/etc/ssl/certs/ca-bundle.crt"
              ];

              # Required to make golangci-lint work.
              Volumes = {
                "/tmp" = {};
              };
            };
          };

          default = self.packages."${system}".sts;
        };

        apps = {
          sts = {
            type = "app";
            program = "${self.packages."${system}".sts}/bin/sts";
          };

          default = self.apps."${system}".sts;
        };
      });
}
