# 业务 API 接口文档

本文档汇总**用户管理 → K线管理**各业务子页的列表/统计接口。字段定义以 `desc/system/api/` 下 `.api` 文件为准；**如何从契约跟读到 Handler/Logic/DAL** 见 [read-api-flow.md](./read-api-flow.md)。

> **维护说明**：每个接口下方预留「响应示例」代码块，后续补充含真实字段值的完整 JSON。

## 通用约定

| 项 | 说明 |
| --- | --- |
| 服务 | `system`（本地默认 `http://127.0.0.1:8092`） |
| 网关 | Traefik 可选 `http://127.0.0.1:28080`，前端 Vite 代理一般为 `/api` → system |
| 鉴权 | 需登录 Token；中间件：`Auth`、`Sign` |
| 分页参数 | 所有列表接口继承 `PageReq`：`pageNum`、`pageSize`、`params[beginTime]`、`params[endTime]`（见下节默认值） |

### 分页请求默认参数

所有**列表分页** `GET` 请求，默认携带查询参数：

```text
pageNum=1&pageSize=10
```

示例（用户列表）：

```http
GET /member/user/list?pageNum=1&pageSize=10
```

| 参数 | 类型 | 默认 | 说明 |
| --- | --- | --- | --- |
| `pageNum` | int | `1` | 当前页码，从 1 开始 |
| `pageSize` | int | `10` | 每页条数；前端 VxeGrid 全局默认 10（`adapter/vxe-table.ts`） |
| `params[beginTime]` | string | 无 | 筛选开始时间，可选 |
| `params[endTime]` | string | 无 | 筛选结束时间，可选 |

> **说明**：各接口还可叠加业务筛选参数（如 `keyword`、`status`）。团队统计 `GET /member/team/stats` **非分页**，无 `pageNum`/`pageSize`。若直接调后端且省略 `pageSize`，`desc/system/api/base.api` 中 go-zero 默认值为 `20`，与前端列表页实际传参可能不一致，联调时建议显式传 `pageSize=10`。

### 分页列表响应结构

含 `rows` 字段的响应由 `toolkit/helper/resp.go` **展平**到顶层（`rows` 不在 `data` 内）：

```json
{
  "code": 200,
  "msg": "操作成功",
  "total": 0,
  "rows": []
}
```

### 非分页对象响应结构

不含 `rows` 的响应包在 `data` 内，例如团队统计：

```json
{
  "code": 200,
  "msg": "操作成功",
  "data": {}
}
```

## 接口总览

| 模块 | 页面 | 方法 | 路径 | 权限标识 |
| --- | --- | --- | --- | --- |
| 用户管理 | 用户列表 | GET | `/member/user/list` | `member:list:list` |
| 用户管理 | 活跃用户统计 | GET | `/member/user/stats` | `member:list:list` |
| 用户管理 | 实名认证 | GET | `/member/kyc/list` | `member:kyc:list` |
| 用户管理 | 钱包管理 | GET | `/member/wallet/list` | `member:wallet:list` |
| 用户管理 | 用户报表 | GET | `/member/report/list` | `member:report:list` |
| 用户管理 | 团队管理 | GET | `/member/team/list` | `member:team:list` |
| 用户管理 | 团队统计 | GET | `/member/team/stats` | `member:team:list` |
| 用户管理 | 登录记录 | GET | `/member/loginLog/list` | `member:loginLog:list` |
| 资金管理 | 钱包申请 | GET | `/fund/walletApply/list` | `fund:walletApply:list` |
| 资金管理 | 账户流水 | GET | `/fund/statement/list` | `fund:statement:list` |
| 资金管理 | 提现管理 | GET | `/fund/withdraw/list` | `fund:withdraw:list` |
| 资金管理 | 充值管理 | GET | `/fund/recharge/list` | `fund:recharge:list` |
| 订单管理 | 合约订单 | GET | `/trade/contract/list` | `trade:contract:list` |
| 订单管理 | 委托订单 | GET | `/trade/entrust/list` | `trade:entrust:list` |
| 订单管理 | 成交订单 | GET | `/trade/deal/list` | `trade:deal:list` |
| 投信管理 | 持仓订单 | GET | `/invest/position/list` | `invest:position:list` |
| 投信管理 | 投信列表 | GET | `/invest/list/list` | `invest:list:list` |
| 产品管理 | 产品配置 | GET | `/product/config/list` | `product:config:list` |
| 产品管理 | 产品实时数据 | GET | `/product/realtime/list` | `product:realtime:list` |
| 产品管理 | 产品历史数据 | GET | `/product/history/list` | `product:history:list` |
| 通知管理 | 市场新闻 | GET | `/notify/news/list` | `notify:news:list` |
| 通知管理 | 通知发布 | GET | `/notify/publish/list` | `notify:publish:list` |
| K线管理 | K线管理 | GET | `/kline/main/list` | `kline:main:list` |

---

## 一、用户管理

### 1.1 用户列表

| 项 | 值 |
| --- | --- |
| 方法 | `GET` |
| 路径 | `/member/user/list` |
| 权限 | `member:list:list` |
| API 定义 | `desc/system/api/member/user.api` |
| 行实体 | `MemberUserItem` |

**查询参数**（除分页外）：

| 参数 | 类型 | 说明 |
| --- | --- | --- |
| `keyword` | string | 关键词 |
| `authStatus` | string | 认证状态 |
| `deleted` | string | 是否已删用户 |

**`rows[]` 字段**：

| 字段 | 类型 | 说明 |
| --- | --- | --- |
| `id` | string | 主键 |
| `username` | string | 用户名 |
| `realName` | string | 真实姓名 |
| `idCard` | string | 身份证号 |
| `agentLevel` | int64 | 代理层级 |
| `inviteCode` | string | 邀请码 |
| `commissionRate` | float64 | 佣金比例 |
| `totalCommission` | float64 | 累计佣金 |
| `totalBalance` | float64 | 总余额（USD 钱包 `balance>0` 汇总） |
| `fundPositionAmount` | float64 | 投信持仓（`fb_fund_position` PENDING 本金） |
| `fundPositionDividend` | float64 | 投信分红（`fb_fund_profit_log.status=1` 汇总） |
| `status` | string | 账号状态 |
| `onlineStatus` | int64 | 在线状态（1=Redis token 在线或近 5 分钟 presence 活跃 / 0 离线） |
| `lastLogin` | string | 最后登录时间 |
| `createdAt` | string | 注册时间 |

**响应示例**（待补充）：

```json
{
  "code": 200,
  "msg": "操作成功",
  "total": 0,
  "rows": [
    {
      "id": "2069776834893524993",
      "username": "18691188277",
      "realName": "马小梅",
      "idCard": "610623197011190122",
      "agentLevel": 3,
      "inviteCode": "X4A6B6LM",
      "commissionRate": 0.00,
      "totalCommission": 0.00,
      "totalBalance": 2244.60,
      "fundPositionAmount": 1000.0000,
      "fundPositionDividend": 30.0000,
      "status": "ACTIVE",
      "onlineStatus": 1,
      "lastLogin": "2026-06-24 21:37:37",
      "createdAt": "2026-06-24 21:37:07" 
    }
  ]
}

```

---

### 1.2 用户列表（活跃统计）

| 项 | 值 |
| --- | --- |
| 方法 | `GET` |
| 路径 | `/member/user/stats` |
| 权限 | `member:list:list` |
| API 定义 | `desc/system/api/member/user.api` |
| 响应实体 | `MemberUserStatsResp`（在 `data` 内） |

**查询参数**：无业务筛选（全局 Redis 统计，对齐 Java `online-statistics`）；**无分页参数**。

**`data` 字段**：

| 字段 | 类型 | 说明 |
| --- | --- | --- |
| `totalOnlineUsers` | int64 | Redis `online:user:*` 有效会话用户数 |
| `todayLogins` | int64 | Redis `login:today` Set 大小（今日登录用户） |
| `totalActiveSessions` | int64 | Redis `fb:presence:active` 近 **5 分钟**内有鉴权请求的活跃用户数；前端 Tag「活跃用户:{totalActiveSessions}」 |

**响应示例**（待补充）：

```json
{
  "code": 200,
  "msg": "操作成功",
  "data": {
    "totalOnlineUsers": "355",
    "todayLogins": "399",
    "totalActiveSessions": "2"
  }
}
```

---

### 1.3 实名认证

| 项 | 值 |
| --- | --- |
| 方法 | `GET` |
| 路径 | `/member/kyc/list` |
| 权限 | `member:kyc:list` |
| API 定义 | `desc/system/api/member/kyc.api` |
| 行实体 | `MemberKycItem` |

**查询参数**：

| 参数 | 类型 | 说明 |
| --- | --- | --- |
| `keyword` | string | 关键词 |
| `authStatus` | string | 认证状态 |

**`rows[]` 字段**：

| 字段 | 类型 | 说明 |
| --- | --- | --- |
| `id` | string | 主键 |
| `userId` | string | 用户 ID |
| `username` | string | 用户名 |
| `mobile` | string | 手机号 |
| `realName` | string | 真实姓名 |
| `idCardNo` | string | 身份证号 |
| `idCardFront` | string | 身份证正面 |
| `idCardBack` | string | 身份证反面 |
| `status` | string | 认证状态 |
| `rejectReason` | string | 驳回原因 |
| `verifiedAt` | string | 认证时间 |
| `createdAt` | string | 提交时间 |
| `updatedAt` | string | 更新时间 |

**响应示例**（待补充）：

```json
{
  "code": 200,
  "msg": "操作成功",
  "total": 0,
  "rows": [
    {
      "id": "2067908255562674177",
      "userId": "2067908054152196097",
      "username": "15164459830",
      "mobile": null,
      "realName": "测试吧吧",
      "idCardNo": "220724199803123045",
      "idCardFront": "/ids/3efabddbb0454a64a8a51717cc343f86_2067908054152196097_front.jpeg",
      "idCardBack": "/ids/3efabddbb0454a64a8a51717cc343f86_2067908054152196097_back.jpeg",
      "status": "VERIFIED",
      "rejectReason": null,
      "verifiedAt": "2026-06-19 17:52:18",
      "createdAt": "2026-06-19 17:52:03",
      "updatedAt": "2026-06-19 17:52:18"
    }
  ]
}
```

---

### 1.4 钱包管理

| 项 | 值 |
| --- | --- |
| 方法 | `GET` |
| 路径 | `/member/wallet/list` |
| 权限 | `member:wallet:list` |
| API 定义 | `desc/system/api/member/wallet.api` |
| 行实体 | `MemberWalletItem` |

**查询参数**：

| 参数 | 类型 | 说明 |
| --- | --- | --- |
| `keyword` | string | 关键词 |
| `accountType` | string | 账户类型 |
| `currency` | string | 币种 |
| `frozenStatus` | string | 冻结状态 |

**`rows[]` 字段**：

| 字段 | 类型 | 说明 |
| --- | --- | --- |
| `id` | string | 主键 |
| `userId` | string | 用户 ID |
| `username` | string | 用户名 |
| `mobile` | string | 手机号 |
| `realName` | string | 真实姓名 |
| `accountType` | string | 账户类型 |
| `balance` | float64 | 余额 |
| `frozenAmount` | float64 | 冻结金额 |
| `frozen` | bool | 是否冻结 |
| `version` | string | 乐观锁版本 |
| `currency` | string | 币种 |
| `drawTicket` | int64 | 抽奖券数量 |
| `createdAt` | string | 创建时间 |
| `updatedAt` | string | 更新时间 |

**响应示例**（待补充）：

```json
{
  "code": 200,
  "msg": "操作成功",
  "total": 0,
  "rows": [
    {
      "id": "2068980123812048897",
      "userId": "2068980122423734274",
      "username": "FDD123",
      "mobile": null,
      "realName": "冯吕彬",
      "accountType": "main",
      "balance": 2244.60,
      "frozenAmount": 0.00,
      "frozen": false,
      "version": "3",
      "currency": "USD",
      "drawTicket": 0,
      "createdAt": "2026-06-22 16:51:16",
      "updatedAt": "2026-06-24 21:47:24"
    }
  ]
}
```

---

### 1.5 用户报表

| 项 | 值 |
| --- | --- |
| 方法 | `GET` |
| 路径 | `/member/report/list` |
| 权限 | `member:report:list` |
| API 定义 | `desc/system/api/member/report.api` |
| 行实体 | `MemberReportItem` |

**查询参数**：

| 参数 | 类型 | 说明 |
| --- | --- | --- |
| `keyword` | string | 关键词 |
| `level` | string | 用户等级 |

**`rows[]` 字段**：

| 字段 | 类型 | 说明 |
| --- | --- | --- |
| `id` | string | 主键 |
| `userId` | string | 用户 ID |
| `username` | string | 用户名 |
| `mobile` | string | 手机号 |
| `realName` | string | 真实姓名 |
| `level` | int64 | 等级 |
| `amount` | float64 | 账户余额 |
| `rechargeAmount` | float64 | 累计充值 |
| `withdrawAmount` | float64 | 累计提现 |
| `rechargeDiff` | float64 | 充提差 |
| `totalProfit` | float64 | 累计盈亏 |
| `teamCount` | int64 | 团队人数 |
| `registerTime` | string | 注册时间 |
| `lastLogin` | string | 最后登录 |
| `loginIp` | string | 登录 IP |
| `parentUser` | object | 上级用户 `{ id, username }` |

**响应示例**（待补充）：

```json
{
  "code": 200,
  "msg": "操作成功",
  "total": 0,
  "rows": [
    {
      "id": "2068980122423734274",
      "userId": "2068980122423734274",
      "username": "FDD123",
      "mobile": null,
      "realName": "冯吕彬",
      "level": 0,
      "amount": 2244.60,
      "rechargeAmount": 3000.00,
      "withdrawAmount": 0,
      "rechargeDiff": 3000.00,
      "totalProfit": 0,
      "teamCount": 0,
      "registerTime": "2026-06-22 16:51:16",
      "lastLogin": "2026-06-25 00:46:45",
      "loginIp": "223.104.72.141",
      "parentUser": {
        "id": "2000844918397870081",
        "username": "wu52886"
      }
    }
  ]
}
```

### 1.5.1 用户报表流水

| 项 | 值 |
| --- | --- |
| 方法 | `GET` |
| 路径 | `/member/report/flow/{userId}` |
| 权限 | `member:report:list` |
| API 定义 | `desc/system/api/member/report.api` |
| 行实体 | `MemberReportFlowItem` |

**路径参数**：`userId` — 业务用户 ID。

**`rows[]` 字段**：

| 字段 | 类型 | 说明 |
| --- | --- | --- |
| `id` | string | 主键 |
| `userId` | string | 用户 ID |
| `username` | string | 用户名 |
| `mobile` | string | 手机号 |
| `realName` | string | 真实姓名 |
| `accountType` | string | 账户类型 |
| `flowType` | string | 流水类型 |
| `beforeAmount` | float64 | 变动前余额 |
| `flowAmount` | float64 | 变动金额 |
| `afterAmount` | float64 | 变动后余额 |
| `businessNo` | string | 业务单号 |
| `remark` | string | 备注 |
| `createdAt` | string | 创建时间 |
| `walletId` | string | 钱包 ID |
| `currency` | string | 币种 |
| `description` | string | 描述 |
| `status` | string | 状态 |
| `updatedAt` | string | 更新时间 |

**响应示例**：

```json
{
  "code": 200,
  "msg": "操作成功",
  "total": 0,
  "rows": [
    {
      "id": "2069775650615013378",
      "userId": "2068980122423734274",
      "username": "FDD123",
      "mobile": null,
      "realName": "冯吕彬",
      "accountType": "main",
      "flowType": "CONTRACT_BUY",
      "beforeAmount": 2229.00,
      "flowAmount": -2229.00,
      "afterAmount": 0.00,
      "businessNo": "CONTRACT_BUY_2069775650594041858",
      "remark": "购买ETH/EUR合约",
      "createdAt": "2026-06-24 21:32:24",
      "walletId": "2068980123812048897",
      "currency": "USD",
      "description": "购买ETH/EUR合约",
      "status": "SUCCESS",
      "updatedAt": null
    }
  ]
}
```

---

### 1.6 团队管理（列表）

| 项 | 值 |
| --- | --- |
| 方法 | `GET` |
| 路径 | `/member/team/list` |
| 权限 | `member:team:list` |
| API 定义 | `desc/system/api/member/team.api` |
| 行实体 | `MemberTeamItem` |

**查询参数**：

| 参数 | 类型 | 说明 |
| --- | --- | --- |
| `keyword` | string | 关键词 |
| `status` | string | 状态 |

**`rows[]` 字段**：

| 字段 | 类型 | 说明 |
| --- | --- | --- |
| `id` | string | 主键 |
| `username` | string | 用户名 |
| `agent` | object | 上级代理 `{ id, username, realName }` |
| `realName` | string | 真实姓名 |
| `level` | int64 | 用户等级 |
| `status` | string | 状态 |
| `teamSize` | int64 | 团队规模 |
| `level1Members` ~ `level5Members` | int64 | 各级成员数 |
| `agentLevel` | int64 | 代理层级 |
| `walletCount` | int64 | 钱包数量 |
| `balance` | float64 | 个人余额 |
| `totalTeamBalance` | float64 | 团队总余额 |
| `createdAt` | string | 注册时间 |

**响应示例**（待补充）：

```json
{
  "code": 200,
  "msg": "操作成功",
  "total": 0,
  "rows": [
    {
      "id": "2068980122423734274",
      "username": "FDD123",
      "agent": {
        "id": "2000844918397870081",
        "username": "wu52886",
        "realName": "吴彩艳"
      },
      "realName": "冯吕彬",
      "level": 0,
      "status": "ACTIVE",
      "teamSize": 0,
      "level1Members": 0,
      "level2Members": 0,
      "level3Members": 0,
      "level4Members": 0,
      "level5Members": 0,
      "agentLevel": 3,
      "walletCount": 0,
      "balance": 2244.60,
      "totalTeamBalance": 0,
      "createdAt": null 
    }
  ]
}
```

---

### 1.7 团队管理（统计）

| 项 | 值 |
| --- | --- |
| 方法 | `GET` |
| 路径 | `/member/team/stats` |
| 权限 | `member:team:list` |
| API 定义 | `desc/system/api/member/team.api` |
| 响应实体 | `MemberTeamStatsResp`（在 `data` 内） |

**查询参数**：同团队列表（`keyword`、`status`）。

**`data` 字段**：

| 字段 | 类型 | 说明 |
| --- | --- | --- |
| `totalMembers` | int64 | 成员总数 |
| `level1Members` | int64 | 一级代理数 |
| `level2Members` | int64 | 二级代理数 |
| `level3Members` | int64 | 三级代理数 |
| `level4Members` | int64 | 四级代理数 |
| `level5Members` | int64 | 五级代理数 |
| `totalTeamBalance` | float64 | 团队总余额 |

**响应示例**（待补充）：

```json
{
  "code": 200,
  "msg": "操作成功",
  "data": {
    "totalMembers": 29,
    "level1Members": 4,
    "level2Members": 5,
    "level3Members": 5,
    "level4Members": 5,
    "level5Members": 10,
    "totalTeamBalance": 790041.88
  }
}
```

---

### 1.8 登录记录

| 项 | 值 |
| --- | --- |
| 方法 | `GET` |
| 路径 | `/member/loginLog/list` |
| 权限 | `member:loginLog:list` |
| API 定义 | `desc/system/api/member/login_log.api` |
| 行实体 | `MemberLoginLogItem` |

**查询参数**：

| 参数 | 类型 | 说明 |
| --- | --- | --- |
| `keyword` | string | 关键词 |
| `loginResult` | string | 登录结果 |
| `loginMethod` | string | 登录方式 |
| `riskLevel` | string | 风险等级 |
| `ipType` | string | IP 类型 |
| `timeRange` | string | 时间范围 |
| `failReason` | string | 失败原因 |
| `deviceType` | string | 设备类型 |
| `browser` | string | 浏览器 |

**`rows[]` 字段**：

| 字段 | 类型 | 说明 |
| --- | --- | --- |
| `id` | string | 主键 |
| `userId` | string | 用户 ID |
| `username` | string | 用户名 |
| `realName` | string | 真实姓名 |
| `deviceId` | string | 设备 ID |
| `loginTime` | string | 登录时间 |
| `loginIp` | string | 登录 IP |
| `loginLocation` | string | 登录地点 |
| `loginType` | string | 登录方式 |
| `loginResult` | string | 登录结果 |
| `failReason` | string | 失败原因 |
| `riskLevel` | string | 风险等级 |
| `riskDetail` | string | 风险详情 |

**响应示例**（待补充）：

```json
{
  "code": 200,
  "msg": "操作成功",
  "total": 0,
  "rows": [
    {
      "id": "2069849388882669569",
      "userId": "2051263591276912641",
      "username": "15009113390",
      "realName": "卢琴琴",
      "deviceId": "4d2e6539d2b783c3bf1f886344fdca1e",
      "loginTime": "2026-06-25 02:25:25",
      "loginIp": "111.19.76.126",
      "loginLocation": "Unknown",
      "loginType": "PASSWORD",
      "loginResult": "SUCCESS",
      "failReason": null,
      "riskLevel": "LOW",
      "riskDetail": null
    }
  ]
}
```

---

## 二、资金管理

### 2.1 钱包申请

| 项 | 值 |
| --- | --- |
| 方法 | `GET` |
| 路径 | `/fund/walletApply/list` |
| 权限 | `fund:walletApply:list` |
| API 定义 | `desc/system/api/fund/wallet_apply.api` |
| 行实体 | `FundWalletApplyItem` |

**查询参数**：`keyword`、`status`、`accountType`

**`rows[]` 字段**：`id`、`userName`、`realName`、`phoneNumber`、`accountType`、`status`、`riskScore`、`auditOpinion`、`applyTime`、`auditTime`、`auditor`

**响应示例**（待补充）：

```json
{
  "code": 200,
  "msg": "操作成功",
  "total": 0,
  "rows": []
}
```

---

### 2.2 账户流水

| 项 | 值 |
| --- | --- |
| 方法 | `GET` |
| 路径 | `/fund/statement/list` |
| 权限 | `fund:statement:list` |
| API 定义 | `desc/system/api/fund/statement.api` |
| 行实体 | `FundStatementItem` |

**查询参数**：`keyword`、`tradeType`、`tradeStatus`、`currency`

**`rows[]` 字段**：`id`、`userName`、`phoneNumber`、`realName`、`accountType`、`tradeType`、`changeAmount`、`balanceBefore`、`balanceAfter`、`tradeStatus`、`currency`、`tradeDesc`、`remark`、`tradeTime`

**响应示例**（待补充）：

```json
{
  "code": 200,
  "msg": "操作成功",
  "total": 0,
  "rows": []
}
```

---

### 2.3 提现管理

| 项 | 值 |
| --- | --- |
| 方法 | `GET` |
| 路径 | `/fund/withdraw/list` |
| 权限 | `fund:withdraw:list` |
| API 定义 | `desc/system/api/fund/withdraw.api` |
| 行实体 | `FundWithdrawItem` |

**查询参数**：`keyword`、`withdrawStatus`、`withdrawType`

**`rows[]` 字段**：`id`、`userName`、`realName`、`withdrawAmount`、`usdtAddress`、`withdrawStatus`、`withdrawType`、`createTime`、`updateTime`

**响应示例**（待补充）：

```json
{
  "code": 200,
  "msg": "操作成功",
  "total": 0,
  "rows": []
}
```

---

### 2.4 充值管理

| 项 | 值 |
| --- | --- |
| 方法 | `GET` |
| 路径 | `/fund/recharge/list` |
| 权限 | `fund:recharge:list` |
| API 定义 | `desc/system/api/fund/recharge.api` |
| 行实体 | `FundRechargeItem` |

**查询参数**：`keyword`、`status`

**`rows[]` 字段**：`id`、`userName`、`phoneNumber`、`realName`、`rechargeAmount`、`status`、`rechargeImage`、`remark`、`createTime`、`updateTime`

**响应示例**（待补充）：

```json
{
  "code": 200,
  "msg": "操作成功",
  "total": 0,
  "rows": []
}
```

---

## 三、订单管理

### 3.1 合约订单

| 项 | 值 |
| --- | --- |
| 方法 | `GET` |
| 路径 | `/trade/contract/list` |
| 权限 | `trade:contract:list` |
| API 定义 | `desc/system/api/trade/contract.api` |
| 行实体 | `TradeContractItem` |

**查询参数**：`keyword`、`status`、`symbol`、`direction`

**`rows[]` 字段**：`id`、`userName`、`realName`、`balanceU`、`symbol`、`direction`、`tradeAmount`、`actualProfit`、`status`、`durationSec`、`openPrice`、`closePrice`、`orderControl`、`controlResult`、`globalControl`、`openTime`、`settleTime`

**响应示例**（待补充）：

```json
{
  "code": 200,
  "msg": "操作成功",
  "total": 0,
  "rows": []
}
```

---

### 3.2 委托订单

| 项 | 值 |
| --- | --- |
| 方法 | `GET` |
| 路径 | `/trade/entrust/list` |
| 权限 | `trade:entrust:list` |
| API 定义 | `desc/system/api/trade/entrust.api` |
| 行实体 | `TradeEntrustItem` |

**查询参数**：`keyword`、`orderStatus`、`marketCode`

**`rows[]` 字段**：`id`、`orderNo`、`userName`、`marketCode`、`stockCode`、`orderType`、`entrustPrice`、`entrustQty`、`dealQty`、`orderStatus`、`stopLossPrice`、`priceType`、`priceFloatRange`、`validity`、`fee`、`direction`、`limitPrice`、`entrustTime`

**响应示例**（待补充）：

```json
{
  "code": 200,
  "msg": "操作成功",
  "total": 0,
  "rows": []
}
```

---

### 3.3 成交订单

| 项 | 值 |
| --- | --- |
| 方法 | `GET` |
| 路径 | `/trade/deal/list` |
| 权限 | `trade:deal:list` |
| API 定义 | `desc/system/api/trade/deal.api` |
| 行实体 | `TradeDealItem` |

**查询参数**：`keyword`、`marketCode`、`dealType`

**`rows[]` 字段**：`id`、`dealNo`、`userName`、`marketCode`、`stockCode`、`dealType`、`dealPrice`、`dealQty`、`dealAmount`、`fee`、`dealTime`

**响应示例**（待补充）：

```json
{
  "code": 200,
  "msg": "操作成功",
  "total": 0,
  "rows": []
}
```

---

## 四、投信管理

### 4.1 持仓订单

| 项 | 值 |
| --- | --- |
| 方法 | `GET` |
| 路径 | `/invest/position/list` |
| 权限 | `invest:position:list` |
| API 定义 | `desc/system/api/invest/position.api` |
| 行实体 | `InvestPositionItem` |

**查询参数**：`keyword`、`investCode`、`positionStatus`

**`rows[]` 字段**：`id`、`userName`、`realName`、`investCode`、`positionAmount`、`buyDate`、`startDate`、`endDate`、`periodDays`、`fixedYieldRate`、`positionStatus`、`endStatus`、`lastProfitDate`、`createTime`、`updateTime`

**响应示例**（待补充）：

```json
{
  "code": 200,
  "msg": "操作成功",
  "total": 0,
  "rows": []
}
```

---

### 4.2 投信列表

| 项 | 值 |
| --- | --- |
| 方法 | `GET` |
| 路径 | `/invest/list/list` |
| 权限 | `invest:list:list` |
| API 定义 | `desc/system/api/invest/list.api` |
| 行实体 | `InvestListItem` |

**查询参数**：`keyword`、`status`、`soldOut`

**`rows[]` 字段**：`id`、`investCode`、`investName`、`logo`、`productDesc`、`status`、`soldOut`、`sortOrder`、`yieldDisplay`、`yieldRate`、`investableAmount`、`minAddAmount`、`period`、`yieldType`、`expireDate`、`createTime`、`updateTime`

**响应示例**（待补充）：

```json
{
  "code": 200,
  "msg": "操作成功",
  "total": 0,
  "rows": []
}
```

---

## 五、产品管理

### 5.1 产品配置

| 项 | 值 |
| --- | --- |
| 方法 | `GET` |
| 路径 | `/product/config/list` |
| 权限 | `product:config:list` |
| API 定义 | `desc/system/api/product/config.api` |
| 行实体 | `ProductConfigItem` |

**查询参数**：`keyword`、`productType`、`status`

**`rows[]` 字段**：`id`、`productCode`、`productAlias`、`symbol`、`productName`、`productType`、`market`、`odds`、`tradeTime`、`currencies`、`status`、`createTime`、`updateTime`

**响应示例**（待补充）：

```json
{
  "code": 200,
  "msg": "操作成功",
  "total": 0,
  "rows": []
}
```

---

### 5.2 产品实时数据

| 项 | 值 |
| --- | --- |
| 方法 | `GET` |
| 路径 | `/product/realtime/list` |
| 权限 | `product:realtime:list` |
| API 定义 | `desc/system/api/product/realtime.api` |
| 行实体 | `ProductRealtimeItem` |

**查询参数**：`keyword`、`productCode`、`type`

**`rows[]` 字段**：`id`、`productCode`、`tradeCode`、`name`、`type`、`currentPrice`、`changeAmount`、`changePercent`、`openPrice`、`highPrice`、`lowPrice`、`volume`、`turnover`、`direction`、`tradeDate`、`updateTime`

**响应示例**（待补充）：

```json
{
  "code": 200,
  "msg": "操作成功",
  "total": 0,
  "rows": []
}
```

---

### 5.3 产品历史数据

| 项 | 值 |
| --- | --- |
| 方法 | `GET` |
| 路径 | `/product/history/list` |
| 权限 | `product:history:list` |
| API 定义 | `desc/system/api/product/history.api` |
| 行实体 | `ProductHistoryItem` |

**查询参数**：`keyword`、`productCode`、`market`

**`rows[]` 字段**：`id`、`productCode`、`productName`、`productType`、`market`、`tradeDate`、`openPrice`、`highPrice`、`lowPrice`、`closePrice`、`changeAmount`、`changePercent`、`volume`、`turnover`、`updateTime`

**响应示例**（待补充）：

```json
{
  "code": 200,
  "msg": "操作成功",
  "total": 0,
  "rows": []
}
```

---

## 六、通知管理

### 6.1 市场新闻

| 项 | 值 |
| --- | --- |
| 方法 | `GET` |
| 路径 | `/notify/news/list` |
| 权限 | `notify:news:list` |
| API 定义 | `desc/system/api/notify/news.api` |
| 行实体 | `NotifyNewsItem` |

**查询参数**：`keyword`、`source`

**`rows[]` 字段**：`id`、`title`、`summary`、`content`、`source`、`link`、`image`、`viewCount`、`publishTime`

**响应示例**（待补充）：

```json
{
  "code": 200,
  "msg": "操作成功",
  "total": 0,
  "rows": []
}
```

---

### 6.2 通知发布

| 项 | 值 |
| --- | --- |
| 方法 | `GET` |
| 路径 | `/notify/publish/list` |
| 权限 | `notify:publish:list` |
| API 定义 | `desc/system/api/notify/publish.api` |
| 行实体 | `NotifyPublishItem` |

**查询参数**：`keyword`、`notifyType`、`readStatus`

**`rows[]` 字段**：`id`、`receiveUserId`、`title`、`content`、`notifyType`、`readStatus`、`deleted`、`sendTime`、`expireTime`

**响应示例**（待补充）：

```json
{
  "code": 200,
  "msg": "操作成功",
  "total": 0,
  "rows": []
}
```

---

## 七、K线管理

### 7.1 K线管理

| 项 | 值 |
| --- | --- |
| 方法 | `GET` |
| 路径 | `/kline/main/list` |
| 权限 | `kline:main:list` |
| API 定义 | `desc/system/api/kline/main.api` |
| 行实体 | `KlineMainItem` |

**查询参数**：`keyword`、`productCode`、`symbol`

**`rows[]` 字段**：`id`、`productCode`、`symbol`、`redisKey`、`barCount`、`latestTime`、`open`、`high`、`low`、`close`、`volume`

**响应示例**（待补充）：

```json
{
  "code": 200,
  "msg": "操作成功",
  "total": 0,
  "rows": []
}
```

---

## 相关文档

- [本地开发说明](./local-dev.md)
- [API 定义目录](../desc/system/api/)
