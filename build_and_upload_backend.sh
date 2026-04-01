#!/bin/bash

# 設定變數
IMAGE_NAME="sel_backend"
OUTPUT_DIR="./tmp"
TAR_NAME="backend.tar"
REMOTE_USER="subarya"
REMOTE_HOST="ds.subarya.me"
REMOTE_PATH="/var/www/sel-backend"

# 建立 tmp 資料夾（如果不存在）
mkdir -p "$OUTPUT_DIR"

# 移除舊的 tar 檔
rm -f "$OUTPUT_DIR/$TAR_NAME"

# 建立 Docker 映像
echo "🔨 Building Docker image..."
docker build --platform linux/amd64 -t "$IMAGE_NAME" .

# 匯出 Docker 映像為 .tar 檔
echo "📦 Saving Docker image to $OUTPUT_DIR/$TAR_NAME..."
docker save "$IMAGE_NAME" -o "$OUTPUT_DIR/$TAR_NAME"

# 上傳到遠端伺服器
echo "🚀 Uploading to $REMOTE_HOST:$REMOTE_PATH..."
scp "$OUTPUT_DIR/$TAR_NAME" "$REMOTE_USER@$REMOTE_HOST:$REMOTE_PATH"

echo "✅ 完成！"
