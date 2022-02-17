#!/usr/bin/env bash

function check_compatibility() {
  echo ""
  echo ""

  if [ -f "$HOME/.zshrc" ]; then
    echo "✅ Shell is compatible"
  else
    echo "❌ Shell is not compatible"
    exit 1
  fi

  if [ $(uname -s) == "Darwin" ]; then
    echo "✅ OS is compatible"
  else
    echo "❌ OS is not compatible"
    exit 1
  fi

  if [ $(uname -p) == "i386" ]; then
    echo "✅ Architecture is compatible"
  elif [ $(uname -p) == "arm" ]; then
    echo "❌ Architecture is not compatible"
    exit 1
  else
    echo "❌ Architecture is not compatible"
    exit 1
  fi

  echo ""
  echo ""
}

function echo_shell_instructions() {
  echo ""
  echo ""
  echo "Finalize your shell integration by copying the following lines to your .zshrc:"
  echo ""
  echo ""
  echo "  nvm_strip_path() {"
  echo "    command printf %s "\${1-}" | command awk -v NVM_DIR="\${NVM_DIR}" -v RS=: '"
  echo "     index(\$0, NVM_DIR) == 1 {"
  echo "       path = substr(\$0, length(NVM_DIR) + 1)"
  echo "       if (path ~ "^\(/versions/[^/]*\)?/[^/]*'"${2-}"'.*$") { next }"
  echo "     }"
  echo "     { print }' | command paste -s -d: -"
  echo "   }"
  echo "  "
  echo "   load-nvmrc() {"
  echo "     NODE_PATH=\$(\${NVM_DIR}/resolve_node_version)"
  echo "     if [ -n "\$NODE_PATH" ]; then"
  echo "       PATH="\$PATH:\$NODE_PATH""
  echo "       echo "Note location set to \$NODE_PATH""
  echo "     fi"
  echo "   }"
  echo "  "
  echo "  autoload -U add-zsh-hook"
  echo "  add-zsh-hook chpwd load-nvmrc"
  echo "  load-nvmrc"
  echo ""
  echo ""
}

function download_and_install_utility() {
  cd $NVM_DIR

  curl -L -O https://github.com/abejfehr/fast-nvm-switcher/releases/download/v0.1.5/resolve_node_version

  chmod +x $NVM_DIR/resolve_node_version

  echo ""
  echo ""
  echo "Successfully downloaded the resolve_node_version utility to your $NVM_DIR directory"
}

check_compatibility
download_and_install_utility
echo_shell_instructions

