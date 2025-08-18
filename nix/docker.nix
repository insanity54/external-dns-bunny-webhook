{
  lib,
  system,
  dockerTools,
  external-dns-bunny-webhook,
  external-dns-bunny-webhook-rev,
}:
let
  spec = {
    name = "nossa.ee/talya/external-dns-bunny-webhook";
    tag = external-dns-bunny-webhook-rev;
    config = {
      Cmd = [ "/bin/webhook" ];
      User = "1000:1000";
    };

    contents = [
      external-dns-bunny-webhook
      dockerTools.caCertificates
    ];
  };
in
if builtins.elem system lib.platforms.darwin then
  dockerTools.buildImage spec
else
  dockerTools.streamLayeredImage (spec // { maxLayers = 120; })
