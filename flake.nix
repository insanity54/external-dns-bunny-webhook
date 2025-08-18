{
  description = "external-dns provider for bunny.net";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-25.05";
  };

  outputs =
    { self, nixpkgs }:
    let
      systems = [
        "aarch64-darwin"
        "aarch64-linux"
        "x86_64-darwin"
        "x86_64-linux"
      ];
      eachSystem = nixpkgs.lib.genAttrs systems;
    in
    {
      formatter = eachSystem (system: nixpkgs.legacyPackages.${system}.nixfmt-rfc-style);

      packages = eachSystem (
        system:
        let
          pkgs = nixpkgs.legacyPackages.${system};
        in
        rec {
          default = external-dns-bunny-webhook;

          external-dns-bunny-webhook = pkgs.callPackage ./nix/package.nix { };

          external-dns-bunny-webhook-docker = pkgs.callPackage ./nix/docker.nix {
            inherit external-dns-bunny-webhook;
            external-dns-bunny-webhook-rev = self.rev or "dev";
          };
        }
      );

      devShell = eachSystem (
        system:
        let
          pkgs = nixpkgs.legacyPackages.${system};
        in
        pkgs.mkShell {
          name = "external-dns-bunny-webhook";
          packages = with pkgs; [
            go
            gopls
          ];
        }
      );
    };
}
