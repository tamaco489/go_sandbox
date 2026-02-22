# issue テンプレート

## タイトル

```text
[Go <version>] <トピックのタイトル (### の見出し文字列から番号を除いた部分)>
```

## ブランチ名

```text
feature/go-<version>/<トピックのタイトルを英単語のケバブケースに変換したもの>
```

例: `feature/go-1.24/bytes-strings-iterator`

## 本文

```markdown
## 概要

<トピックの冒頭の説明文をそのまま使う>

## 変更内容

<トピック内の表をそのまま使う。表がない場合はこのセクションを省略する>

## 検証ポイント

<「検証する内容:」以下の箇条書きをそのまま使う>

## ブランチ

`feature/go-<version>/<トピックのタイトルを英単語のケバブケースに変換したもの>`

## 参照

- [Go <version> Release Notes](https://go.dev/doc/go<version>)
- `tmp/<version>/learning_topics.md` トピック <番号>
```
