#!/bin/bash
# Go 学習トピックの GitHub issue を作成する汎用スクリプト
#
# 使い方:
#   bash .claude/skills/create-learning-issue/scripts/create_issue.sh <version> <topic-number> <branch-suffix>
#
# 引数:
#   version       Go バージョン (例: 1.24, 1.25, 1.26)
#   topic-number  トピック番号  (例: 1, 2 ... 10)
#   branch-suffix ブランチ名のサフィックス (例: bytes-strings-iterator)
#
# 例:
#   bash .claude/skills/create-learning-issue/scripts/create_issue.sh 1.24 1 bytes-strings-iterator
#   bash .claude/skills/create-learning-issue/scripts/create_issue.sh 1.25 3 vet-waitgroup-hostport
set -euo pipefail

create_issue() {
  local version="${1:?第1引数にGoバージョンを指定してください (例: 1.24)}"
  local topic_num="${2:?第2引数にトピック番号を指定してください (例: 1)}"
  local branch_suffix="${3:?第3引数にブランチサフィックスを指定してください (例: bytes-strings-iterator)}"
  local dir_version="${version//./_}"
  local file="tmp/${dir_version}/learning_topics.md"

  if [[ ! -f "$file" ]]; then
    echo "Error: $file が見つかりません" >&2
    exit 1
  fi

  # タイトル抽出: "### N. タイトル" の行からタイトル部分のみ取得
  local title
  title=$(grep -m1 "^### ${topic_num}\. " "$file" | sed "s/^### ${topic_num}\. //")

  # トピックセクション抽出: "### N." から次の "### " までの行
  local section
  section=$(awk -v n="${topic_num}" '
    $0 ~ ("^### "n"\\. ") { found=1; next }
    found && /^### [0-9]+\./ { exit }
    found { print }
  ' "$file")

  # 概要: 最初の "####" または "検証する内容:" が現れるまでの段落
  local overview
  overview=$(printf '%s\n' "$section" | awk '
    /^####/ || /^検証する内容:/ || /^---$/ { exit }
    { print }
  ' | sed -e '/^[[:space:]]*$/{ N; /^\n[[:space:]]*$/d; }')

  # 変更内容: "####" 見出しからテーブルを含む部分 ("検証する内容:" の手前まで)
  local changes
  changes=$(printf '%s\n' "$section" | awk '
    /^####/ { found=1 }
    found && (/^検証する内容:/ || /^---$/) { exit }
    found { print }
  ')

  # 検証ポイント: "検証する内容:" 以降の箇条書き
  local verification
  verification=$(printf '%s\n' "$section" | awk '
    /^検証する内容:/ { found=1; next }
    found && /^---$/ { exit }
    found { print }
  ' | sed '/^[[:space:]]*$/d')

  local branch="feature/go-${version}/${branch_suffix}"

  # issue 本文を組み立て
  local body
  body="## 概要

${overview}"

  if [[ -n "$changes" ]]; then
    body+="

## 変更内容

${changes}"
  fi

  body+="

## 検証ポイント

${verification}

## ブランチ

\`${branch}\`

## 参照

- [Go ${version} Release Notes](https://go.dev/doc/go${version})
- \`tmp/${dir_version}/learning_topics.md\` トピック ${topic_num}"

  gh issue create \
    --title "[Go ${version}] ${title}" \
    --body "$body"
}

create_issue "$@"
