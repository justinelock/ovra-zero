# 变更日志

本文件记录 Ovra-Zero 及仓库内配套前端/脚本的**功能性与配置类**变更。  
格式约定见 [.cursor/rules/code-comments-changelog.mdc](.cursor/rules/code-comments-changelog.mdc)。

## [未发布]

### 维护
- **业务占位代码注释**：为 fund/trade/invest 等 21 个 Logic、前端 API/列表页/data 补全步骤级注释，对齐 `.cursor/rules/code-comments-changelog.mdc`（`app/system/internal/logic/`、`ruoyi-plus-vben5/apps/web-antd/src/views/biz/`、`src/api/biz/`）

### 文档
- **业务 API 文档**：列表分页 GET 默认参数补充为 `pageNum=1&pageSize=10`（`docs/biz-api.md`）
- **业务 API 文档**：汇总用户管理～K线管理 22 个列表/统计接口，每接口预留响应 JSON 占位（`docs/biz-api.md`）

### 新增
- **用户列表活跃 Tag**：标题「用户列表」后展示 `活跃用户:{totalActiveSessions}`，新增 `GET /member/user/stats`（`views/member/list/index.vue`、`desc/system/api/member/user.api`）

- **业务用户列表页**：参考系统用户管理实现筛选+分页表格（无部门树），含占位列表 API（`ruoyi-plus-vben5/apps/web-antd/src/views/member/list/`、`desc/system/api/member/user.api`）
- **业务菜单批量占位**：新增资金管理/订单/投信/产品/通知/K线/客服/App 等 9 个一级目录及 18 个子页占位（`bin/sql/patch_biz_menus.sql`、`ruoyi-plus-vben5/apps/web-antd/src/views/biz/`）
- **业务用户管理菜单**：新增一级「用户管理」及 6 个子菜单占位页，与系统管理下后台账号管理并存（`bin/sql/patch_member_menu.sql`、`ruoyi-plus-vben5/apps/web-antd/src/views/member/`）
- **开发参考菜单**：`sys_menu` 新增「开发参考」目录，挂载 `演示使用自行删除` 下全部演示页；补 SSE/加解密演示 API（`bin/sql/patch_dev_reference_menu.sql`、`desc/system/api/dev/`、`app/system/internal/logic/dev/`）

### 变更
- **业务子页列表**：订单/投信/产品/通知/K线等 11 个模块实现筛选+分页表格（`ruoyi-plus-vben5/apps/web-antd/src/views/biz/`、`desc/system/api/trade/` 等）
- **资金管理子页**：钱包申请/账户流水/提现/充值管理实现筛选+分页表格，含占位列表 API（`ruoyi-plus-vben5/apps/web-antd/src/views/biz/fund/`、`desc/system/api/fund/`）
- **业务用户管理子页**：钱包管理/用户报表/团队管理/登录记录参考用户列表实现筛选+分页表格（`ruoyi-plus-vben5/apps/web-antd/src/views/member/`、`desc/system/api/member/`）
- **团队管理统计**：代理统计改为一行圆角卡片，亮/暗主题分别配色且背景透明度 0.8（`ruoyi-plus-vben5/apps/web-antd/src/views/member/team/team-stats.vue`）
- **业务实名认证**：参考用户列表实现筛选+分页表格，含占位列表 API（`ruoyi-plus-vben5/apps/web-antd/src/views/member/kyc/`、`desc/system/api/member/kyc.api`）
- **业务用户列表**：查询区改为单行 inline 布局（关键词/认证状态/日期 + 查询/导出/已删用户）（`ruoyi-plus-vben5/apps/web-antd/src/views/member/list/`）
- **开发规范**：新增 Cursor 规则，要求代码步骤级注释、更新本变更日志、Git 提交优先中文（`.cursor/rules/code-comments-changelog.mdc`）
- **前端动画**：页面切换默认改为 `fade` 淡入淡出，并在 `preferences.ts` 补充动画配置说明注释（`ruoyi-plus-vben5/apps/web-antd/src/preferences.ts`）

### 修复
- **登录过期**：认证中间件改返回 `{code:401}` JSON，修复仅弹提示不跳转登录页（`toolkit/middlewares/auth_middleware.go`、`ruoyi-plus-vben5/apps/web-antd/src/utils/http/checkStatus.ts`）
- **部门管理**：新增/编辑接口兼容 `parentId` 数值类型入参，统一转字符串后再反序列化，修复编辑时报 `type mismatch for field "parentId"`（`app/system/internal/handler/system/dept/add_handler.go`、`app/system/internal/handler/system/dept/update_handler.go`）
- **通知公告**：新增时生成 `notice_id` 雪花主键，修复 Duplicate entry '' for key PRIMARY（`app/system/internal/logic/system/notice/add_logic.go`）
- **通知公告**：列表从嵌入 `SysNotice` 正确映射 `noticeId`，修复编辑请求 `row_xx`、删除假成功（`app/system/internal/logic/system/notice/page_set_logic.go`、`app/system/internal/dal/sys_notice.go`）
- **通知公告**：详情接口将 `notice_content` 转为字符串返回（`app/system/internal/logic/system/notice/info_logic.go`）
- **租户管理**：详情/删除/改状态按主键 `id` 查询，修复列表点编辑报「数据不存在」（`app/system/internal/dal/sys_tenant.go`）
- **租户管理**：编辑时 `username`/`password` 改为可选，仅新增租户时校验必填（`desc/system/api/system/tenant.api`、`app/system/internal/types/types.go`、`app/system/internal/logic/system/tenant/add_logic.go`）

### 文档
- **本地开发**：补充 RSA 密钥格式、SSE 404、前端代理等说明（`docs/local-dev.md`）
