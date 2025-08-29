{
  buildEnv,
  dockerTools,
  external-dns-bunny-webhook,
  external-dns-bunny-webhook-rev,
}:
let
  paths = [
    external-dns-bunny-webhook
    dockerTools.caCertificates
  ];

  pathsToLink = [
    "/bin"
    "/etc"
  ];

  spec = {
    name = "nossa.ee/talya/external-dns-bunny-webhook";
    tag = external-dns-bunny-webhook-rev;
    config = {
      Cmd = [ "/bin/webhook" ];
      User = "1000:1000";
    };
  };
in
{
  build = dockerTools.buildImage (
    spec
    // {
      copyToRoot = buildEnv {
        name = "external-dns-bunny-webhook-root";
        inherit paths pathsToLink;
      };
    }
  );
  stream-layered = dockerTools.streamLayeredImage (
    spec
    // {
      contents = paths;
      maxLayers = 120;
    }
  );
}
