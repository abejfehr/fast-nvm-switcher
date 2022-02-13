# Fast NVM Switcher

This repository provides an alternative shell configuration (specifically for zsh) switches node versions on directory change, but in a _significantly_ faster way than the default configuration provided in nvm's setup.

## How to use


Begin by installing [nvm](https://github.com/nvm-sh/nvm) using the download script from their README.

If you already have nvm installed, remove anything nvm-related you might already have from your local .zshrc file.

Next, run the following script to download the `resolve_node_version` binary to your machine:

```sh
(cd $HOME/.nvm/ && curl -L -O https://github.com/abejfehr/fast-nvm-switcher/releases/download/0.1.1/resolve_node_version)
```

Add the following lines to your .zshrc

```sh
export NVM_DIR="${HOME}/.nvm"

# Lazy loads nvm when running any nvm command
nvm() {
  [[ -s "${NVM_DIR}/nvm.sh" ]] && \. "${NVM_DIR}/nvm.sh"
  nvm $@
}

# Strips the path of previous nvm node directories
nvm_strip_path() {
  command printf %s "${1-}" | command awk -v NVM_DIR="${NVM_DIR}" -v RS=: '
  index($0, NVM_DIR) == 1 {
    path = substr($0, length(NVM_DIR) + 1)
    if (path ~ "^(/versions/[^/]*)?/[^/]*'"${2-}"'.*$") { next }
  }
  { print }' | command paste -s -d: -
}

# Resolves node version based on nearest nvmrc and adds its directory to the PATH
load-nvmrc() {
  NODE_PATH=$(${NVM_DIR}/resolve_node_version)
  echo "Updating node location to $NODE_PATH"
  PATH="$PATH:$NODE_PATH"
}

autoload -U add-zsh-hook
add-zsh-hook chpwd load-nvmrc
load-nvmrc
```

## Limitations

- This script only supports the 'default' alias, it won't work with lts, etc.
- The install script assumes that you're using the default `$NVM_DIR` of `$HOME/.nvm`
