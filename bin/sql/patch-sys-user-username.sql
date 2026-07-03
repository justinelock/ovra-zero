-- sys_user 登录账号列与线上一致：user_name -> username
ALTER TABLE `sys_user`
  CHANGE COLUMN `user_name` `username` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户账号';
