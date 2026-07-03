# 配置目录

本地开发配置位于 `etc/dev/`，详细启动说明见 [docs/local-dev.md](../docs/local-dev.md)。

## 服务端口

| 模块名称     | REST端口 | RPC端口 | Prometheus |
|----------|--------|-------|------------|
| auth     | 8085   | -     | 4001       |
| system   | 8086   | 9092  | 4002       |
| demo     | 8099   | 9099  | 4009       |

## 新建模块
```shell
goctl api new demo --style go_zero
```