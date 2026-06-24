-- 开发参考菜单：挂载「演示使用自行删除」前端示例页
-- 已有库执行本脚本；全新安装已含在 ovra_zero-2.0.sql

BEGIN;

-- 一级目录
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_dept`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES
('7', '开发参考', '0', 7, 'dev-ref', NULL, '', 1, 0, 'M', '0', '0', '', 'lucide:book-open', 103, 1, NOW(), NULL, NULL, 'Vben 组件与能力演示，上线前可删');

-- 二级页面（component 对应 views/演示使用自行删除/...）
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_dept`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES
('1710', '表单示例', '7', 1, 'form', '演示使用自行删除/form/index', '', 1, 0, 'C', '0', '0', 'devref:form:list', 'lucide:file-input', 103, 1, NOW(), NULL, NULL, ''),
('1711', '查询示例', '7', 2, 'query', '演示使用自行删除/query/index', '', 1, 0, 'C', '0', '0', 'devref:query:list', 'lucide:search', 103, 1, NOW(), NULL, NULL, ''),
('1712', '字典示例', '7', 3, 'dict', '演示使用自行删除/dict/index', '', 1, 0, 'C', '0', '0', 'devref:dict:list', 'fluent-mdl2:dictionary', 103, 1, NOW(), NULL, NULL, ''),
('1713', '菜单示例', '7', 4, 'menu', '演示使用自行删除/menu/index', '', 1, 0, 'C', '0', '0', 'devref:menu:list', 'ic:sharp-menu', 103, 1, NOW(), NULL, NULL, ''),
('1714', 'Vxe表格', '7', 5, 'vxe', '演示使用自行删除/vxe/index', '', 1, 0, 'C', '0', '0', 'devref:vxe:list', 'lucide:table', 103, 1, NOW(), NULL, NULL, ''),
('1715', '富文本', '7', 6, 'tinymce', '演示使用自行删除/tinymce/index', '', 1, 0, 'C', '0', '0', 'devref:tinymce:list', 'lucide:type', 103, 1, NOW(), NULL, NULL, ''),
('1716', '文件上传', '7', 7, 'upload', '演示使用自行删除/upload/index', '', 1, 0, 'C', '0', '0', 'devref:upload:list', 'lucide:upload', 103, 1, NOW(), NULL, NULL, ''),
('1717', 'SSE推送', '7', 8, 'sse', '演示使用自行删除/sse/index', '', 1, 0, 'C', '0', '0', 'devref:sse:list', 'lucide:radio', 103, 1, NOW(), NULL, NULL, ''),
('1718', 'API加解密', '7', 9, 'encrypt', '演示使用自行删除/other/encrypt', '', 1, 0, 'C', '0', '0', 'devref:encrypt:list', 'lucide:lock', 103, 1, NOW(), NULL, NULL, ''),
('1719', '微信示例', '7', 10, 'wechat', '演示使用自行删除/wechat/index', '', 1, 0, 'C', '0', '0', 'devref:wechat:list', 'mdi:wechat', 103, 1, NOW(), NULL, NULL, ''),
('1720', '更新记录', '7', 11, 'changelog', '演示使用自行删除/changelog/index', '', 1, 0, 'C', '0', '0', 'devref:changelog:list', 'lucide:scroll-text', 103, 1, NOW(), NULL, NULL, '');

-- SSE 演示按钮权限
INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `parent_id`, `order_num`, `path`, `component`, `query_param`, `is_frame`, `is_cache`, `menu_type`, `visible`, `status`, `perms`, `icon`, `create_dept`, `create_by`, `create_time`, `update_by`, `update_time`, `remark`) VALUES
('1721', 'SSE查询', '1717', 1, '#', '', '', 1, 0, 'F', '0', '0', 'devref:sse:query', '#', 103, 1, NOW(), NULL, NULL, ''),
('1722', 'SSE发送', '1717', 2, '#', '', '', 1, 0, 'F', '0', '0', 'devref:sse:send', '#', 103, 1, NOW(), NULL, NULL, '');

-- 赋权给系统管理员角色（admin 默认角色）
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES
('228333863670124544', '7'),
('228333863670124544', '1710'),
('228333863670124544', '1711'),
('228333863670124544', '1712'),
('228333863670124544', '1713'),
('228333863670124544', '1714'),
('228333863670124544', '1715'),
('228333863670124544', '1716'),
('228333863670124544', '1717'),
('228333863670124544', '1718'),
('228333863670124544', '1719'),
('228333863670124544', '1720'),
('228333863670124544', '1721'),
('228333863670124544', '1722');

COMMIT;
