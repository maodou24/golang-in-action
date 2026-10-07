# Release Hub

Release Hub 是一个分支管理服务，目标是把手工发版流程产品化

- 产品/项目依赖管理
- 创建 Branch，并同步 Git 仓库
- 基于 token 的访问控制
- 导出项目依赖仓库

项目、默认依赖、版本快照及分支同步结果以 SQLite 为唯一数据源；Manifest XML 从数据库记录实时生成，不作为可编辑配置的第二份存储。

启动前需要设置 `TOKEN_ENCRYPTION_KEY` 和 `SESSION_SECRET`。前者加密保存于数据库的 Git provider token，后者用于签名 HttpOnly 会话 cookie。数据库文件路径由 `DATA_PATH` 配置。

## 架构设计

User --> Release Hub --> SQLite DB
            |
            +--> Git hosting provider branches
            +--> Generated Manifest XML

## 结构关系

- Namespace
    - Project
        - Release
            - Dependencies
            - Branch sync results

