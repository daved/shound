{ pkgs ? import <nixpkgs> {} }:

pkgs.stdenv.mkDerivation {
  name = "shound";
  src = ./.;

  nativeBuildInputs = [
    pkgs.pkg-config
    pkgs.alsa-lib
    pkgs.zig
    pkgs.go
  ];

  buildPhase = ''
    mkdir -p /build/cache/{go,zig}
    export GOCACHE=/build/cache/go
    export ZIG_GLOBAL_CACHE_DIR=/build/cache/zig

    export CGO_ENABLED=1
    export CC="zig cc"

    go build ./cmd/shound
  '';

  installPhase = ''
    mkdir -p $out/bin
    cp shound $out/bin
  '';
}
