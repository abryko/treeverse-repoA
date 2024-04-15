treeverse: {
  # appA: Go webserver, serving hello requests
  appA = treeverse.pkgs.buildGoModule {
    name = "appA";
    src = ./appA;
    vendorHash = null;
    ldflags = [
      "-X main.environment=${treeverse.variables.env}"
    ];
    postInstall = ''
      mv $out/bin/treeverse-repoA $out/bin/appA
    '';
  };

  # nixosModuleA: nixos module declaring systemd unit for appA
  nixosModuleA = {
    systemd.services.appA = {
      description = "appA webservice";
      wantedBy = ["multi-user.target"];
      after = ["network.target"];
      script = "${treeverse.repositories.repoA.appA}/bin/appA";
    };
  };
}
