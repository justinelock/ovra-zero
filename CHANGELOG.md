# 变更日志

本文件记录 Ovra-Zero 及仓库内配套前端/脚本的**功能性与配置类**变更。  
格式约定见 [.cursor/rules/code-comments-changelog.mdc](.cursor/rules/code-comments-changelog.mdc)。

## [未发布]

### 变更
- **用户报表流水抽屉**：`GET /member/report/flow/{userId}` 查询对齐 Java `selectPageWithUser`；抽屉 9 列与「{姓名}的流水记录」标题（`fb_report.go`、`report-flow-drawer.vue`）
- **用户报表列表**：`GET /member/report/list` 两阶段对齐 Java `getPageData`（批量聚合余额/充提/盈亏/团队/登录 IP）（`fb_report.go`、`report.api`）
- **表格用户名列复制**：抽取 `renderCopyableValue`/`copyText` 公共工具，所有含「用户名」列的 member/biz 列表均支持一键复制（`utils/render-copyable.tsx`、`views/**/data.tsx`）
- **实名认证列表查询**：`GET /member/kyc/list` 对齐 Java `selectPageWithUser`（JOIN 字段、keyword/idCardNo/BETWEEN 时间等）（`fb_member.go`、`kyc.api`）
- **操作列按钮间距**：新增全局 `TableActionSpace`（2px），替换各页 `#action` 内默认 8px `Space`（`components/global/table-action-space.vue`、各 `views/**/index.vue`）
- **业务用户编辑判重**：`PUT /member/user` 在用户名或身份证号变更时查询去重，冲突返回「用户名已经存在」「身份证号已经存在」（`fb_member.go`）
- **业务用户编辑头像**：编辑抽屉支持裁剪上传头像（参考系统用户 Avatar + 个人中心 CropperAvatar）；`PUT /member/user` 增加 `avatar` 字段（`user-edit-drawer.vue`、`member/user.api`、`fb_member.go`）
- **用户详情接口**：`GET /member/user/{id}` 响应改为独立实体 `MemberUserInfoResp`，补全 fb_users 字段、上级摘要、钱包/投信聚合与 Redis 在线状态（`member/user.api`、`fb_member.go`、`info_logic.go`）
- **业务用户编辑表单**：改为右侧抽屉双列布局，字段对齐管理端（密码/密保/佣金/实名/合约控制等）（`member/user.api`、`user-edit-drawer.vue`）

### 新增
- **钱包删除**：`DELETE /member/wallet/{ids}` 按主键物理删除 `fb_user_wallets`；钱包管理操作列改为仅「删除」（`wallet.api`、`fb_user_wallet.go`、`views/member/wallet/index.vue`）
- **实名认证详情弹窗**：列表「详情」左侧 Descriptions、右侧身份证正反面预览（`ruoyi-plus-vben5/apps/web-antd/src/views/member/kyc/kyc-detail-modal.vue`）
- **已删用户抽屉**：列表「已删用户」打开抽屉展示 `deleted=1` 列表，操作列「恢复」调用 `PUT /member/user/restore/{ids}`（`deleted-users-drawer.vue`、`restore_logic.go`）
- **业务用户编辑**：`GET /member/user/{id}` 详情与 `PUT /member/user` 保存；列表编辑按 id 拉取并提交（`member/user.api`、`fb_member.go`、`views/member/list/`）
- **业务用户删除/重置密码**：`DELETE /member/user/{ids}` 逻辑删、`PUT /member/user/resetPwd` MD5 更新密码（`member/user.api`、`fb_member.go`、`views/member/list/`）
- **钱包加减款接口**：新增 `PUT /member/wallet/addOrSubtract`，按 Java `addOrSubtract` 实现加款/减款、流水写入与 flowType 校验（`desc/system/api/member/wallet.api`、`dal/fb_user_wallet.go`、`logic/member/wallet/add_or_subtract_logic.go`）
- **用户管理接口对接 fb_* 表**：9 个 member 列表/统计/流水接口接入真实查询；新增 `FbMemberDal` 封装多表 JOIN 与聚合（`app/system/internal/dal/fb_member.go`、`app/system/internal/logic/member/`）
- **报表流水**：新增 `GET /member/report/flow/{userId}` 及前端抽屉展示（`desc/system/api/member/report.api`、`views/member/report/report-flow-drawer.vue`）
- **gentool fb_* 表**：`gen/system/gen.yaml` 增加 8 张业务表代码生成（`app/system/internal/dal/model/fb_*.gen.go`）

### 变更
- **用户列表/活跃统计对齐 Java**：USD 钱包余额、profit_log 投信分红、keyword 精确匹配；`onlineStatus` 与 stats 三项改 Redis 全局（`fb_user_redis.go`、`fb_member.go`、`member/user/*_logic.go`）
- **用户列表表格增强**：真实姓名/邀请码/ID 可复制；余额点开钱包抽屉；投信持仓分红着色；ACTIVE 显示正常（`views/member/list/`）
- **用户钱包抽屉列对齐 Java**：余额抽屉改为钱包类型/可用余额/冻结金额/券数量/货币；操作列合并为 ±款、开/关账户、冻/解金额、±抽奖券、划转五组按钮；加减款弹窗支持按加款/减款切换 flowType，并按 `id/type/amount/remark/flowType` 对接 `addOrSubtract`（`views/member/list/user-wallet-drawer.vue`、`wallet-action-modal.vue`、`api/member/wallet/index.ts`）
- **前端 Vite 代理**：统一 `/api` → Traefik `28080`；网关补全 `/member` 等业务 PathPrefix（`vite.config.ts`、`bin/traefik/dynamic.yaml`）
- **用户管理契约对齐**：`docs/biz-api.md` 1.1～1.8 字段表与响应 JSON 一致；前端 `api/member`、`views/member` 列字段改为 `username`、`createdAt`、`level1Members` 等（`ruoyi-plus-vben5/apps/web-antd/src/api/member/`、`views/member/`）

### 维护
- **注释规范**：步骤说明写在对应代码行旁，函数外保留一行或多行总结；更新 `.cursor/rules/code-comments-changelog.mdc` 示例（`member/user/*_logic.go`、`fb_user_redis.go`、`views/member/list/index.vue`）
- **用户管理 API 类型注释**：在 `desc/system/api/member/*.api` 补充类型/字段说明，`make api-system` 同步至 `types.go`；新增 `types/member_doc.go` 说明生成约定（勿手改 types.go）

### 文档
- **1.2 活跃统计**：`docs/biz-api.md` 修正 `MemberUserStatsResp` 示例数值类型为 int64
- **命名约定**：`docs/biz-api.md` 增加 API 命名约定章节
- **接口代码跟读指南**：新增 `docs/read-api-flow.md`（onboarding：契约 → 路由 → Handler → Logic → DAL，以 `/member/user/list` 为例）；`local-dev.md`、`biz-api.md` 增加交叉链接
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
- **钱包加减款 Unknown column UpdateBy**：AuditPlugin 更新回调改为仅当表模型含 `UpdateBy`/`UpdateTime` 字段时才写入（`toolkit/gorm/plugin/audit_plugin.go`）
- **钱包加减款响应**：`addOrSubtract` 改 `OkJsonCtx` 返回 `{code,msg}`，修复前端无法关弹窗（`handler/member/wallet/add_or_subtract_handler.go`）`GET /member/wallet/list` 补充 `userId` 并按 `w.user_id` 精确筛选；`PageWallets` SQL 对齐 Java `selectPageWithUser`（keyword 仅 username/mobile/realName、BETWEEN 时间、status/verified 等）（`dal/fb_member.go`、`wallet.api`）
- **登录过期**：认证中间件改返回 `{code:401}` JSON，修复仅弹提示不跳转登录页（`toolkit/middlewares/auth_middleware.go`、`ruoyi-plus-vben5/apps/web-antd/src/utils/http/checkStatus.ts`）
- **部门管理**：新增/编辑接口兼容 `parentId` 数值类型入参，统一转字符串后再反序列化，修复编辑时报 `type mismatch for field "parentId"`（`app/system/internal/handler/system/dept/add_handler.go`、`app/system/internal/handler/system/dept/update_handler.go`）
- **通知公告**：新增时生成 `notice_id` 雪花主键，修复 Duplicate entry '' for key PRIMARY（`app/system/internal/logic/system/notice/add_logic.go`）
- **通知公告**：列表从嵌入 `SysNotice` 正确映射 `noticeId`，修复编辑请求 `row_xx`、删除假成功（`app/system/internal/logic/system/notice/page_set_logic.go`、`app/system/internal/dal/sys_notice.go`）
- **通知公告**：详情接口将 `notice_content` 转为字符串返回（`app/system/internal/logic/system/notice/info_logic.go`）
- **租户管理**：详情/删除/改状态按主键 `id` 查询，修复列表点编辑报「数据不存在」（`app/system/internal/dal/sys_tenant.go`）
- **租户管理**：编辑时 `username`/`password` 改为可选，仅新增租户时校验必填（`desc/system/api/system/tenant.api`、`app/system/internal/types/types.go`、`app/system/internal/logic/system/tenant/add_logic.go`）

### 文档
- **本地开发**：补充 RSA 密钥格式、SSE 404、前端代理等说明（`docs/local-dev.md`）
