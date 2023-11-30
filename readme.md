# Unfoldcpp-CLI

## Get started with Unfoldcpp-CLI
This is a cli-application providing function like `g++ -E`. However, unfoldcpp only gathers included files with quotation marks like `#include "hoge.h"`.

If you want to use, please follow this instructions.

1. Clone this repository
    ```bash
    git clone git@github.com:comavius/unfoldcpp-cli.git
    ```
1. Build golang project
    ```bash
    cd /path/to/cloned/repo
    make
    ```
1. Add `/path/to/here` to `PATH` or `mv` executable to directory you like.
1. Check if unfoldcpp is installed successfully.
    ```bash
    unfoldcpp help
    ```

## Usage
```bash    
unfoldcpp unfold /path/to/main.cpp > path/to/unfolded.cpp
```
