{
  description = "A flake for utxorpc/spec";

  inputs = {
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.simpleFlake {
      inherit self nixpkgs;
      name = "utxorpc/spec";
      shell = ./shell.nix;
    };
}
