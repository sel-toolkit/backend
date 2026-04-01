#!/bin/bash

# 資料庫還原腳本
# 使用 Docker 容器內的 mysql 還原資料庫

# 設定變數
DB_NAME="sel"
DB_HOST="localhost"
DB_PORT="3306"
BACKUP_DIR="./backups"
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

# 檢查參數
if [ $# -eq 0 ]; then
    echo "用法: $0 <備份檔案路徑>"
    echo "例如: $0 ./backups/sel_backup_20260401_143022.sql"
    echo ""
    echo "可用的備份檔案："
    ls -la "$BACKUP_DIR"/*.sql* 2>/dev/null || echo "沒有找到備份檔案"
    exit 1
fi

BACKUP_FILE="$1"

# 檢查備份檔案是否存在
if [ ! -f "$BACKUP_FILE" ]; then
    echo "錯誤：備份檔案不存在: $BACKUP_FILE"
    exit 1
fi

echo "⚠️  警告：此操作將覆蓋現有的資料庫資料！"
echo "資料庫名稱: $DB_NAME"
echo "備份檔案: $BACKUP_FILE"
echo "使用容器: $CONTAINER_NAME"
echo ""
read -p "確定要還原資料庫嗎？(y/N): " -n 1 -r
echo
if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    echo "取消還原操作"
    exit 1
fi

echo "開始還原資料庫..."

# 檢查檔案是否為壓縮檔
if [[ "$BACKUP_FILE" == *.gz ]]; then
    echo "檢測到壓縮檔案，先解壓..."
    gunzip -c "$BACKUP_FILE" | docker exec -i "$CONTAINER_NAME" mysql \
        -h "$DB_HOST" \
        -P "$DB_PORT" \
        -u root \
        -p"$MYSQL_ROOT_PASSWORD" \
        "$DB_NAME"
else
    docker exec -i "$CONTAINER_NAME" mysql \
        -h "$DB_HOST" \
        -P "$DB_PORT" \
        -u root \
        -p"$MYSQL_ROOT_PASSWORD" \
        "$DB_NAME" < "$BACKUP_FILE"
fi

# 檢查還原是否成功
if [ $? -eq 0 ]; then
    echo "✅ 資料庫還原成功！"
else
    echo "❌ 資料庫還原失敗！"
    exit 1
fi

echo "還原完成！" 