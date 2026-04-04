# Suggested Commands

## lo/ and sync/ modules

```sh
cd lo/   # or cd sync/
make run  # go run cmd/main.go
```

## slog/ module

```sh
cd slog/
make build        # compile to ./build/main
make run          # run compiled binary
make up           # hot-reload dev server (air)

# Test endpoints
make get_health
make get_me
make get_profile_me
make get_product_by_id
make all_requests
```

## 1_24/testing_b_loop/ module

```sh
cd 1_24/testing_b_loop/
go test -bench=. -benchmem ./...   # 全ベンチマーク実行
go test -bench=BenchmarkNew -benchmem ./...  # b.Loop() のみ実行
go test -bench=BenchmarkOld -benchmem ./...  # 旧 b.N ループのみ実行
```

## Notes

- 1_24/ 配下はすべてテスト専用モジュール (main.go なし、\_test.go のみ)
- Hot-reload (air) configured only for slog/ via .air.toml
- Each module has its own go.mod; commands must be run from within each module directory
