# Unfoldcpp-CLI
[English version](readme.md)
## Get started with Unfoldcpp-CLI
Unfoldcpp-CLIは、`#include "hoge.h"`のようにファイルをパスで指定された`#include`指令について、これをインクルード元のコード中に再帰的に展開します。

下記の手順で導入することができます。

1. このリポジトリをクローンする
    ```bash
    git clone git@github.com:comavius/unfoldcpp-cli.git
    ```
1. golangのコードをビルドする
    ```bash
    cd /path/to/cloned/repo
    make
    ```
1. 環境変数を指定する
1. 適切にインストールされているか確認する
    ```bash
    unfoldcpp help
    ```

## 使い方
```bash    
unfoldcpp unfold /path/to/main.cpp > path/to/unfolded.cpp
```
