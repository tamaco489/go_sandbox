---
name: create-learning-issue
description: Go の学習トピックを GitHub issue として登録する。tmp/<version>/learning_topics.md のトピックを issue 化するときに使う。
argument-hint: "[go-version] [topic-number]"
---

`$ARGUMENTS` から Go バージョンとトピック番号を読み取り、GitHub issue を作成する。

## 引数の形式

```
/create-learning-issue 1.24 1
```

- 第 1 引数: Go バージョン (例: `1.24`, `1.25`, `1.26`)
- 第 2 引数: トピック番号 (例: `1`, `2` ... `10`)

## 手順

1. `tmp/<version>/learning_topics.md` を読み込み、指定トピックの内容を取得する
2. `assets/issue_template.md` を参照し、タイトルと本文を組み立てる
3. `gh issue create` で issue を登録する
4. 作成された issue の URL を出力する

## テンプレート参照

issue のタイトル・本文の形式は `assets/issue_template.md` を参照すること。

## 注意事項

- `gh` コマンドが使えることを前提とする
- 表がないトピックは「変更内容」セクションを省略する
