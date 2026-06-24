# 变更日志

本文件记录 Ovra-Zero 及仓库内配套前端/脚本的**功能性与配置类**变更。  
格式约定见 [.cursor/rules/code-comments-changelog.mdc](.cursor/rules/code-comments-changelog.mdc)。

## [未发布]

### 维护
- **开发规范**：新增 Cursor 规则，要求代码步骤级注释、更新本变更日志、Git 提交优先中文（`.cursor/rules/code-comments-changelog.mdc`）
- **前端动画**：页面切换默认改为 `fade` 淡入淡出，并在 `preferences.ts` 补充动画配置说明注释（`ruoyi-plus-vben5/apps/web-antd/src/preferences.ts`）

### 修复
- **通知公告**：新增时生成 `notice_id` 雪花主键，修复 Duplicate entry '' for key PRIMARY（`app/system/internal/logic/system/notice/add_logic.go`）
- **通知公告**：列表从嵌入 `SysNotice` 正确映射 `noticeId`，修复编辑请求 `row_xx`、删除假成功（`app/system/internal/logic/system/notice/page_set_logic.go`、`app/system/internal/dal/sys_notice.go`）
- **通知公告**：详情接口将 `notice_content` 转为字符串返回（`app/system/internal/logic/system/notice/info_logic.go`）
- **租户管理**：详情/删除/改状态按主键 `id` 查询，修复列表点编辑报「数据不存在」（`app/system/internal/dal/sys_tenant.go`）
- **租户管理**：编辑时 `username`/`password` 改为可选，仅新增租户时校验必填（`desc/system/api/system/tenant.api`、`app/system/internal/types/types.go`、`app/system/internal/logic/system/tenant/add_logic.go`）

### 文档
- **本地开发**：补充 RSA 密钥格式、SSE 404、前端代理等说明（`docs/local-dev.md`）
