{ pkgs ? import <nixpkgs> {} }:

pkgs.stdenv.mkDerivation {
  name = "shound";
  src = ./.;
  #dontUnpack = true;

  nativeBuildInputs = [
    pkgs.pkg-config
    pkgs.alsaLib
    pkgs.zig
    pkgs.go
  ];

  buildPhase = ''
    mkdir -p /build/cache/{go,zig}
    export GOCACHE=/build/cache/go
    export ZIG_GLOBAL_CACHE_DIR=/build/cache/zig

    export CGO_ENABLED=1
    export CC="zig cc -target x86_64-linux-gnu -static -I${pkgs.alsaLib}/lib"
    export CXX="zig c++ -target x86_64-linux-gnu -static -I${pkgs.alsaLib}/lib"

    go build -ldflags '-linkmode "external" -extldflags "-static"' ./cmd/shound
  '';

  installPhase = ''
    mkdir -p $out/bin
    cp shound $out/bin
  '';
}
