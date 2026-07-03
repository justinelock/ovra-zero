-- 各表登录账号/联系人列统一命名（已执行过的语句可跳过）

ALTER TABLE `sys_user`
  CHANGE COLUMN `user_name` `username` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户账号';

ALTER TABLE `sys_logininfor`
  CHANGE COLUMN `user_name` `username` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '用户账号';

ALTER TABLE `sys_social`
  CHANGE COLUMN `user_name` `username` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '登录账号';

ALTER TABLE `sys_tenant`
  CHANGE COLUMN `contact_user_name` `contact_username` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '联系人';
