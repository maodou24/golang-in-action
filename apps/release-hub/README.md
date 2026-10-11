# Release Hub

Release Hub 是一个分支管理服务，目标是把手工发版流程产品化

- 产品/项目依赖管理
- 创建 Branch，并同步 Git 仓库
- 基于 token 的访问控制
- 导出项目依赖仓库

项目、默认依赖、版本快照及分支同步结果以 SQLite 为唯一数据源；Manifest XML 从数据库记录实时生成，不作为可编辑配置的第二份存储。

启动前需要设置 `TOKEN_ENCRYPTION_KEY` 和 `SESSION_SECRET`。前者加密保存于数据库的 Git provider token，后者用于签名 HttpOnly 会话 cookie。数据库文件路径由 `DATA_PATH` 配置。

## Release 创建与 Manifest 发布

- `MANIFEST_REPO_URL`：必填，指定 GitHub manifest 仓库，例如 `https://github.com/your-org/manifests.git`。
- `MANIFEST_BASE_BRANCH`：可选，创建 `releasehub` 分支时使用的基线；未设置则使用仓库默认分支。
- 使用当前用户的 GitHub token 创建依赖仓库分支并提交 manifest，需要这些仓库的写权限。

创建顺序：先将 release 和依赖快照提交到数据库，再从各依赖基线创建目标分支，最后从数据库生成 XML 并提交到 manifest 仓库的 `releasehub` 分支。

Manifest 路径为 `projects/<projectId>/releases/<releaseName>/manifest.xml`，使用项目 ID 避免重命名影响路径。分支不存在时创建，存在时复用；同一路径的文件使用现有 SHA 更新，不强制重置分支。默认继承的依赖目标分支为 release 名称；请求中的显式依赖使用各自的 `branchName`。

提交使用 `commitSummary` 和 `commitDescription`，未提供 summary 时自动生成。提交 SHA 保存为 `manifestCommitSha`；GitHub 操作失败时保留数据库记录和逐项 `branchResults`，状态为 `partial_success`，不会回滚已创建的远端分支。数据库仍是配置来源，仓库中的 XML 为发布产物。

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

