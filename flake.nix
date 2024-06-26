{
  description = "mdhtml flake";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in
      {
        defaultPackage = (import ./default.nix { inherit pkgs; });
        devShells.default = pkgs.mkShell
          {
            buildInputs = with pkgs; [
              go
              golangci-lint
            ];
          };
      }
    );
}
