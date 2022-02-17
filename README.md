# ⚡️ Fast nvm switcher

Is your nvm slow? This repository provides some suggestions (and a utility) for speeding up nvm in your shell.

## Speeding up a basic installation

The most expensive part of having nvm installed on your machine is running `nvm.sh` in a new shell.

That's the second line in the initialization below:

```bash
export NVM_DIR="$HOME/.nvm"
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"  # This loads nvm
[ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"  # This loads nvm bash_completion
```

Simply define your own `nvm` function right in your shell configuration that loads nvm only when you call it, like so:

```bash
export NVM_DIR="$HOME/.nvm"
[ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"  # This loads nvm bash_completion

function nvm() {
  [ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"  # This loads nvm

  nvm $@
}
```

## Deeper shell integration

For the zsh shell hook that automatically switches node versions on directory change, you can simply start by downloading the `resolve_node_version` binary to your `.nvm` folder:

```bash
curl -o- https://raw.githubusercontent.com/abejfehr/fast-nvm-switcher/v0.1.6/install.sh | bash
```

And follow the instructions given.

Once you've complete that, changing directories should now also be blazing fast ⚡️

> **Note:** If you already use one of the "Deeper shell integrations" from the [nvm README](https://github.com/nvm-sh/nvm#deeper-shell-integration), you'll have to remove those before following these instructions.

## Limitations

- This script only supports the 'default' alias, it won't work with 'lts' or other custom aliases
- The `resolve_node_version` utility does not update the `$MANPATH` environment variable
