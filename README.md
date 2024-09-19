# Hello Shiron / Solufit GO Template

## ABOUT

このテンプレートはSolufit / Shiron間で共通のベースリポジトリを用いて作成されています。  
各Orgで変更を加える場合はこのREADMEファイルの末尾に説明を加えるか、変更を加えた箇所の説明文章を取り消し線で消し追記してください。

Created by sasanqua and solufit developers

## 環境

- [6000](http://localhost:6000) Go API
    - Gin
    - xo
    - dbmate
- [6080](http://localhost:6080) Swagger API
- Maria DB

## ディレクトリ構成

```
|- .devcontainer
|- devcontainer.json                # devcontainerの設定
    |- docker-compose.yml           # 開発コンテナ関連のコンテナ立ち上げ
    |- Dockerfile                   # Go APIコンテナのビルド
|- .github
    |- dependabot.yml
    |- workflows
        |- actions.yml              # Lint / Generate / Build / Test
|- .vscode
    |- settings.json                # vscode 設定
|- batch
    |- main.go                      # バッチ処理用
|- db
    |- migrations
        |- xxxxxxxxxxxxxx_xxxxx.sql # マイグレートデータ
    |- seed
        |- 00000000000000_init.sql  # 初期データ
|- docs                             # swagger関連の自動生成ファイル
    |- docs.go
    |- swagger.json
    |- swagger.yaml
|- internal                         # メインのコード記述部分
    |- middleware
    |- pkg
    |- presenter
        |- server.go                # エンドポイント定義
    |- repository                   # db関連処理（CRUDはmodelsでやる）
    |- service                      # エンドポイントごとの処理
        |- version
            |- handler.go
|- models                           # xoが生成したDBのスキーマに基づくCRUDデータ
|- .gitignore
|- go.mod                           # Goのパッケージデータ
|- go.sun
|- main.go                          # サーバー起動エントリーポイント
|- README
```

## API

### エンドポイントの追加

1. `internal/service`配下にフォルダーを追加する
2. 作成したフォルダーは以下に`handler.go`を作成する
3. `persenter/server.go`にエンドポイントに関する記述を追記する

### ドキュメントの作成

```bash
swag init ./main.go
```
```
2024/09/19 04:19:51 Generate swagger docs....
2024/09/19 04:19:51 Generate general API Info, search dir:./
2024/09/19 04:19:52 Generating version.Response
2024/09/19 04:19:52 create docs.go at docs/docs.go
2024/09/19 04:19:52 create swagger.json at docs/swagger.json
2024/09/19 04:19:52 create swagger.yaml at docs/swagger.yaml
```

### テストの実行

```bash
go test -cover ./... -coverprofile=cover.out
go tool cover -html=cover.out -o cover.html
open cover.html
```

## DB

### テーブルの作成

#### sqlファイルの作成

```bash
dbmate new
```


#### sqlファイルの適用

```bash
dbmate --url mysql://root:root@db:3306/test -s db/schema.sql up 
```

### 初期データの投入

#### db/seedフォルダー配下に.sqlファイルを配置します

db/seed/sample.sql
```sql
-- migrate:up
INSERT INTO users (id, name, email, created_At, updated_at)
VALUES ('01F8MECHZX3TBDSZ7XY8E9ZHDH', '須藤史郎', 'sudo.shiron@solufit.net', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)

-- migrate:down
```


### マイグレート状態の確認

```bash
dbmate --url mysql://root:root@db:3306/test status
```

```
^[[3~[X] 00000000000000_init.sql
[X] 20240919030921_init.sql

Applied: 2
Pending: 0
```


### マイグレーションの適用

```bash
dbmate --url mysql://root:root@db:3306/test -d db/seed -s db/schema.sql up
```


### スキーマファイルの作成

```bash
xo schema mysql://root:root@db:3306/test
```
