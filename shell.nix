{ pkgs ? import <nixpkgs> {} }:

with pkgs;

# proto-lens-protoc is broken for later ghc versions
let haskellPackages = haskell.packages.ghc928; in

mkShell {
  buildInputs = [
    buf
    protobuf
    protoc-gen-prost
    protoc-gen-prost-crate
    protoc-gen-prost-serde
    protoc-gen-tonic
    haskellPackages.proto-lens-protoc
  ];
}
