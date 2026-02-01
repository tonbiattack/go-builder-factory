# go-builder-factory

Goのテスト駆動開発向けに、ビルダーパターンとファクトリメソッドの使い分け、およびDB接続テストのテストデータ作成を再現できるサンプルです。

## 収録内容
- ビルダーパターンとファクトリメソッドのサンプル
- 実DB（MySQL）に接続するテスト
- トランザクションでロールバックし、テストデータを汚さない構成

## 事前準備
MySQL を起動し、テスト用DBを作成してください。

```sql
CREATE DATABASE test_db;
```

## 環境変数
テスト実行前にDSNを指定します。

```
export TEST_DSN="root:password@tcp(localhost:3306)/test_db?parseTime=true"
```

## テスト実行
```
go test ./...
```

## サンプルコード
- [sample/user.go](sample/user.go)
- [sample/repository.go](sample/repository.go)
- [sample/user_test.go](sample/user_test.go)
