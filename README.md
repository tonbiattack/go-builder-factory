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

## 注記
- このサンプルでは「テストが完璧に動作するかどうか」を確認する目的で、実行確認を行っています。
- フィールド変数の値を直接確認するテストは非常に壊れやすく、原則として導入しない方が良いと考えています。
- ただし本サンプルでは動作確認のために、あえてフィールド確認テストを含めています。
- ルールから外れる点がある場合は、サンプルであることを踏まえて明記しています。

## サンプルコード
- [sample/user.go](sample/user.go)
- [sample/repository.go](sample/repository.go)
- [sample/user_test.go](sample/user_test.go)
- テスト専用のビルダー/ファクトリ: [sample/test_helpers_test.go](sample/test_helpers_test.go)
