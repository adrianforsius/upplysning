with import <nixpkgs> { };
mkShell {
  NIX_LD_LIBRARY_PATH = lib.makeLibraryPath [
    stdenv.cc.cc
    openssl

    # chrome on nixos dependencies
    glibc
    glib
    nss
    nspr
    dbus
    atk
    cups
    libdrm
    expat
    libx11
    libxcomposite
    libxdamage
    libxext
    libxfixes
    libxrandr
    libgbm
    libxcb
    libxkbcommon
    pango
    cairo
    alsa-lib
    # ...

  ];
  NIX_LD = lib.fileContents "${stdenv.cc}/nix-support/dynamic-linker";
}
