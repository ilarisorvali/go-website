#!/usr/bin/env bash
set -euo pipefail

QUADLET_DIR="${HOME}/.config/containers/systemd"

mkdir -p "$QUADLET_DIR"

moved_services=()

echo "Installing Quadlet files..."

for file in *.container; do
    [[ -f "$file" ]] || continue

    install -Dm0644 "$file" "$QUADLET_DIR/"
    echo "  Installed: $file"

    service="$(basename "$file")"
    service="${service%.*}.service"

    moved_services+=("$service")
done

if [[ ${#moved_services[@]} -eq 0 ]]; then
    echo "No .container files found."
    exit 0
fi

echo "Reloading user systemd..."
systemctl --user daemon-reload

echo "Done."
