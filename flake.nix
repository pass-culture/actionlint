{
  description = "actionlint";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
  };

  outputs =
    { self, nixpkgs }:
    let
      supportedSystems = [
        "x86_64-linux"
        "aarch64-linux"
        "x86_64-darwin"
        "aarch64-darwin"
      ];
      forAllSystems = nixpkgs.lib.genAttrs supportedSystems;
    in
    {
      packages = forAllSystems (
        system:
        let
          pkgs = nixpkgs.legacyPackages.${system};
        in
        {
          default = pkgs.buildGoModule {
            name = "actionlint";
            src = ./.;
            vendorHash = "sha256-E5CqRChnq6ccwu7kiyfJJnkLtER5LzgOi6tK0lnpcrs=";
            subPackages = "cmd/actionlint";
          };
        }
      );
      devShells = forAllSystems (
        system:
        let
          pkgs = nixpkgs.legacyPackages.${system};
        in
        {
          default = pkgs.mkShell {
            packages = [
              pkgs.go_1_25
              pkgs.go-tools
              pkgs.govulncheck
              pkgs.shellcheck
              pkgs.shellcheck
              pkgs.python314Packages.pyflakes
            ];
          };
        }
      );
    };
}
