#!/bin/bash

# 停止並移除容器（即使沒在跑也不會出錯）
docker-compose -p sel_toolkit down || true

# 清除不用的資源
docker system prune -f

# 移除舊的 image（不存在也不會出錯）
docker rmi sel_backend || true

# 重新 build image
docker build -t sel_backend .

# 啟動 container
docker-compose -p sel_toolkit up -d
