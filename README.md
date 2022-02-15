# Fast NVM Switcher

This repository provides an alternative shell configuration (specifically for zsh) switches node versions on directory change, but in a _significantly_ faster way than the default configuration provided in nvm's setup.

## How fast is it?

The following benchmarks are from a 2019 16-inch MacBook Pro, where nvm was particularly slow. The output was obtained using [zshprof](https://github.com/raboof/zshprof).

### Default nvm

Shell startup times on my machine were always _on the order of seconds_.

<details>

<summary>View zshprof output</summary>

```
num  calls                time                       self            name
-----------------------------------------------------------------------------------
 1)    1        1140.57  1140.57   32.40%   1140.57  1140.57   32.40%  load-nvmrc
 2)    1        2301.00  2301.00   65.36%   1054.46  1054.46   29.95%  nvm_auto
 3)    2        1246.54   623.27   35.41%    694.29   347.15   19.72%  nvm
 4)    1         498.15   498.15   14.15%    454.94   454.94   12.92%  nvm_ensure_version_installed
 5)    2          50.12    25.06    1.42%     50.12    25.06    1.42%  compaudit
 6)    1          53.81    53.81    1.53%     46.70    46.70    1.33%  nvm_die_on_prefix
 7)    1          43.21    43.21    1.23%     43.21    43.21    1.23%  nvm_is_version_installed
 8)    1          78.15    78.15    2.22%     28.03    28.03    0.80%  compinit
 9)    1           6.56     6.56    0.19%      6.56     6.56    0.19%  nvm_grep
10)    4           7.10     1.78    0.20%      0.54     0.14    0.02%  nvm_npmrc_bad_news_bears
11)    1           0.32     0.32    0.01%      0.32     0.32    0.01%  add-zsh-hook
12)    1           0.29     0.29    0.01%      0.29     0.29    0.01%  nvm_has
13)    1           0.08     0.08    0.00%      0.08     0.08    0.00%  compdef
14)    1           0.15     0.15    0.00%      0.07     0.07    0.00%  complete
15)    1        2301.03  2301.03   65.37%      0.03     0.03    0.00%  nvm_process_parameters
16)    1           0.02     0.02    0.00%      0.02     0.02    0.00%  bashcompinit
17)    1           0.01     0.01    0.00%      0.01     0.01    0.00%  nvm_is_zsh

-----------------------------------------------------------------------------------

15)    1        2301.03  2301.03   65.37%      0.03     0.03    0.00%  nvm_process_parameters
       1/1      2301.00  2301.00   65.36%   1054.46  1054.46             nvm_auto [2]

-----------------------------------------------------------------------------------

       1/1      2301.00  2301.00   65.36%   1054.46  1054.46             nvm_process_parameters [15]
 2)    1        2301.00  2301.00   65.36%   1054.46  1054.46   29.95%  nvm_auto
       1/2      1246.54  1246.54   35.41%     25.76    25.76             nvm [3]

-----------------------------------------------------------------------------------

       1/2      1246.54  1246.54   35.41%     25.76    25.76             nvm_auto [2]
       1/2      1220.78  1220.78   34.68%    668.53   668.53             nvm [3]
 3)    2        1246.54   623.27   35.41%    694.29   347.15   19.72%  nvm
       1/1         0.29     0.29    0.01%      0.29     0.29             nvm_has [12]
       1/1        53.81    53.81    1.53%     46.70    46.70             nvm_die_on_prefix [6]
       1/1       498.15   498.15   14.15%    454.94   454.94             nvm_ensure_version_installed [4]
       1/2      1220.78  1220.78   34.68%    668.53   668.53             nvm [3]

-----------------------------------------------------------------------------------

 1)    1        1140.57  1140.57   32.40%   1140.57  1140.57   32.40%  load-nvmrc

-----------------------------------------------------------------------------------

       1/1       498.15   498.15   14.15%    454.94   454.94             nvm [3]
 4)    1         498.15   498.15   14.15%    454.94   454.94   12.92%  nvm_ensure_version_installed
       1/1        43.21    43.21    1.23%     43.21    43.21             nvm_is_version_installed [7]

-----------------------------------------------------------------------------------

 8)    1          78.15    78.15    2.22%     28.03    28.03    0.80%  compinit
       1/2        50.12    50.12    1.42%      0.74     0.74             compaudit [5]

-----------------------------------------------------------------------------------

       1/1        53.81    53.81    1.53%     46.70    46.70             nvm [3]
 6)    1          53.81    53.81    1.53%     46.70    46.70    1.33%  nvm_die_on_prefix
       4/4         7.10     1.78    0.20%      0.54     0.14             nvm_npmrc_bad_news_bears [10]

-----------------------------------------------------------------------------------

       1/2        50.12    50.12    1.42%      0.74     0.74             compinit [8]
       1/2        49.38    49.38    1.40%     49.38    49.38             compaudit [5]
 5)    2          50.12    25.06    1.42%     50.12    25.06    1.42%  compaudit
       1/2        49.38    49.38    1.40%     49.38    49.38             compaudit [5]

-----------------------------------------------------------------------------------

       1/1        43.21    43.21    1.23%     43.21    43.21             nvm_ensure_version_installed [4]
 7)    1          43.21    43.21    1.23%     43.21    43.21    1.23%  nvm_is_version_installed

-----------------------------------------------------------------------------------

       4/4         7.10     1.78    0.20%      0.54     0.14             nvm_die_on_prefix [6]
10)    4           7.10     1.78    0.20%      0.54     0.14    0.02%  nvm_npmrc_bad_news_bears
       1/1         6.56     6.56    0.19%      6.56     6.56             nvm_grep [9]

-----------------------------------------------------------------------------------

       1/1         6.56     6.56    0.19%      6.56     6.56             nvm_npmrc_bad_news_bears [10]
 9)    1           6.56     6.56    0.19%      6.56     6.56    0.19%  nvm_grep

-----------------------------------------------------------------------------------

11)    1           0.32     0.32    0.01%      0.32     0.32    0.01%  add-zsh-hook

-----------------------------------------------------------------------------------

       1/1         0.29     0.29    0.01%      0.29     0.29             nvm [3]
12)    1           0.29     0.29    0.01%      0.29     0.29    0.01%  nvm_has

-----------------------------------------------------------------------------------

14)    1           0.15     0.15    0.00%      0.07     0.07    0.00%  complete
       1/1         0.08     0.08    0.00%      0.08     0.08             compdef [13]

-----------------------------------------------------------------------------------

       1/1         0.08     0.08    0.00%      0.08     0.08             complete [14]
13)    1           0.08     0.08    0.00%      0.08     0.08    0.00%  compdef

-----------------------------------------------------------------------------------

16)    1           0.02     0.02    0.00%      0.02     0.02    0.00%  bashcompinit

-----------------------------------------------------------------------------------

17)    1           0.01     0.01    0.00%      0.01     0.01    0.00%  nvm_is_zsh
```

</details>

### With "fast nvm switcher"

The output is always on the order of ~10ms.

<details>

<summary>View zshprof output</summary>

```
num  calls                time                       self            name
-----------------------------------------------------------------------------------
 1)    1           7.77     7.77   94.53%      7.77     7.77   94.53%  load-nvmrc
 2)    1           0.45     0.45    5.47%      0.45     0.45    5.47%  add-zsh-hook

-----------------------------------------------------------------------------------

 1)    1           7.77     7.77   94.53%      7.77     7.77   94.53%  load-nvmrc

-----------------------------------------------------------------------------------

 2)    1           0.45     0.45    5.47%      0.45     0.45    5.47%  add-zsh-hook
```

</details>

## How to use it

Begin by installing [nvm](https://github.com/nvm-sh/nvm) using the download script from their README.

If you already have nvm installed, remove anything nvm-related you might already have from your local .zshrc file.

Next, run the following script to download the `resolve_node_version` binary to your machine:

```sh
(cd $HOME/.nvm/ && curl -L -O https://github.com/abejfehr/fast-nvm-switcher/releases/download/0.1.2/resolve_node_version)
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
  if [ -n "$NODE_PATH" ]; then
    echo "Updating node location to $NODE_PATH"
    PATH="$PATH:$NODE_PATH"
  fi
}

autoload -U add-zsh-hook
add-zsh-hook chpwd load-nvmrc
load-nvmrc
```

## Limitations

- This script only supports the 'default' alias, it won't work with lts, etc.
- The install script assumes that you're using the default `$NVM_DIR` of `$HOME/.nvm`
