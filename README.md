# backend
Backend service for SEL Toolkit

## Requirements

- Go 1.23+
- Docker & Docker Compose
- [golang-migrate CLI](https://github.com/golang-migrate/migrate) (`brew install golang-migrate`)

## Setup

複製環境變數範本並填入設定：

```bash
cp .env.example .env
```

啟動所有服務（MySQL、Redis、Backend）：

```bash
docker compose up -d
```

Migration 會在 Backend 啟動前自動執行。

## Database Migration

新增 migration：

```bash
migrate create -ext sql -dir ./manifest/sql -seq <migration_name>
```

會產生 `00000X_<migration_name>.up.sql` 和 `.down.sql`，填入 SQL 內容後重新部署即可。

手動執行（本機開發）：

```bash
# 先 build binary
go build -o main .

# 套用所有未執行的 migration
./main -migrate up

# rollback 最後一個 migration
./main -migrate down
```

執行成功後可確認 DB 狀態：

```bash
docker exec -it sel_toolkit-sel_mysql-1 mysql -uroot -p"$MYSQL_ROOT_PASSWORD" sel -e "SELECT * FROM schema_migrations;"
```

`version` 為最新版號、`dirty=0` 代表執行成功。

### 正式區部署

```bash
docker compose up
```

`sel_migrate` service 會在 MySQL ready 後自動執行 migration，成功後才啟動 `sel_backend`。

## Development

```bash
# 產生 DAO / DO / Entity
make dao

# 產生 Controller
make ctrl

# 產生 Service
make service
```
