-- 业务菜单批量占位（资金管理/订单/投信/产品/通知/K线/客服/App）
-- 已有库执行本脚本；全新安装已含在 ovra_zero-2.0.sql

BEGIN;

-- 一级目录
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_dept`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES
('2100', '资金管理', '0', 9, 'fund', NULL, '', 1, 0, 'M', '0', '0', '', 'mdi:shield-check-outline', 103, 1, NOW(), NULL, NULL, ''),
('2200', '订单管理', '0', 10, 'trade', NULL, '', 1, 0, 'M', '0', '0', '', 'mdi:database', 103, 1, NOW(), NULL, NULL, ''),
('2300', '投信管理', '0', 11, 'invest', NULL, '', 1, 0, 'M', '0', '0', '', 'mdi:view-grid-outline', 103, 1, NOW(), NULL, NULL, ''),
('2400', '产品管理', '0', 12, 'product', NULL, '', 1, 0, 'M', '0', '0', '', 'mdi:flag-outline', 103, 1, NOW(), NULL, NULL, ''),
('2500', '通知管理', '0', 13, 'notify', NULL, '', 1, 0, 'M', '0', '0', '', 'mdi:cloud-outline', 103, 1, NOW(), NULL, NULL, ''),
('2600', 'K线管理', '0', 14, 'kline', NULL, '', 1, 0, 'M', '0', '0', '', 'mdi:chart-line', 103, 1, NOW(), NULL, NULL, ''),
('2700', '在线客服', '0', 15, 'cs', NULL, '', 1, 0, 'M', '0', '0', '', 'mdi:chat-processing-outline', 103, 1, NOW(), NULL, NULL, ''),
('2800', 'App品牌资源', '0', 16, 'app-brand', NULL, '', 1, 0, 'M', '0', '0', '', 'mdi:earth', 103, 1, NOW(), NULL, NULL, ''),
('2900', 'App版本管理', '0', 17, 'app-version', NULL, '', 1, 0, 'M', '0', '0', '', 'mdi:view-grid', 103, 1, NOW(), NULL, NULL, '');

-- 二级页面
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_dept`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES
('2101', '钱包申请', '2100', 1, 'wallet-apply', 'biz/fund/walletApply/index', '', 1, 0, 'C', '0', '0', 'fund:walletApply:list', 'mdi:file-document-outline', 103, 1, NOW(), NULL, NULL, ''),
('2102', '账户流水', '2100', 2, 'statement', 'biz/fund/statement/index', '', 1, 0, 'C', '0', '0', 'fund:statement:list', 'mdi:shield-yen', 103, 1, NOW(), NULL, NULL, ''),
('2103', '提现管理', '2100', 3, 'withdraw', 'biz/fund/withdraw/index', '', 1, 0, 'C', '0', '0', 'fund:withdraw:list', 'mdi:cash-minus', 103, 1, NOW(), NULL, NULL, ''),
('2104', '充值管理', '2100', 4, 'recharge', 'biz/fund/recharge/index', '', 1, 0, 'C', '0', '0', 'fund:recharge:list', 'mdi:cash-plus', 103, 1, NOW(), NULL, NULL, ''),
('2201', '合约订单', '2200', 1, 'contract', 'biz/trade/contract/index', '', 1, 0, 'C', '0', '0', 'trade:contract:list', 'mdi:printer', 103, 1, NOW(), NULL, NULL, ''),
('2202', '委托订单', '2200', 2, 'entrust', 'biz/trade/entrust/index', '', 1, 0, 'C', '0', '0', 'trade:entrust:list', 'mdi:cloud-download-outline', 103, 1, NOW(), NULL, NULL, ''),
('2203', '成交订单', '2200', 3, 'deal', 'biz/trade/deal/index', '', 1, 0, 'C', '0', '0', 'trade:deal:list', 'mdi:cloud-check-outline', 103, 1, NOW(), NULL, NULL, ''),
('2301', '持仓订单', '2300', 1, 'position', 'biz/invest/position/index', '', 1, 0, 'C', '0', '0', 'invest:position:list', 'mdi:gold', 103, 1, NOW(), NULL, NULL, ''),
('2302', '投信列表', '2300', 2, 'list', 'biz/invest/list/index', '', 1, 0, 'C', '0', '0', 'invest:list:list', 'mdi:view-list', 103, 1, NOW(), NULL, NULL, ''),
('2401', '产品配置', '2400', 1, 'config', 'biz/product/config/index', '', 1, 0, 'C', '0', '0', 'product:config:list', 'mdi:book-open-outline', 103, 1, NOW(), NULL, NULL, ''),
('2402', '产品实时数据', '2400', 2, 'realtime', 'biz/product/realtime/index', '', 1, 0, 'C', '0', '0', 'product:realtime:list', 'mdi:calendar-clock', 103, 1, NOW(), NULL, NULL, ''),
('2403', '产品历史数据', '2400', 3, 'history', 'biz/product/history/index', '', 1, 0, 'C', '0', '0', 'product:history:list', 'mdi:code-brackets', 103, 1, NOW(), NULL, NULL, ''),
('2501', '市场新闻', '2500', 1, 'news', 'biz/notify/news/index', '', 1, 0, 'C', '0', '0', 'notify:news:list', 'mdi:file-search-outline', 103, 1, NOW(), NULL, NULL, ''),
('2502', '通知发布', '2500', 2, 'publish', 'biz/notify/publish/index', '', 1, 0, 'C', '0', '0', 'notify:publish:list', 'mdi:bullhorn-outline', 103, 1, NOW(), NULL, NULL, ''),
('2601', 'K线管理', '2600', 1, 'main', 'biz/kline/main/index', '', 1, 0, 'C', '0', '0', 'kline:main:list', 'mdi:chart-line', 103, 1, NOW(), NULL, NULL, ''),
('2701', '在线客服', '2700', 1, 'main', 'biz/cs/main/index', '', 1, 0, 'C', '0', '0', 'cs:main:list', 'mdi:chat-processing-outline', 103, 1, NOW(), NULL, NULL, ''),
('2801', 'App品牌资源', '2800', 1, 'main', 'biz/appBrand/main/index', '', 1, 0, 'C', '0', '0', 'appBrand:main:list', 'mdi:earth', 103, 1, NOW(), NULL, NULL, ''),
('2901', 'App版本管理', '2900', 1, 'main', 'biz/appVersion/main/index', '', 1, 0, 'C', '0', '0', 'appVersion:main:list', 'mdi:view-grid', 103, 1, NOW(), NULL, NULL, '');

-- 赋权给系统管理员角色（admin）
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES
('228333863670124544', '2100'),
('228333863670124544', '2101'),
('228333863670124544', '2102'),
('228333863670124544', '2103'),
('228333863670124544', '2104'),
('228333863670124544', '2200'),
('228333863670124544', '2201'),
('228333863670124544', '2202'),
('228333863670124544', '2203'),
('228333863670124544', '2300'),
('228333863670124544', '2301'),
('228333863670124544', '2302'),
('228333863670124544', '2400'),
('228333863670124544', '2401'),
('228333863670124544', '2402'),
('228333863670124544', '2403'),
('228333863670124544', '2500'),
('228333863670124544', '2501'),
('228333863670124544', '2502'),
('228333863670124544', '2600'),
('228333863670124544', '2601'),
('228333863670124544', '2700'),
('228333863670124544', '2701'),
('228333863670124544', '2800'),
('228333863670124544', '2801'),
('228333863670124544', '2900'),
('228333863670124544', '2901');

COMMIT;
