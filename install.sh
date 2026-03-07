#!/bin/sh
set -e

REPO="almeidafm/clifileconverter"
BINARY_NAME="clifileconverter"
INSTALL_DIR="/usr/local/bin"
VERSION="${VERSION:-latest}"

info()    { printf '\033[0;34m[INFO]\033[0m  %s\n' "$1"; }
success() { printf '\033[0;32m[OK]\033[0m    %s\n' "$1"; }
error()   { printf '\033[0;31m[ERROR]\033[0m %s\n' "$1" >&2; exit 1; }

detect_platform() {
    OS="$(uname -s)"
    ARCH="$(uname -m)"

    case "$OS" in
        Linux)  OS="linux" ;;
        Darwin) OS="darwin" ;;
        *)      error "Unsupported OS: $OS" ;;
    esac

    case "$ARCH" in
        x86_64 | amd64) ARCH="amd64" ;;
        arm64  | aarch64) ARCH="arm64" ;;
        *)      error "Unsupported architecture: $ARCH" ;;
    esac

    PLATFORM="${OS}_${ARCH}"
    info "Detected platform: $PLATFORM"
}

resolve_version() {
    if [ "$VERSION" = "latest" ]; then
        info "Fetching latest release version..."
        VERSION="$(curl -fsSL "https://api.github.com/repos/${REPO}/releases/latest" \
            | grep '"tag_name"' \
            | sed -E 's/.*"tag_name": *"([^"]+)".*/\1/')"
        [ -z "$VERSION" ] && error "Could not determine latest version."
    fi
    info "Version: $VERSION"
}


download_and_install() {

    ASSET="${BINARY_NAME}_${PLATFORM}.tar.gz"
    URL="https://github.com/${REPO}/releases/download/${VERSION}/${ASSET}"

    TMP_DIR="$(mktemp -d)"
    trap 'rm -rf "$TMP_DIR"' EXIT

    info "Downloading $URL ..."
    curl -fsSL "$URL" -o "$TMP_DIR/$ASSET" \
        || error "Download failed. Check the version and your internet connection."

    info "Extracting..."
    tar -xzf "$TMP_DIR/$ASSET" -C "$TMP_DIR"

    BINARY_PATH="$(find "$TMP_DIR" -type f -name "$BINARY_NAME" | head -n 1)"
    [ -z "$BINARY_PATH" ] && error "Binary '$BINARY_NAME' not found in archive."

    chmod +x "$BINARY_PATH"

    if [ -w "$INSTALL_DIR" ]; then
        mv "$BINARY_PATH" "$INSTALL_DIR/$BINARY_NAME"
    else
        info "Root access required to install to $INSTALL_DIR"
        sudo mv "$BINARY_PATH" "$INSTALL_DIR/$BINARY_NAME"
    fi

    success "Installed $BINARY_NAME to $INSTALL_DIR/$BINARY_NAME"
}

verify() {
    if command -v "$BINARY_NAME" >/dev/null 2>&1; then
        success "Installation complete! Run: $BINARY_NAME --help"
    else
        info "Binary installed but not found in PATH."
        info "Add this to your shell config (~/.bashrc or ~/.zshrc):"
        info "  export PATH=\"$INSTALL_DIR:\$PATH\""
    fi
}

main() {
    detect_platform
    resolve_version
    download_and_install
    verify
}

main