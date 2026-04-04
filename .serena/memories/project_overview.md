# Project Overview: go_sandbox

## Purpose

Personal sandbox repository for verifying the behavior of Go standard library and third-party libraries. Each subdirectory is an independent Go module.

## Tech Stack

- Language: Go
- Go versions: see each module's go.mod
- External libraries: github.com/samber/lo (lo/), github.com/google/uuid (slog/)

## Modules

### lo/

- Demonstrates usage of the samber/lo utility library
- Key features: GroupBy, Map, Reduce, SliceToMap, UniqMap, FilterMap
- Models: User, Player, Pokemon, Achievement, Item

### sync/

- Demonstrates Go's sync package (stdlib only)
- Key features: Mutex, RWMutex, Once

### slog/

- HTTP server demonstrating Go 1.21+ structured logging (log/slog)
- Uses pure stdlib net/http (no frameworks)
- Routes: /api/v1/health, /api/v1/products/{id}, /api/v1/users/me, /api/v1/users/profile/me
- Middleware: auth, logging
- Logging levels by status: 5xx→error, 4xx→warn, 2xx→info

### 1_24/

- Go 1.24 新機能の動作検証モジュール群 (各サブディレクトリが独立した Go モジュール)
- 1_24/testing_b_loop/: testing.(*B).Loop() の検証
  - bench_strings_test.go: 旧 b.N ループ vs 新 b.Loop() の文字列処理ベンチマーク
  - bench_file_test.go: setup/cleanup を伴うファイル操作ベンチマーク

### 1_25/ / 1_26/

- Go 1.25 / 1.26 新機能検証用ディレクトリ (現時点では .gitkeep のみ、未着手)

## tmp/1_24/

- Go 1.24 学習・ブログ執筆用の作業ディレクトリ
- learning_topics.md: 優先度付き学習トピック 10 件
- release_notes.md: リリースノートのメモ
- blog/testing_b_loop/: testing.(*B).Loop() のブログ素案 (draft.md, blog.md)
- blog/bytes_strings_iterator/: bytes/strings イテレータのブログ素案 (plan.md、執筆中)
