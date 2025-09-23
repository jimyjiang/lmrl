#!/bin/bash

# 设置输出目录
OUTPUT_DIR="./bin"

# 支持的平台列表
PLATFORMS=(
  "linux/amd64"
  "darwin/amd64"
)

# 清理并创建输出目录
rm -rf ${OUTPUT_DIR}
mkdir -p ${OUTPUT_DIR}

# 遍历所有平台进行编译
for PLATFORM in "${PLATFORMS[@]}"; do
  # 分割平台字符串
  OS="${PLATFORM%/*}"
  ARCH="${PLATFORM#*/}"
  
  # 创建平台特定目录
  BIN_DIR="${OUTPUT_DIR}/${OS}_${ARCH}"
  mkdir -p ${BIN_DIR}
  
  
  # 设置环境变量并编译
  echo "Building for ${OS}/${ARCH}..."
  env GOOS=${OS} GOARCH=${ARCH} go build -o "${BIN_DIR}" ./...
  
  # 检查是否编译成功
  if [ $? -ne 0 ]; then
    echo "Error building for ${OS}/${ARCH}"
    exit 1
  fi
  
  # 可选：复制配置文件等资源
  # cp config.json ${BIN_DIR}/
done

echo "Build completed! Binaries are in ${OUTPUT_DIR}"