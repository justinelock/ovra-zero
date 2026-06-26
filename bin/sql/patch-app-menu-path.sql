-- App 管理子菜单路由冲突修复：两条菜单 path 均为 main → 均解析为 /app/main
-- 执行后请重新登录或刷新菜单缓存

UPDATE sys_menu SET path = 'brand', order_num = 1 WHERE menu_id = '2801';
UPDATE sys_menu SET path = 'version', order_num = 2 WHERE menu_id = '2901';
