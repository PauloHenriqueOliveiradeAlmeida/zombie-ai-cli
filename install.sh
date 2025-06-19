#!/bin/bash

BIN_NAME="zombie"
INSTALL_DIR="$HOME/.local/bin"
INSTALL_PATH="$INSTALL_DIR/$BIN_NAME"

OS=$(uname | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

if [[ "$ARCH" == "x86_64" ]]; then
  ARCH="amd64"
fi

BIN_FILE="$BIN_NAME-$OS-$ARCH"

BIN_URL="https://github.com/PauloHenriqueOliveiradeAlmeida/zombie-ai-cli/releases/latest/download/$BIN_FILE"

echo "Instalando $BIN_NAME em $INSTALL_DIR..."
mkdir -p "$INSTALL_DIR"

echo "Baixando binário de $BIN_URL..."
curl -L "$BIN_URL" -o "$INSTALL_PATH"

if [[ $? -ne 0 ]]; then
  echo "Falha no download do binário."
  exit 1
fi

chmod +x "$INSTALL_PATH"

if ! echo "$PATH" | grep -q "$INSTALL_DIR"; then
  echo ""
  echo "ATENÇÃO: $INSTALL_DIR não está no PATH."
  echo "Adicione ao seu ~/.bashrc ou ~/.zshrc:"
  echo "export PATH=\"$INSTALL_DIR:\$PATH\""
else
  echo "$INSTALL_DIR já está no PATH."
fi

echo "Use '$BIN_NAME configure' para começar!"
