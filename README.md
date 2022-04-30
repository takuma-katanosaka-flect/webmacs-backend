# webmacs-backend

## 開発環境

- Editor
  - VSCode
- Infrastructure
  - Heroku

## セットアップ

### anyenv + goenvのインストール

参考文献を使って設定を行う。

1. [anyenv + macOS環境構築](https://qiita.com/rinpa/items/81766cd6a7b23dea9f3c)
1. [anyenvのgoenvにGoLangをインストール](https://qiita.com/suzuki_cs/items/62799b5035e130818486)

### VSCode Extensionのインストール

`golang.go`をインストールする。

### GOPATH、GOROOTの設定

1. 以下のコマンドを実行する

    ```sh
    go env GOROOT
    go env GOPATH
    ```

1. setting.jsonにコマンドで取得したパスをそれぞれ設定する。

    ```json
    "go.gopath": "go env GOPATHで取得したパスを設定"
    "go.goroot": "go env GOROOTで取得したパスを設定",
    ```

## ローカルサーバ起動

TO BE DEFINED

## デプロイ方法

TO BE DEFINED
