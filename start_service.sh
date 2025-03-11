#!/bin/bash

# 设置Gomall根目录
GOMALL_ROOT="."
cd $GOMALL_ROOT

# 设置颜色输出
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${GREEN}正在启动Gomall微服务...${NC}"

# 确保Docker容器已启动
echo -e "${YELLOW}启动Docker容器...${NC}"
docker-compose up -d

# 等待Docker容器完全启动
echo -e "${YELLOW}等待Docker容器完全启动...${NC}"

# 检测所有容器是否都已启动并运行
wait_for_containers() {
  local max_attempts=30
  local attempt=1
  local all_running=false
  
  while [ $attempt -le $max_attempts ]; do
    echo -e "${YELLOW}检查容器状态 (尝试 $attempt/$max_attempts)...${NC}"
    
    # 获取所有容器状态
    local containers=$(docker-compose ps -q)
    all_running=true
    
    for container in $containers; do
      local status=$(docker inspect --format='{{.State.Status}}' $container)
      if [ "$status" != "running" ]; then
        all_running=false
        echo -e "${YELLOW}容器 $container 状态: $status, 继续等待...${NC}"
        break
      fi
    done
    
    if $all_running; then
      echo -e "${GREEN}所有容器已成功启动!${NC}"
      return 0
    fi
    
    sleep 2
    ((attempt++))
  done
  
  echo -e "${YELLOW}警告: 等待容器启动超时，将继续执行后续步骤...${NC}"
  return 1
}

wait_for_containers

# 定义服务列表
services=("user" "product" "cart" "order" "checkout" "payment" "email")

# 创建日志目录
mkdir -p $GOMALL_ROOT/logs

# 检查终端类型并使用适当的命令
if command -v gnome-terminal &> /dev/null; then
    # 使用gnome-terminal
    for service in "${services[@]}"; do
        echo -e "${YELLOW}启动 ${service} 服务...${NC}"
        gnome-terminal --title="${service} Service" -- bash -c "cd $GOMALL_ROOT/app/${service} && go mod tidy && go run .; exec bash"
        sleep 1
    done
elif command -v konsole &> /dev/null; then
    # 使用KDE的konsole
    for service in "${services[@]}"; do
        echo -e "${YELLOW}启动 ${service} 服务...${NC}"
        konsole --new-tab --title "${service} Service" -e bash -c "cd $GOMALL_ROOT/app/${service} && go mod tidy && go run .; exec bash" &
        sleep 1
    done
elif command -v xterm &> /dev/null; then
    # 使用xterm
    for service in "${services[@]}"; do
        echo -e "${YELLOW}启动 ${service} 服务...${NC}"
        xterm -title "${service} Service" -e "cd $GOMALL_ROOT/app/${service} && go mod tidy && go run .; exec bash" &
        sleep 1
    done
elif command -v tmux &> /dev/null; then
    # 使用tmux
    echo -e "${YELLOW}使用tmux启动所有服务...${NC}"
    tmux new-session -d -s gomall
    
    # 创建第一个窗口用于user服务
    tmux rename-window -t gomall:0 "user"
    tmux send-keys -t gomall:0 "cd $GOMALL_ROOT/app/user && go mod tidy && go run ." C-m
    
    # 为其他服务创建新窗口
    for i in {1..6}; do
        service=${services[$i]}
        tmux new-window -t gomall:$i -n "$service"
        tmux send-keys -t gomall:$i "cd $GOMALL_ROOT/app/$service && go mod tidy && go run ." C-m
    done
    
    echo -e "${GREEN}所有服务已在tmux会话中启动！${NC}"
    echo -e "${YELLOW}使用 'tmux attach -t gomall' 命令查看服务状态${NC}"
    echo -e "${YELLOW}在tmux中，使用Ctrl+b然后按数字键0-6切换不同服务窗口${NC}"
else
    # 如果没有图形终端，则使用后台进程
    echo -e "${YELLOW}未检测到支持的终端模拟器，将在后台启动服务...${NC}"
    for service in "${services[@]}"; do
        echo -e "${YELLOW}启动 ${service} 服务...${NC}"
        cd $GOMALL_ROOT/app/${service}
        go mod tidy
        nohup go run . > $GOMALL_ROOT/logs/${service}.log 2>&1 &
        echo "$service 服务已在后台启动，日志保存在 $GOMALL_ROOT/logs/${service}.log"
        sleep 1
    done
fi

echo -e "${GREEN}所有服务启动命令已发送！${NC}"
echo -e "${GREEN}请在各个终端窗口中查看服务启动状态。${NC}"
echo -e "${YELLOW}您也可以通过访问 http://localhost:8500 查看Consul面板来确认服务状态${NC}"