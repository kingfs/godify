#!/bin/bash

# Dify Golang SDK 测试覆盖率脚本

echo "Running Dify Golang SDK tests with coverage..."

# 创建coverage目录
mkdir -p coverage

# 运行所有包的测试并生成覆盖率报告
go test -v -coverprofile=coverage/coverage.out ./...

# 如果测试失败，退出
if [ $? -ne 0 ]; then
    echo "Tests failed!"
    exit 1
fi

# 生成HTML覆盖率报告
go tool cover -html=coverage/coverage.out -o coverage/coverage.html

# 显示总体覆盖率
echo ""
echo "=== Test Coverage Summary ==="
go tool cover -func=coverage/coverage.out | tail -1

echo ""
echo "Coverage report generated: coverage/coverage.html"
echo "Open coverage/coverage.html in your browser to view detailed coverage."