#!/bin/sh

# Dapatkan direktori tempat skrip ini berada
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"

# Set path konfigurasi relatif terhadap root proyek
CONFIG_PATH="$SCRIPT_DIR/config/config.yaml"

# Jalankan aplikasi dengan path konfigurasi
"$SCRIPT_DIR/bin/main"
