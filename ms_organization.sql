CREATE TABLE `ms_organization` (
    `id` bigint(0) NOT NULL AUTO_INCREMENT,
    `name` varchar(255) CHARACTER SET utf8mb4 COLLATE  utf8mb4_general_ci NULL DEFAULT '' COMMENT '名称',
    `avatar` VARCHAR(511) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '头像',
    `description` VARCHAR(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  NULL COMMENT '备注',
    `member_id` bigint(0) NULL DEFAULT NULL COMMENT '拥有者',
    `create_time` bigint(0) NULL DEFAULT NULL COMMENT '创建时间',
    `personal` tinyint(1) NULL DEFAULT 0 COMMENT '是否个人项目',
    `address` varchar(100) CHARACTER SET utf8mb4 COLLATE  utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地址',
    `province` int(0) NULL DEFAULT 0 COMMENT '省',
    `city` int(0) NULL DEFAULT 0 COMMENT '市',
    `area` int(0) NULL DEFAULT 0 COMMENT '区',
    PRIMARY KEY(`id`) USING BTREE
) ENGINE = INNODB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '组织表' ROW_FORMAT = COMPACT;