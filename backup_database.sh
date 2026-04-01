#!/bin/bash

# 資料庫備份腳本
# 使用 Docker 容器內的 mysqldump 備份 MySQL 資料庫

# 設定變數
DB_NAME="sel"
DB_HOST="localhost"
DB_PORT="3306"
BACKUP_DIR="./backups"
TIMESTAMP=$(date +"%Y%m%d_%H%M%S")
BACKUP_FILE="${BACKUP_DIR}/sel_backup_${TIMESTAMP}.sql"
CONTAINER_NAME="backend-mysql-1"  # Docker Compose 預設容器名稱

# 檢查環境變數
if [ -z "$MYSQL_ROOT_PASSWORD" ]; then
    echo "錯誤：請先設定 MYSQL_ROOT_PASSWORD 環境變數"
    echo "例如：export MYSQL_ROOT_PASSWORD=your_password"
    exit 1
fi

# 檢查 Docker 容器是否運行
if ! docker ps | grep -q "$CONTAINER_NAME"; then
    echo "錯誤：MySQL 容器未運行"
    echo "請先啟動 Docker Compose：docker-compose up -d mysql"
    exit 1
fi

# 建立備份目錄
mkdir -p "$BACKUP_DIR"

echo "開始備份資料庫..."
echo "資料庫名稱: $DB_NAME"
echo "備份檔案: $BACKUP_FILE"
echo "使用容器: $CONTAINER_NAME"

# 執行備份（使用 Docker 容器內的 mysqldump）
docker exec "$CONTAINER_NAME" mysqldump \
    -h "$DB_HOST" \
    -P "$DB_PORT" \
    -u root \
    -p"$MYSQL_ROOT_PASSWORD" \
    --single-transaction \
    --routines \
    --triggers \
    --events \
    --add-drop-database \
    --create-options \
    "$DB_NAME" > "$BACKUP_FILE"

# 檢查備份是否成功
if [ $? -eq 0 ]; then
    echo "✅ 資料庫備份成功！"
    echo "備份檔案位置: $BACKUP_FILE"
    echo "檔案大小: $(du -h "$BACKUP_FILE" | cut -f1)"
else
    echo "❌ 資料庫備份失敗！"
    exit 1
fi

# 建立壓縮備份（可選）
echo "建立壓縮備份..."
gzip "$BACKUP_FILE"
COMPRESSED_FILE="${BACKUP_FILE}.gz"
echo "壓縮備份檔案: $COMPRESSED_FILE"
echo "壓縮後檔案大小: $(du -h "$COMPRESSED_FILE" | cut -f1)"

echo "備份完成！" 