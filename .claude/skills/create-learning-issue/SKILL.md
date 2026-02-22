---
name: create-learning-issue
description: Go の学習トピックを GitHub issue として登録する。tmp/<version>/learning_topics.md のトピックを issue 化するときに使う。
argument-hint: "[go-version] [topic-number(s)]"
---

`$ARGUMENTS` から Go バージョンとトピック番号を読み取り、GitHub issue を作成する。

## 引数の形式

```
/create-learning-issue 1.24 1
/create-learning-issue 1.25 1から10まで
```

- 第 1 引数: Go バージョン (例: `1.24`, `1.25`, `1.26`)
- 第 2 引数: トピック番号、またはトピック範囲 (例: `1`, `3から7まで`, `1から10まで`)

## 手順

1. `tmp/<version>/learning_topics.md` を読み込み、対象トピックの内容を把握する
2. 各トピックのタイトルから英単語のケバブケースでブランチサフィックスを決定する
3. 下記のスクリプトを実行して issue を登録する
4. 作成された issue の URL を出力する

## スクリプト実行

issue の作成には必ず以下のスクリプトを使用すること。

```bash
bash .claude/skills/create-learning-issue/scripts/create_issue.sh <version> <topic-number> <branch-suffix>
```

### 引数

| 引数            | 説明                                                 | 例                       |
| --------------- | ---------------------------------------------------- | ------------------------ |
| `version`       | Go バージョン                                        | `1.24`                   |
| `topic-number`  | トピック番号                                         | `1`                      |
| `branch-suffix` | タイトルを英単語のケバブケースに変換したサフィックス | `bytes-strings-iterator` |

### 実行例

```bash
bash .claude/skills/create-learning-issue/scripts/create_issue.sh 1.24 1 bytes-strings-iterator
bash .claude/skills/create-learning-issue/scripts/create_issue.sh 1.25 3 vet-waitgroup-hostport
bash .claude/skills/create-learning-issue/scripts/create_issue.sh 1.26 5 some-feature-name
```

## スクリプトの動作

`scripts/create_issue.sh` は以下を自動で行う。

- `tmp/<version>/learning_topics.md` からトピック番号に対応するセクションを抽出する
- 概要・変更内容 (表がある場合のみ)・検証ポイントを解析して issue 本文を組み立てる
- `gh issue create` で issue を登録する

## 注意事項

- `gh` コマンドが使えることを前提とする
- ブランチサフィックスは Claude がトピックタイトルを読んで英語ケバブケースで決定する
- 表がないトピックは「変更内容」セクションが自動的に省略される
