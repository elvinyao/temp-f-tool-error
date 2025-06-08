# Focalboard Tool

## 系统功能介绍

Focalboard Tool 是一个用于管理和操作 Focalboard 看板的工具服务。该系统提供了一套完整的RESTful API，允许用户通过HTTP请求来获取、修改和管理Focalboard看板数据。本系统主要作为Focalboard和Mattermost的API集成工具，简化了与这些平台交互的复杂性。

### 主要功能

- **看板管理**：获取单个看板信息和配置
- **卡片操作**：获取、移动、批量操作卡片
- **状态管理**：按状态筛选卡片，批量移动卡片状态
- **用户状态**：更新用户自定义状态
- **关联查询**：获取卡片关联的Lead信息
- **完整的错误处理和日志记录**
- **Swagger API文档**
- **Basic Auth认证**

## 系统架构图

### 系统功能架构图

```
┌─────────────────────┐
│    客户端应用       │
└─────────┬───────────┘
          │
          ▼
┌─────────────────────┐      ┌─────────────────────┐
│  Focalboard Tool    │◄────►│    配置文件         │
│  RESTful API服务    │      │  (application.toml) │
└─────────┬───────────┘      └─────────────────────┘
          │
          ▼
┌─────────────────────┐      ┌─────────────────────┐
│  服务层 (Service)   │◄────►│    错误定义         │
└─────────┬───────────┘      │   (errors.yaml)     │
          │                  └─────────────────────┘
          ▼
┌─────────────────────┐
│  数据访问层 (DAO)   │
└─────────┬───────────┘
          │
          ▼
┌─────────────────────┐      ┌─────────────────────┐
│  HTTP API客户端     │◄────►│   Focalboard        │
└─────────────────────┘      └─────────────────────┘
          │
          ▼
┌─────────────────────┐
│    Mattermost       │
└─────────────────────┘
```

### 业务架构图

```
┌───────────────────────────────────────────┐
│              客户端应用                    │
└───────────────────┬───────────────────────┘
                    │
                    ▼
┌───────────────────────────────────────────┐
│              认证层 (Basic Auth)           │
└───────────────────┬───────────────────────┘
                    │
                    ▼
┌───────────────────────────────────────────┐
│              API层                        │
│   ┌─────────────────────────────────┐     │
│   │      /api/v1/focalboard/*       │     │
│   └─────────────────┬───────────────┘     │
└───────────────────┬─┴───────────────────┘
                    │
                    ▼
┌───────────────────────────────────────────┐
│              业务逻辑层                    │
│   ┌─────────────────────────────────┐     │
│   │    配置验证       │    看板配置  │     │
│   └─────────────────────────────────┘     │
│   ┌─────────────────────────────────┐     │
│   │    错误处理       │    日志记录  │     │
│   └─────────────────────────────────┘     │
└───────────────────┬───────────────────────┘
                    │
                    ▼
┌───────────────────────────────────────────┐
│              数据访问层                    │
│   ┌─────────────────────────────────┐     │
│   │ Focalboard API │ Mattermost API │     │
│   └─────────────────────────────────┘     │
└───────────────────────────────────────────┘
```

### 技术架构图

```
┌───────────────────────────────────────────┐
│              应用层                        │
│   ┌─────────────────────────────────┐     │
│   │         Gin Web框架             │     │
│   └─────────────────────────────────┘     │
└───────────────────┬───────────────────────┘
                    │
                    ▼
┌───────────────────────────────────────────┐
│              中间件层                      │
│   ┌─────────────────────────────────┐     │
│   │ 请求ID │ 错误处理 │ 访问控制     │     │
│   └─────────────────────────────────┘     │
└───────────────────┬───────────────────────┘
                    │
                    ▼
┌───────────────────────────────────────────┐
│              核心服务层                    │
│   ┌─────────────────────────────────┐     │
│   │ 配置管理 │ 日志服务 │ 错误管理   │     │
│   └─────────────────────────────────┘     │
└───────────────────┬───────────────────────┘
                    │
                    ▼
┌───────────────────────────────────────────┐
│              工具库层                      │
│   ┌─────────────────────────────────┐     │
│   │ HTTP客户端 │ 日志库 │ 错误处理库 │     │
│   └─────────────────────────────────┘     │
└───────────────────────────────────────────┘
```

## 技术栈

- **Go** 1.24.3
- **Gin** Web框架
- **Swagger** API文档
- **Zap** 高性能日志库
- **Lumberjack** 日志轮转
- **Focalboard** 官方SDK集成
- **Mattermost** 官方SDK集成

## 错误处理

系统实现了一套完整的错误处理机制，分为两种类型的错误：

1. **客户端错误(client_error)**: 由客户端引起的错误，如参数无效、资源不存在等。
2. **服务器错误(server_error)**: 系统内部错误，如数据库连接失败、API调用失败等。

错误配置集中在 `configs/errors.yaml` 文件中定义，包含错误代码、HTTP状态码、错误消息模板等。

错误处理流程：

1. 使用中间件捕获所有API请求处理过程中产生的错误
2. 将错误转换为标准的错误响应格式
3. 根据错误类型记录不同级别的日志
4. 对于服务器错误，隐藏敏感的错误详情，只返回基本错误信息
5. 在响应中包含请求ID，方便问题追踪

## 日志处理

系统使用 Uber 的 zap 日志库实现高性能日志记录，主要功能：

1. 支持多种日志级别(DEBUG, INFO, WARN, ERROR, FATAL, PANIC)
2. 支持日志轮转(log rotation)，通过 lumberjack 实现
3. 日志内容包含时间戳、日志级别、请求ID、错误详情等信息
4. 可通过配置文件调整日志行为，如文件大小、保留数量、保留天数等

日志配置位于 `application.toml` 的 `[log]` 部分：

```toml
[log]
logLevel = "debug"
useLogRotation = true
[log.logProps]
fileName = "./logs/aa.log"
maxSize = 1
maxBackups = 5
maxAge = 31
```

## API接口文档

系统集成了Swagger文档，可通过访问 `/swagger/*` 路径查看完整API文档。

### 认证

API使用Basic Auth认证，默认用户名和密码：
- 用户名：`admin` / 密码：`admin`
- 用户名：`user` / 密码：`user`

### 主要API接口

#### 1. 获取单个看板信息

```http
GET /api/v1/focalboard/boards/single
```

**参数:**
- `boardId`: 看板ID (必填)
- `token`: 用户认证令牌 (必填)

**响应示例:**
```json
{
  "success": true,
  "data": {
    "id": "board-id",
    "teamId": "team-id",
    "title": "看板标题",
    "cardProperties": [...]
  }
}
```

#### 2. 获取看板所有卡片信息

```http
GET /api/v1/focalboard/boards/cards/allinfo
```

**参数:**
- `boardId`: 看板ID (必填)

**描述:** 获取指定看板中的所有卡片，包括属性和值

#### 3. 获取看板卡片列表

```http
GET /api/v1/focalboard/boards/cards/list/all
```

**参数:**
- `boardId`: 看板ID (必填)

**描述:** 获取看板中所有卡片的列表

#### 4. 按状态筛选卡片

```http
GET /api/v1/focalboard/boards/cards/list/filter/status
```

**参数:**
- `boardId`: 看板ID (必填)
- `statusName`: 状态名称 (可选)

**描述:** 根据指定状态筛选看板中的卡片

#### 5. 移动单个卡片

```http
PUT /api/v1/focalboard/boards/cards/move/one
```

**参数:**
- `boardId`: 看板ID (必填)
- `asleadId`: Aslead ID (必填)
- `statusName`: 目标状态名称 (必填)

**描述:** 将单个卡片移动到不同的状态列

#### 6. 批量移动卡片

```http
PUT /api/v1/focalboard/boards/cards/move/batch
```

**参数:**
- `boardId`: 看板ID (必填)
- `fromStatusName`: 源状态名称 (必填)
- `toStatusName`: 目标状态名称 (必填)

**描述:** 批量移动卡片到不同的状态列

#### 7. 更新用户自定义状态

```http
PATCH /api/v1/focalboard/boards/one/status/patch
```

**参数:**
- `asleadId`: Aslead ID (必填)
- `statusName`: 状态名称 (必填)
- `userToken`: 用户令牌 (必填)

**描述:** 为用户设置新的自定义状态

#### 8. 获取卡片关联信息

```http
GET /api/v1/focalboard/cards/single/asleadinfo
```

**参数:**
- `boardId`: 看板ID (必填)
- `cardId`: 卡片ID (必填)

**描述:** 使用卡片ID和看板ID获取关联的Lead信息

### 错误响应格式

```json
{
  "success": false,
  "error": {
    "code": "错误代码",
    "message": "错误信息",
    "params": {
      "request_id": "请求ID",
      "timestamp": 1234567890
    }
  }
}
```

## 快速开始

### 安装和运行

1. **克隆项目**
```bash
git clone <repository-url>
cd focalboard-tool
```

2. **安装依赖**
```bash
go mod download
```

3. **配置设置**
   - 复制 `configs/application.toml` 并根据需要修改配置
   - 确保Focalboard和Mattermost服务器地址正确

4. **运行服务**
```bash
go run main.go
```

5. **访问API文档**
   - 打开浏览器访问 `http://localhost:8080/swagger/index.html`

### 配置文件

主要配置在 `configs/application.toml` 中：

```toml
[app]
httpPort = 8080
version = "0.0.1"
appName = "focalboard-tool"
runMode = "debug"

[httpClient.focalboardClient]
addr = "http://your-focalboard-server"
apiVersionPath = "/api/v2"
timeout = "10s"

[httpClient.mattermostClient]
addr = "http://your-mattermost-server"
apiVersionPath = "/api/v4"
timeout = "10s"
```

## 开发指南

### 项目结构

```
├── configs/            # 配置文件目录
│   ├── application.toml # 应用配置
│   └── errors.yaml     # 错误定义
├── docs/              # Swagger文档
├── internal/          # 内部代码
│   ├── apimodel/      # API模型定义
│   ├── appconst/      # 应用常量
│   ├── conf/          # 配置处理
│   ├── dao/           # 数据访问层
│   ├── middleware/    # HTTP中间件
│   ├── model/         # 内部模型
│   ├── server/        # HTTP服务器
│   └── service/       # 业务逻辑
├── library/           # 工具库
├── logs/              # 日志文件
├── pkg/               # 公共包
├── main.go            # 入口文件
├── go.mod             # Go模块定义
└── go.sum             # 依赖版本锁定
```

### 添加新API

1. 在 `internal/server/handler` 中创建新的处理函数
2. 在 `internal/server/http/focalboard.go` 中添加路由
3. 在 `internal/service` 中实现业务逻辑
4. 在 `internal/dao` 中添加数据访问方法
5. 更新Swagger文档注释

### 错误处理

添加新错误类型：
1. 在 `configs/errors.yaml` 中定义错误
2. 使用 `pkg/errors` 包中的方法创建具体错误

```go
// 使用示例
if token == "" {
    return errors.ConfigMissingParam("token")
}
```

### 日志记录

```go
import (
    "focalboard-tool/library/log"
    "go.uber.org/zap"
)

// 记录信息
log.Info("操作成功", zap.String("operation", "GetBoard"))

// 记录错误
log.Error("操作失败", zap.Error(err), zap.String("boardId", boardId))
```

## 部署

### Docker部署

1. 构建镜像：
```bash
docker build -t focalboard-tool .
```

2. 运行容器：
```bash
docker run -d -p 8080:8080 \
  -v /path/to/config:/app/configs \
  -v /path/to/logs:/app/logs \
  focalboard-tool
```

### 生产环境

1. 修改配置文件中的运行模式为 `release`
2. 配置适当的日志级别
3. 设置合适的超时时间
4. 配置负载均衡和监控

## 监控和维护

### 健康检查

系统提供了基本的健康检查端点，可以通过监控工具检查服务状态。

### 日志监控

- 日志文件位置：`./logs/aa.log`
- 支持日志轮转，自动清理过期日志
- 可以通过配置调整日志保留策略

### 性能监控

- 集成了pprof性能分析工具
- 可以通过HTTP接口查看性能指标
- 支持请求追踪和性能分析

## 许可证

本项目使用MIT许可证，详情请参见LICENSE文件。

## 贡献

欢迎提交问题报告和功能请求。如需贡献代码，请遵循项目的代码规范和提交流程。

## 支持

如果您在使用过程中遇到问题，请通过以下方式获取帮助：

1. 查看Swagger API文档
2. 检查日志文件中的错误信息
3. 提交Issue报告问题
4. 参考项目文档和代码注释 