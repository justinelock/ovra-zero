-- 业务用户管理一级菜单（与「系统管理 → 用户管理」并存）
-- 已有库执行本脚本；全新安装已含在 ovra_zero-2.0.sql

BEGIN;

-- 一级目录
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_dept`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES
('2000', '用户管理', '0', 8, 'member', NULL, '', 1, 0, 'M', '0', '0', '', 'ant-design:user-outlined', 103, 1, NOW(), NULL, NULL, '业务用户中心');

-- 二级页面（component 对应 views/member/...）
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_dept`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES
('2001', '用户列表', '2000', 1, 'list', 'member/list/index', '', 1, 0, 'C', '0', '0', 'member:list:list', 'ant-design:user-outlined', 103, 1, NOW(), NULL, NULL, ''),
('2002', '实名认证', '2000', 2, 'kyc', 'member/kyc/index', '', 1, 0, 'C', '0', '0', 'member:kyc:list', 'mdi:card-account-details-outline', 103, 1, NOW(), NULL, NULL, ''),
('2003', '钱包管理', '2000', 3, 'wallet', 'member/wallet/index', '', 1, 0, 'C', '0', '0', 'member:wallet:list', 'ant-design:wallet-outlined', 103, 1, NOW(), NULL, NULL, ''),
('2004', '用户报表', '2000', 4, 'report', 'member/report/index', '', 1, 0, 'C', '0', '0', 'member:report:list', 'mdi:chart-bar', 103, 1, NOW(), NULL, NULL, ''),
('2005', '团队管理', '2000', 5, 'team', 'member/team/index', '', 1, 0, 'C', '0', '0', 'member:team:list', 'mdi:account-group-outline', 103, 1, NOW(), NULL, NULL, ''),
('2006', '登录记录', '2000', 6, 'login-log', 'member/loginLog/index', '', 1, 0, 'C', '0', '0', 'member:loginLog:list', 'mdi:login', 103, 1, NOW(), NULL, NULL, '');

-- 赋权给系统管理员角色（admin）
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES
('228333863670124544', '2000'),
('228333863670124544', '2001'),
('228333863670124544', '2002'),
('228333863670124544', '2003'),
('228333863670124544', '2004'),
('228333863670124544', '2005'),
('228333863670124544', '2006');

COMMIT;
