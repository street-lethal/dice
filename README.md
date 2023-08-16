# dice

## 準備

```shell
cp config/settings.sample.json config/settings.json
cp data/dice.sample.txt data/dice.txt
```

## 設定

```shell
vi config/settings.txt
```
* `base` : 進数(サイコロの面の数)
* `length` : ランダム文字列の文字数
* `char_kinds` : ランダム文字列の文字種
    * `A` : 英大文字 (`A` を指定するだけで A-Z が全て加わる)
    * `a` : 英小文字 (`a` を指定するだけで a-z が全て加わる)
    * `0` : 数字 (`0` を指定するだけで 0-9 が全て加わる)
    * `#` `$` `%` `&` `_` `-` : 記号(指定された文字のみ加わる)

## サイコロの出目を書き込む

```shell
vi data/dice.txt
```

## ランダム文字列生成

### バイナリー直接実行
```shell
./dice
```

### ソースから実行

```shell
goenv exec go run main.go
```
