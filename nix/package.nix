{
  lib,
  buildGoModule,
  ...
}:
let
  src = lib.fileset.toSource rec {
    root = ../.;
    fileset = lib.fileset.unions [
      (root + /cmd)
      (root + /go.mod)
      (root + /go.sum)
      (root + /internal)
    ];
  };
in
buildGoModule {
  pname = "external-dns-bunny-webhook";
  version = "0.4.0";

  inherit src;

  vendorHash = "sha256-goHQNnDh2vzfnIMlIhY5QgJ0StioG54QHSC3VvP9Y+U=";
}
