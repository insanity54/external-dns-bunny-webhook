{
  dockerTools,
  external-dns-bunny-webhook,
  external-dns-bunny-webhook-rev,
  ...
}:
dockerTools.streamLayeredImage {
  name = "nossa.ee/talya/external-dns-bunny-webhook";
  tag = external-dns-bunny-webhook-rev;
  config = {
    Cmd = [ "/bin/webhook" ];
    User = "1000:1000";
  };

  contents = [
    external-dns-bunny-webhook
    dockerTools.binSh
  ];

  maxLayers = 120;
}
