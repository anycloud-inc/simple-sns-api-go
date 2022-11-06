## Test

```sh
# Install gotest
go install github.com/rakyll/gotest

# Run test
gotest -v ./...
```

## DB スキーマの変更

```sh
# Entityの生成
go run -mod=mod entgo.io/ent/cmd/ent init [EntityName]

# スキーマの変更
vim ./ent/schema/[entity_name].go

# ファイルの自動生成
go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/execquery ./ent/schema
```
