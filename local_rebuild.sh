#!/bin/bash

# 停止並移除容器（即使沒在跑也不會出錯）
docker-compose -p sel_toolkit down || true

# 清除不用的資源
docker system prune -f

echo "啟動 MySQL 和 Redis 服務..."

# 只啟動 MySQL 和 Redis，不啟動 backend
docker-compose -p sel_toolkit up -d sel_mysql sel_redis

# 啟動 gf run main.go
gf run main.go
