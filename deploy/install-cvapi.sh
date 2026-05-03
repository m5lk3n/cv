#!/usr/bin/env bash
# Install cvapi as a systemd service on Ubuntu.
#
# Usage (run on the server, as root or via sudo):
#   sudo ./deploy/install-cvapi.sh /path/to/built/cvapi
#
# Defaults to ./bin/cvapi if no argument is given.
set -euo pipefail

BIN_SRC="${1:-./bin/cvapi}"
BIN_DST="/usr/local/bin/cvapi"
UNIT_SRC="$(dirname "$0")/cvapi.service"
UNIT_DST="/etc/systemd/system/cvapi.service"
LOG_DIR="/var/log/caddy"
LOG_FILE="${LOG_DIR}/cvapi.log"
SVC_USER="caddy"
SVC_GROUP="caddy"

if [[ $EUID -ne 0 ]]; then
    echo "error: must be run as root (try sudo)" >&2
    exit 1
fi

if [[ ! -f "$BIN_SRC" ]]; then
    echo "error: binary not found at $BIN_SRC" >&2
    echo "build it first with 'make build-api', or pass an explicit path." >&2
    exit 1
fi

if ! id "$SVC_USER" >/dev/null 2>&1; then
    echo "error: user '$SVC_USER' does not exist (install Caddy first, or edit User= in cvapi.service)" >&2
    exit 1
fi

install -m 0755 -o root -g root "$BIN_SRC" "$BIN_DST"
install -m 0644 -o root -g root "$UNIT_SRC" "$UNIT_DST"

mkdir -p "$LOG_DIR"
chown "$SVC_USER:$SVC_GROUP" "$LOG_DIR"
touch "$LOG_FILE"
chown "$SVC_USER:$SVC_GROUP" "$LOG_FILE"
chmod 0644 "$LOG_FILE"

systemctl daemon-reload
systemctl enable --now cvapi.service
systemctl --no-pager --full status cvapi.service || true

echo
echo "cvapi installed and started."
echo "  binary: $BIN_DST"
echo "  unit:   $UNIT_DST"
echo "  log:    $LOG_FILE"
echo
echo "follow logs with: sudo tail -f $LOG_FILE"
