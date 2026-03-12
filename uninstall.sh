#!/bin/sh
set -e

BINARY_NAME="clifileconverter"
INSTALL_DIR="/usr/local/bin"
TARGET="$INSTALL_DIR/$BINARY_NAME"

info()    { printf '\033[0;34m[INFO]\033[0m  %s\n' "$1"; }
success() { printf '\033[0;32m[OK]\033[0m    %s\n' "$1"; }
error()   { printf '\033[0;31m[ERROR]\033[0m %s\n' "$1" >&2; exit 1; }

if [ ! -f "$TARGET" ]; then
    error "Binary not found at $TARGET"
fi

info "Removing $TARGET"

if [ -w "$INSTALL_DIR" ]; then
    rm "$TARGET"
else
    info "Root access required"
    sudo rm "$TARGET"
fi

success "clifileconverter has been uninstalled"