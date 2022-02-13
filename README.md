# Fast NVM Switcher

This repository provides an alternative shell configuration (specifically for zsh) switches node versions on directory change, but in a _significantly_ faster way than the default configuration provided in nvm's setup.

## How to use


Begin by installing [nvm](https://github.com/nvm-sh/nvm) using the download script from their README.

If you already have nvm installed, remove anything nvm-related you might already have from your local .zshrc file.

Next, run the following script to download the `resolve_node_version` binary to your machine:

```sh
curl -O --output-dir $HOME/.nvm URL_ONCE_I_HAVE_IT
```

Add the following lines to your .zshrc

```sh
export NVM_DIR="${HOME}/.nvm"

# Lazy loads nvm when running any nvm command
nvm() {
  [[ -s "${NVM_DIR}/nvm.sh" ]] && \. "${NVM_DIR}/nvm.sh"
  nvm $@
}

# Resolves node version based on nearest nvmrc and adds its directory to the PATH
load-nvmrc() {
  NODE_PATH=$(${NVM_DIR}/resolve_node_version)
  PATH="$PATH:$NODE_PATH"
}

autoload -U add-zsh-hook
add-zsh-hook chpwd load-nvmrc
load-nvmrc
```

## Limitations

- This script only supports the 'default' alias, it won't work with lts, etc.
- The install script assumes that you're using the default `$NVM_DIR` of `$HOME/.nvm`
