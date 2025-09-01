# Project Layout

## single app

```text
my-microservice/                  # 服务根目录
├── cmd/                         # 主程序入口
│   └── server/                  # 服务主程序
│       ├── main.go              # 入口文件
│       └── wire.go              # 依赖注入（可选）
├── internal/                    # 私有代码（外部无法导入）
│   ├── config/                  # 配置处理
│   │   ├── config.go            # 配置结构体
│   │   ├── load.go              # 配置加载
│   │   └── config.yaml          # 配置文件示例
│   ├── domain/                  # 领域层
│   │   ├── user.go              # 领域模型
│   │   ├── value_object.go      # 值对象
│   │   └── repository.go        # 领域仓库接口
│   ├── repository/              # 数据持久层
│   │   ├── user_repository.go   # 用户仓库实现
│   │   └── mysql/               # MySQL 实现
│   │       └── user_repo.go
│   ├── service/                 # 应用服务层
│   │   ├── user_service.go      # 用户服务
│   │   └── auth_service.go      # 认证服务
│   ├── handler/                 # 传输层（HTTP/gRPC等）
│   │   ├── http/                # HTTP处理器
│   │   │   ├── user_handler.go
│   │   │   ├── middleware.go    # 中间件
│   │   │   └── router.go        # 路由配置
│   │   └── grpc/                # gRPC处理器
│   │       └── user_grpc.go
│   ├── client/                  # 外部服务客户端
│   │   ├── user_client.go
│   │   └── http_client.go
│   └── pkg/                     # 内部公共包（可被internal内其他包导入）
│       ├── logger/              # 日志工具
│       ├── database/            # 数据库连接
│       ├── cache/               # 缓存处理
│       └── utils/               # 工具函数
├── api/                         # API定义
│   ├── protos/                  # gRPC proto文件
│   │   └── user.proto
│   └── openapi/                 # OpenAPI/Swagger定义
│       └── swagger.yaml
├── deployments/                 # 部署配置
│   ├── docker/                  # Docker相关
│   │   └── Dockerfile
│   ├── kubernetes/              # K8s配置
│   │   ├── deployment.yaml
│   │   └── service.yaml
│   └── docker-compose.yaml      # 本地开发
├── scripts/                     # 脚本文件
│   ├── migrate.sh               # 数据库迁移
│   ├── start.sh                 # 启动脚本
│   └── build.sh                 # 构建脚本
├── test/                        # 测试相关
│   ├── integration/             # 集成测试
│   ├── e2e/                     # 端到端测试
│   └── fixtures/                # 测试数据
├── go.mod                       # Go模块定义
├── go.sum                       # 依赖校验
├── Makefile                     # 构建命令
├── .gitignore                   # Git忽略配置
├── README.md                    # 项目说明
└── .env.example                 # 环境变量示例
```

## micro servers

```text
monorepo/                          # 项目根目录
├── go.work                        # Go 工作区文件
├── Makefile                       # 全局构建命令
├── README.md                      # 项目说明
├── .gitignore                     # Git 忽略配置
├── .github/                       # GitHub 配置
│   └── workflows/                 # CI/CD 流水线
│       ├── ci.yml                 # 持续集成
│       └── cd.yml                 # 持续部署
├── services/                      # 微服务目录
│   ├── user-service/              # 用户服务
│   │   ├── cmd/                   # 入口程序
│   │   │   └── server/
│   │   │       └── main.go
│   │   ├── internal/              # 私有代码
│   │   │   ├── handler/
│   │   │   ├── service/
│   │   │   └── repository/
│   │   ├── api/                   # API 定义
│   │   ├── Dockerfile             # 服务Dockerfile
│   │   ├── go.mod                 # 独立模块
│   │   └── config/                # 服务配置
│   ├── order-service/             # 订单服务
│   │   └── ...                    # 类似结构
│   ├── product-service/           # 商品服务
│   │   └── ...
│   └── gateway/                   # API 网关
│       └── ...
├── libs/                          # 公共库目录
│   ├── go.mod                     # 公共库模块
│   ├── database/                  # 数据库相关
│   │   ├── mysql/
│   │   ├── redis/
│   │   └── models.go
│   ├── utils/                     # 工具函数
│   │   ├── validator/
│   │   ├── crypto/
│   │   └── strings/
│   ├── middleware/                # 中间件
│   │   ├── auth/
│   │   ├── logging/
│   │   └── recovery/
│   ├── config/                    # 配置处理
│   │   ├── loader.go
│   │   └── types.go
│   ├── logger/                    # 日志库
│   │   ├── zap/
│   │   └── logger.go
│   └── errors/                    # 错误处理
│       ├── types.go
│       └── handler.go
├── api/                           # 全局API定义
│   ├── protos/                    # gRPC proto文件
│   │   ├── user.proto
│   │   ├── order.proto
│   │   └── product.proto
│   └── openapi/                   # OpenAPI定义
│       └── swagger.yaml
├── deployments/                   # 部署配置
│   ├── kubernetes/                # K8s配置
│   │   ├── user-service/
│   │   ├── order-service/
│   │   └── base/                  # 基础配置
│   ├── docker-compose/            # Docker compose
│   │   ├── dev.yaml
│   │   └── prod.yaml
│   └── scripts/                   # 部署脚本
├── scripts/                       # 开发脚本
│   ├── build.sh                   # 构建脚本
│   ├── test.sh                    # 测试脚本
│   └── migrate.sh                 # 迁移脚本
└── docs/                          # 文档
    ├── architecture/              # 架构文档
    ├── api/                       # API文档
    └── development/               # 开发指南
```