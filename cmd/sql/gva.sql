/*
 Navicat MySQL Data Transfer

 Source Server         : 本地
 Source Server Type    : MySQL
 Source Server Version : 50734
 Source Host           : localhost:3306
 Source Schema         : gva

 Target Server Type    : MySQL
 Target Server Version : 50734
 File Encoding         : 65001

 Date: 21/12/2021 16:44:02
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for auto_code_examples
-- ----------------------------
DROP TABLE IF EXISTS `auto_code_examples`;
CREATE TABLE `auto_code_examples`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `auto_code_example_field` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '仅作示例条目无实际作用',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_auto_code_examples_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of auto_code_examples
-- ----------------------------

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`  (
  `p_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES ('p', '888', '/base/login', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/register', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/api/createApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/api/getApiList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/api/getApiById', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/api/deleteApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/api/updateApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/api/getAllApis', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/authority/createAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/authority/deleteAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/authority/getAuthorityList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/authority/setDataAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/authority/updateAuthority', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/authority/copyAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/getMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/getMenuList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/addBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/getBaseMenuTree', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/addMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/getMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/deleteBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/updateBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/getBaseMenuById', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/changePassword', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/getUserList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/setUserAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/deleteUser', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/casbin/updateCasbin', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/casbin/getPolicyPathByAuthorityId', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/casbin/casbinTest/:pathParam', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/jwt/jsonInBlacklist', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysOperationRecord/createSysOperationRecord', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysOperationRecord/deleteSysOperationRecord', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysOperationRecord/updateSysOperationRecord', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysOperationRecord/findSysOperationRecord', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysOperationRecord/getSysOperationRecordList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysOperationRecord/deleteSysOperationRecordByIds', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/setUserInfo', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/api/deleteApisByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/setUserAuthorities', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/getUserInfo', 'GET', '', '', '');

-- ----------------------------
-- Table structure for exa_customers
-- ----------------------------
DROP TABLE IF EXISTS `exa_customers`;
CREATE TABLE `exa_customers`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `customer_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '客户名',
  `customer_phone_data` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '客户手机号',
  `sys_user_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '管理ID',
  `sys_user_authority_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '管理角色ID',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_exa_customers_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of exa_customers
-- ----------------------------

-- ----------------------------
-- Table structure for exa_file_chunks
-- ----------------------------
DROP TABLE IF EXISTS `exa_file_chunks`;
CREATE TABLE `exa_file_chunks`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `exa_file_id` bigint(20) UNSIGNED NULL DEFAULT NULL,
  `file_chunk_number` bigint(20) NULL DEFAULT NULL,
  `file_chunk_path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_exa_file_chunks_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of exa_file_chunks
-- ----------------------------

-- ----------------------------
-- Table structure for exa_file_upload_and_downloads
-- ----------------------------
DROP TABLE IF EXISTS `exa_file_upload_and_downloads`;
CREATE TABLE `exa_file_upload_and_downloads`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '文件名',
  `url` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '文件地址',
  `tag` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '文件标签',
  `key` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '编号',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_exa_file_upload_and_downloads_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of exa_file_upload_and_downloads
-- ----------------------------
INSERT INTO `exa_file_upload_and_downloads` VALUES (1, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, '10.png', 'http://qmplusimg.henrongyi.top/gvalogo.png', 'png', '158787308910.png');
INSERT INTO `exa_file_upload_and_downloads` VALUES (2, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, 'logo.png', 'http://qmplusimg.henrongyi.top/1576554439myAvatar.png', 'png', '1587973709logo.png');

-- ----------------------------
-- Table structure for exa_files
-- ----------------------------
DROP TABLE IF EXISTS `exa_files`;
CREATE TABLE `exa_files`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `file_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `file_md5` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `file_path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `chunk_total` bigint(20) NULL DEFAULT NULL,
  `is_finish` tinyint(1) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_exa_files_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of exa_files
-- ----------------------------

-- ----------------------------
-- Table structure for exa_simple_uploaders
-- ----------------------------
DROP TABLE IF EXISTS `exa_simple_uploaders`;
CREATE TABLE `exa_simple_uploaders`  (
  `chunk_number` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '当前切片标记',
  `current_chunk_size` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '当前切片容量',
  `current_chunk_path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '切片本地路径',
  `total_size` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '总容量',
  `identifier` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '文件标识（md5）',
  `filename` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '文件名',
  `total_chunks` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '切片总数',
  `is_done` tinyint(1) NULL DEFAULT NULL COMMENT '是否上传完成',
  `file_path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '文件本地路径'
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of exa_simple_uploaders
-- ----------------------------

-- ----------------------------
-- Table structure for jwt_blacklists
-- ----------------------------
DROP TABLE IF EXISTS `jwt_blacklists`;
CREATE TABLE `jwt_blacklists`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `status` int(1) NULL DEFAULT NULL,
  `jwt` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT 'jwt',
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_jwt_blacklists_deleted_at`(`status`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of jwt_blacklists
-- ----------------------------
INSERT INTO `jwt_blacklists` VALUES (1, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiYzQyZDZlYzItM2RkMS00ODZmLTk5YmItNGMwZjczMDAyM2MxIiwiSUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiQnVmZmVyVGltZSI6ODY0MDAsImV4cCI6MTY0MDU5NTU1MiwiaXNzIjoicW1QbHVzIiwibmJmIjoxNjM5OTg5NzUyfQ.RfLuh1sTkPMAlj2A-ODDwZD3lgpyXPiV5X5FMAjhlZ0', NULL, NULL, NULL);
INSERT INTO `jwt_blacklists` VALUES (2, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiYzQyZDZlYzItM2RkMS00ODZmLTk5YmItNGMwZjczMDAyM2MxIiwiSUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiQnVmZmVyVGltZSI6ODY0MDAsImV4cCI6MTY0MDU5OTk1MiwiaXNzIjoicW1QbHVzIiwibmJmIjoxNjM5OTk0MTUyfQ.aaZAwNz78oyM038FdYEAQreqb7iHvphoR6AL-t4iYVs', NULL, NULL, NULL);
INSERT INTO `jwt_blacklists` VALUES (3, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiYzQyZDZlYzItM2RkMS00ODZmLTk5YmItNGMwZjczMDAyM2MxIiwiSUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiQnVmZmVyVGltZSI6ODY0MDAsImV4cCI6MTY0MDYwMDk3OCwiaXNzIjoicW1QbHVzIiwibmJmIjoxNjM5OTk1MTc4fQ.0MYaLphNHp3eZdaNo-FKtDjZCBJyqPzBbm_OcJHB8x8', NULL, NULL, NULL);
INSERT INTO `jwt_blacklists` VALUES (4, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiYzQyZDZlYzItM2RkMS00ODZmLTk5YmItNGMwZjczMDAyM2MxIiwiSUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiQnVmZmVyVGltZSI6ODY0MDAsImV4cCI6MTY0MDY1MzU4MiwiaXNzIjoicW1QbHVzIiwibmJmIjoxNjQwMDQ3NzgyfQ.T5YcPnxHiRA3DOeG_66E0RvN1L3Z0C-BLsbP1l2eOnw', NULL, NULL, NULL);
INSERT INTO `jwt_blacklists` VALUES (5, NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiYzQyZDZlYzItM2RkMS00ODZmLTk5YmItNGMwZjczMDAyM2MxIiwiSUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiQnVmZmVyVGltZSI6ODY0MDAsImV4cCI6MTY0MDY1NTU2NywiaXNzIjoicW1QbHVzIiwibmJmIjoxNjQwMDQ5NzY3fQ.CJWqydHa_2HL2WaFwleWhFLN8phOMIeQALbpkl7Fmyg', '2021-12-21 09:40:15', '2021-12-21 09:40:15', NULL);
INSERT INTO `jwt_blacklists` VALUES (6, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiYzQyZDZlYzItM2RkMS00ODZmLTk5YmItNGMwZjczMDAyM2MxIiwiSUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiQnVmZmVyVGltZSI6ODY0MDAsImV4cCI6MTY0MDY1NTYzMCwiaXNzIjoicW1QbHVzIiwibmJmIjoxNjQwMDQ5ODMwfQ.kuz6bAaqQbZahDpU_MDUFaQUWOzJ_k2LfyzPDjLfe6A', NULL, NULL, NULL);
INSERT INTO `jwt_blacklists` VALUES (7, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiYzQyZDZlYzItM2RkMS00ODZmLTk5YmItNGMwZjczMDAyM2MxIiwiSUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiQnVmZmVyVGltZSI6ODY0MDAsImV4cCI6MTY0MDY1NjA2NCwiaXNzIjoicW1QbHVzIiwibmJmIjoxNjQwMDUwMjY0fQ.XfUq_rxMJKt7Q-gNLETBkeYodQYkbX7rlGqZsWTAc7Q', NULL, NULL, NULL);
INSERT INTO `jwt_blacklists` VALUES (8, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiYzQyZDZlYzItM2RkMS00ODZmLTk5YmItNGMwZjczMDAyM2MxIiwiSUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiQnVmZmVyVGltZSI6ODY0MDAsImV4cCI6MTY0MDY1NzY2MiwiaXNzIjoicW1QbHVzIiwibmJmIjoxNjQwMDUxODYyfQ.0IOQPX5Yn6yiTjfkVp2CDoKDwPv9Q2yEqrsCP8eSGjE', NULL, NULL, NULL);
INSERT INTO `jwt_blacklists` VALUES (9, NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiYzQyZDZlYzItM2RkMS00ODZmLTk5YmItNGMwZjczMDAyM2MxIiwiSUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiQnVmZmVyVGltZSI6ODY0MDAsImV4cCI6MTY0MDY2MDk1NCwiaXNzIjoicW1QbHVzIiwibmJmIjoxNjQwMDU1MTU0fQ.Vub7EutIwXWbX-WY4VxLszDux76tb8Byo-xTImvZ0to', '2021-12-21 15:16:57', '2021-12-21 15:16:57', NULL);

-- ----------------------------
-- Table structure for sys_apis
-- ----------------------------
DROP TABLE IF EXISTS `sys_apis`;
CREATE TABLE `sys_apis`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'api路径',
  `description` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'api中文描述',
  `api_group` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'api组',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'POST' COMMENT '方法',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_apis_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 92 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_apis
-- ----------------------------
INSERT INTO `sys_apis` VALUES (1, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, '/base/login', '用户登录（必选）', 'base', 'POST');
INSERT INTO `sys_apis` VALUES (2, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, '/user/register', '用户注册（必选）', 'user', 'POST');
INSERT INTO `sys_apis` VALUES (3, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, '/api/createApi', '创建api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (4, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, '/api/getApiList', '获取api列表', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (5, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, '/api/getApiById', '获取api详细信息', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (6, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, '/api/deleteApi', '删除Api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (7, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, '/api/updateApi', '更新Api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (8, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, '/api/getAllApis', '获取所有api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (9, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, '/authority/createAuthority', '创建角色', 'authority', 'POST');
INSERT INTO `sys_apis` VALUES (10, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, '/authority/deleteAuthority', '删除角色', 'authority', 'POST');
INSERT INTO `sys_apis` VALUES (11, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, '/authority/getAuthorityList', '获取角色列表', 'authority', 'POST');
INSERT INTO `sys_apis` VALUES (12, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, '/menu/getMenu', '获取菜单树（必选）', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES (13, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, '/menu/getMenuList', '分页获取基础menu列表', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES (14, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, '/menu/addBaseMenu', '新增菜单', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES (15, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, '/menu/getBaseMenuTree', '获取用户动态路由', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES (16, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, '/menu/addMenuAuthority', '增加menu和角色关联关系', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES (17, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, '/menu/getMenuAuthority', '获取指定角色menu', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES (18, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, '/menu/deleteBaseMenu', '删除菜单', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES (19, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, '/menu/updateBaseMenu', '更新菜单', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES (20, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, '/menu/getBaseMenuById', '根据id获取菜单', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES (21, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, '/user/changePassword', '修改密码（建议选择）', 'user', 'POST');
INSERT INTO `sys_apis` VALUES (23, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, '/user/getUserList', '获取用户列表', 'user', 'POST');
INSERT INTO `sys_apis` VALUES (24, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, '/user/setUserAuthority', '修改用户角色（必选）', 'user', 'POST');
INSERT INTO `sys_apis` VALUES (27, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, '/casbin/updateCasbin', '更改角色api权限', 'casbin', 'POST');
INSERT INTO `sys_apis` VALUES (28, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, '/casbin/getPolicyPathByAuthorityId', '获取权限列表', 'casbin', 'POST');
INSERT INTO `sys_apis` VALUES (30, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, '/jwt/jsonInBlacklist', 'jwt加入黑名单(退出，必选)', 'jwt', 'POST');
INSERT INTO `sys_apis` VALUES (31, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, '/authority/setDataAuthority', '设置角色资源权限', 'authority', 'POST');
INSERT INTO `sys_apis` VALUES (39, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, '/casbin/casbinTest/:pathParam', 'RESTFUL模式测试', 'casbin', 'GET');
INSERT INTO `sys_apis` VALUES (41, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, '/authority/updateAuthority', '更新角色信息', 'authority', 'PUT');
INSERT INTO `sys_apis` VALUES (42, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, '/authority/copyAuthority', '拷贝角色', 'authority', 'POST');
INSERT INTO `sys_apis` VALUES (43, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, '/user/deleteUser', '删除用户', 'user', 'DELETE');
INSERT INTO `sys_apis` VALUES (54, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, '/sysOperationRecord/createSysOperationRecord', '新增操作记录', 'sysOperationRecord', 'POST');
INSERT INTO `sys_apis` VALUES (55, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, '/sysOperationRecord/deleteSysOperationRecord', '删除操作记录', 'sysOperationRecord', 'DELETE');
INSERT INTO `sys_apis` VALUES (56, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, '/sysOperationRecord/findSysOperationRecord', '根据ID获取操作记录', 'sysOperationRecord', 'GET');
INSERT INTO `sys_apis` VALUES (57, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, '/sysOperationRecord/getSysOperationRecordList', '获取操作记录列表', 'sysOperationRecord', 'GET');
INSERT INTO `sys_apis` VALUES (61, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, '/sysOperationRecord/deleteSysOperationRecordByIds', '批量删除操作历史', 'sysOperationRecord', 'DELETE');
INSERT INTO `sys_apis` VALUES (65, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, '/user/setUserInfo', '设置用户信息（必选）', 'user', 'PUT');
INSERT INTO `sys_apis` VALUES (85, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, '/api/deleteApisByIds', '批量删除api', 'api', 'DELETE');
INSERT INTO `sys_apis` VALUES (90, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, '/user/setUserAuthorities', '设置权限组', 'user', 'POST');
INSERT INTO `sys_apis` VALUES (91, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, '/user/getUserInfo', '获取自身信息（必选）', 'user', 'GET');

-- ----------------------------
-- Table structure for sys_authorities
-- ----------------------------
DROP TABLE IF EXISTS `sys_authorities`;
CREATE TABLE `sys_authorities`  (
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `authority_id` varchar(90) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色ID',
  `authority_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '角色名',
  `parent_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '父角色ID',
  `default_router` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'dashboard' COMMENT '默认菜单',
  PRIMARY KEY (`authority_id`) USING BTREE,
  UNIQUE INDEX `authority_id`(`authority_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_authorities
-- ----------------------------
INSERT INTO `sys_authorities` VALUES ('2021-12-20 16:49:10', '2021-12-21 10:49:33', NULL, '888', '普通用户', '0', 'dashboard');

-- ----------------------------
-- Table structure for sys_authority_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_authority_menus`;
CREATE TABLE `sys_authority_menus`  (
  `sys_base_menu_id` bigint(20) UNSIGNED NOT NULL,
  `sys_authority_authority_id` varchar(90) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_base_menu_id`, `sys_authority_authority_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_authority_menus
-- ----------------------------
INSERT INTO `sys_authority_menus` VALUES (1, '888');
INSERT INTO `sys_authority_menus` VALUES (3, '888');
INSERT INTO `sys_authority_menus` VALUES (4, '888');
INSERT INTO `sys_authority_menus` VALUES (5, '888');
INSERT INTO `sys_authority_menus` VALUES (6, '888');
INSERT INTO `sys_authority_menus` VALUES (7, '888');
INSERT INTO `sys_authority_menus` VALUES (20, '888');

-- ----------------------------
-- Table structure for sys_auto_code_histories
-- ----------------------------
DROP TABLE IF EXISTS `sys_auto_code_histories`;
CREATE TABLE `sys_auto_code_histories`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `table_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `request_meta` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `auto_code_path` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `injection_meta` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `struct_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `struct_cn_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `api_ids` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `flag` bigint(20) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_auto_code_histories_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_auto_code_histories
-- ----------------------------

-- ----------------------------
-- Table structure for sys_base_menu_parameters
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menu_parameters`;
CREATE TABLE `sys_base_menu_parameters`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `sys_base_menu_id` bigint(20) UNSIGNED NULL DEFAULT NULL,
  `type` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地址栏携带参数为params还是query',
  `key` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地址栏携带参数的key',
  `value` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地址栏携带参数的值',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_base_menu_parameters_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_base_menu_parameters
-- ----------------------------

-- ----------------------------
-- Table structure for sys_base_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menus`;
CREATE TABLE `sys_base_menus`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `menu_level` bigint(20) UNSIGNED NULL DEFAULT NULL,
  `parent_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '父菜单ID',
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '路由path',
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '路由name',
  `hidden` tinyint(1) NULL DEFAULT NULL COMMENT '是否在列表隐藏',
  `component` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '对应前端文件路径',
  `sort` bigint(20) NULL DEFAULT NULL COMMENT '排序标记',
  `keep_alive` tinyint(1) NULL DEFAULT NULL COMMENT '附加属性',
  `default_menu` tinyint(1) NULL DEFAULT NULL COMMENT '附加属性',
  `title` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '附加属性',
  `icon` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '附加属性',
  `close_tab` tinyint(1) NULL DEFAULT NULL COMMENT '附加属性',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_base_menus_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 26 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_base_menus
-- ----------------------------
INSERT INTO `sys_base_menus` VALUES (1, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, 0, '0', 'dashboard', 'dashboard', 0, 'view/dashboard/index.vue', 2, 0, 0, '仪表盘', 's-home', 0);
INSERT INTO `sys_base_menus` VALUES (3, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, 0, '0', 'admin', 'superAdmin', 0, 'view/superAdmin/index.vue', 3, 0, 0, '超级管理员', 'user-solid', 0);
INSERT INTO `sys_base_menus` VALUES (4, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, 0, '3', 'authority', 'authority', 0, 'view/superAdmin/authority/authority.vue', 1, 0, 0, '角色管理', 's-custom', 0);
INSERT INTO `sys_base_menus` VALUES (5, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, 0, '3', 'menu', 'menu', 0, 'view/superAdmin/menu/menu.vue', 2, 1, 0, '菜单管理', 's-order', 0);
INSERT INTO `sys_base_menus` VALUES (6, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, 0, '3', 'api', 'api', 0, 'view/superAdmin/api/api.vue', 3, 1, 0, 'api管理', 's-platform', 0);
INSERT INTO `sys_base_menus` VALUES (7, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, 0, '3', 'user', 'user', 0, 'view/superAdmin/user/user.vue', 4, 0, 0, '用户管理', 'coordinate', 0);
INSERT INTO `sys_base_menus` VALUES (20, '2021-12-20 16:49:10', '2021-12-20 16:49:10', NULL, 0, '3', 'operation', 'operation', 0, 'view/superAdmin/operation/sysOperationRecord.vue', 6, 0, 0, '操作历史', 'time', 0);

-- ----------------------------
-- Table structure for sys_data_authority_id
-- ----------------------------
DROP TABLE IF EXISTS `sys_data_authority_id`;
CREATE TABLE `sys_data_authority_id`  (
  `sys_authority_authority_id` varchar(90) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色ID',
  `data_authority_id_authority_id` varchar(90) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_authority_authority_id`, `data_authority_id_authority_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_data_authority_id
-- ----------------------------
INSERT INTO `sys_data_authority_id` VALUES ('888', '888');
INSERT INTO `sys_data_authority_id` VALUES ('888', '8881');
INSERT INTO `sys_data_authority_id` VALUES ('888', '9528');
INSERT INTO `sys_data_authority_id` VALUES ('9528', '8881');
INSERT INTO `sys_data_authority_id` VALUES ('9528', '9528');

-- ----------------------------
-- Table structure for sys_dictionaries
-- ----------------------------
DROP TABLE IF EXISTS `sys_dictionaries`;
CREATE TABLE `sys_dictionaries`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '字典名（中）',
  `type` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '字典名（英）',
  `status` tinyint(1) NULL DEFAULT NULL COMMENT '状态',
  `desc` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_dictionaries_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dictionaries
-- ----------------------------
INSERT INTO `sys_dictionaries` VALUES (1, '2021-12-20 16:59:04', '2021-12-20 16:59:04', NULL, '性别', 'sex', 1, '性别字典');
INSERT INTO `sys_dictionaries` VALUES (2, '2021-12-20 16:59:04', '2021-12-20 16:59:04', NULL, '数据库int类型', 'int', 1, 'int类型对应的数据库类型');
INSERT INTO `sys_dictionaries` VALUES (3, '2021-12-20 16:59:04', '2021-12-20 16:59:04', NULL, '数据库时间日期类型', 'time.Time', 1, '数据库时间日期类型');
INSERT INTO `sys_dictionaries` VALUES (4, '2021-12-20 16:59:04', '2021-12-20 16:59:04', NULL, '数据库浮点型', 'float64', 1, '数据库浮点型');
INSERT INTO `sys_dictionaries` VALUES (5, '2021-12-20 16:59:04', '2021-12-20 16:59:04', NULL, '数据库字符串', 'string', 1, '数据库字符串');
INSERT INTO `sys_dictionaries` VALUES (6, '2021-12-20 16:59:04', '2021-12-20 16:59:04', NULL, '数据库bool类型', 'bool', 1, '数据库bool类型');

-- ----------------------------
-- Table structure for sys_dictionary_details
-- ----------------------------
DROP TABLE IF EXISTS `sys_dictionary_details`;
CREATE TABLE `sys_dictionary_details`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `label` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '展示值',
  `value` bigint(20) NULL DEFAULT NULL COMMENT '字典值',
  `status` tinyint(1) NULL DEFAULT NULL COMMENT '启用状态',
  `sort` bigint(20) NULL DEFAULT NULL COMMENT '排序标记',
  `sys_dictionary_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '关联标记',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_dictionary_details_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 24 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dictionary_details
-- ----------------------------
INSERT INTO `sys_dictionary_details` VALUES (1, '2021-12-20 16:59:04', '2021-12-20 16:59:04', NULL, 'smallint', 1, 1, 1, 2);
INSERT INTO `sys_dictionary_details` VALUES (2, '2021-12-20 16:59:04', '2021-12-20 16:59:04', NULL, 'mediumint', 2, 1, 2, 2);
INSERT INTO `sys_dictionary_details` VALUES (3, '2021-12-20 16:59:04', '2021-12-20 16:59:04', NULL, 'int', 3, 1, 3, 2);
INSERT INTO `sys_dictionary_details` VALUES (4, '2021-12-20 16:59:04', '2021-12-20 16:59:04', NULL, 'bigint', 4, 1, 4, 2);
INSERT INTO `sys_dictionary_details` VALUES (5, '2021-12-20 16:59:04', '2021-12-20 16:59:04', NULL, 'date', 0, 1, 0, 3);
INSERT INTO `sys_dictionary_details` VALUES (6, '2021-12-20 16:59:04', '2021-12-20 16:59:04', NULL, 'time', 1, 1, 1, 3);
INSERT INTO `sys_dictionary_details` VALUES (7, '2021-12-20 16:59:04', '2021-12-20 16:59:04', NULL, 'year', 2, 1, 2, 3);
INSERT INTO `sys_dictionary_details` VALUES (8, '2021-12-20 16:59:04', '2021-12-20 16:59:04', NULL, 'datetime', 3, 1, 3, 3);
INSERT INTO `sys_dictionary_details` VALUES (9, '2021-12-20 16:59:04', '2021-12-20 16:59:04', NULL, 'timestamp', 5, 1, 5, 3);
INSERT INTO `sys_dictionary_details` VALUES (10, '2021-12-20 16:59:04', '2021-12-20 16:59:04', NULL, 'float', 0, 1, 0, 4);
INSERT INTO `sys_dictionary_details` VALUES (11, '2021-12-20 16:59:04', '2021-12-20 16:59:04', NULL, 'double', 1, 1, 1, 4);
INSERT INTO `sys_dictionary_details` VALUES (12, '2021-12-20 16:59:04', '2021-12-20 16:59:04', NULL, 'decimal', 2, 1, 2, 4);
INSERT INTO `sys_dictionary_details` VALUES (13, '2021-12-20 16:59:04', '2021-12-20 16:59:04', NULL, 'char', 0, 1, 0, 5);
INSERT INTO `sys_dictionary_details` VALUES (14, '2021-12-20 16:59:04', '2021-12-20 16:59:04', NULL, 'varchar', 1, 1, 1, 5);
INSERT INTO `sys_dictionary_details` VALUES (15, '2021-12-20 16:59:04', '2021-12-20 16:59:04', NULL, 'tinyblob', 2, 1, 2, 5);
INSERT INTO `sys_dictionary_details` VALUES (16, '2021-12-20 16:59:04', '2021-12-20 16:59:04', NULL, 'tinytext', 3, 1, 3, 5);
INSERT INTO `sys_dictionary_details` VALUES (17, '2021-12-20 16:59:04', '2021-12-20 16:59:04', NULL, 'text', 4, 1, 4, 5);
INSERT INTO `sys_dictionary_details` VALUES (18, '2021-12-20 16:59:04', '2021-12-20 16:59:04', NULL, 'blob', 5, 1, 5, 5);
INSERT INTO `sys_dictionary_details` VALUES (19, '2021-12-20 16:59:04', '2021-12-20 16:59:04', NULL, 'mediumblob', 6, 1, 6, 5);
INSERT INTO `sys_dictionary_details` VALUES (20, '2021-12-20 16:59:04', '2021-12-20 16:59:04', NULL, 'mediumtext', 7, 1, 7, 5);
INSERT INTO `sys_dictionary_details` VALUES (21, '2021-12-20 16:59:04', '2021-12-20 16:59:04', NULL, 'longblob', 8, 1, 8, 5);
INSERT INTO `sys_dictionary_details` VALUES (22, '2021-12-20 16:59:04', '2021-12-20 16:59:04', NULL, 'longtext', 9, 1, 9, 5);
INSERT INTO `sys_dictionary_details` VALUES (23, '2021-12-20 16:59:04', '2021-12-20 16:59:04', NULL, 'tinyint', 0, 1, 0, 6);

-- ----------------------------
-- Table structure for sys_operation_records
-- ----------------------------
DROP TABLE IF EXISTS `sys_operation_records`;
CREATE TABLE `sys_operation_records`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `ip` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求ip',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求方法',
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求路径',
  `status` bigint(20) NULL DEFAULT NULL COMMENT '请求状态',
  `agent` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '代理',
  `error_message` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '错误信息',
  `body` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '请求Body',
  `user_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '用户id',
  `latency` bigint(20) NULL DEFAULT NULL COMMENT '延迟',
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `resp` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '响应Body',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_operation_records_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 217 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_operation_records
-- ----------------------------
INSERT INTO `sys_operation_records` VALUES (204, '2021-12-21 16:14:04', '127.0.0.1:8888', 'POST', '/sysOperationRecord/getSysOperationRecordList', 1, 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.93 Safari/537.36', '', '{\"page\":1,\"pageSize\":50,\"path\":\"\",\"method\":\"\"}', 1, NULL, NULL, NULL, NULL);
INSERT INTO `sys_operation_records` VALUES (205, '2021-12-21 16:14:12', '127.0.0.1:8888', 'POST', '/sysOperationRecord/deleteSysOperationRecordByIds', 1, 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.93 Safari/537.36', '', '{\"ids\":[203,202,201,40,39,38,37,36,35,34,33,32,31,30,29,28,27,26,25,24,23,22,21,20,19,18,17,16,15,14,13,12,11,10,9,8,7,6,5,4,3]}', 1, NULL, NULL, NULL, NULL);
INSERT INTO `sys_operation_records` VALUES (206, '2021-12-21 16:14:13', '127.0.0.1:8888', 'POST', '/sysOperationRecord/getSysOperationRecordList', 1, 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.93 Safari/537.36', '', '{\"page\":1,\"pageSize\":50,\"path\":\"\",\"method\":\"\"}', 1, NULL, NULL, NULL, NULL);
INSERT INTO `sys_operation_records` VALUES (207, '2021-12-21 16:38:41', '127.0.0.1:8888', 'POST', '/menu/getMenu', 1, 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.93 Safari/537.36', '', '', 1, NULL, NULL, NULL, NULL);
INSERT INTO `sys_operation_records` VALUES (208, '2021-12-21 16:38:41', '127.0.0.1:8888', 'GET', '/user/getUserInfo', 1, 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.93 Safari/537.36', '', '', 1, NULL, NULL, NULL, NULL);
INSERT INTO `sys_operation_records` VALUES (209, '2021-12-21 16:38:41', '127.0.0.1:8888', 'POST', '/menu/getMenuList', 1, 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.93 Safari/537.36', '', '{\"page\":1,\"pageSize\":999,\"path\":\"\",\"method\":\"\"}', 1, NULL, NULL, NULL, NULL);
INSERT INTO `sys_operation_records` VALUES (210, '2021-12-21 16:38:52', '127.0.0.1:8888', 'POST', '/authority/getAuthorityList', 1, 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.93 Safari/537.36', '', '{\"page\":1,\"pageSize\":999,\"path\":\"\",\"method\":\"\"}', 1, NULL, NULL, NULL, NULL);
INSERT INTO `sys_operation_records` VALUES (211, '2021-12-21 16:39:02', '127.0.0.1:8888', 'POST', '/user/getUserList', 1, 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.93 Safari/537.36', '', '{\"page\":1,\"pageSize\":10,\"path\":\"\",\"method\":\"\"}', 1, NULL, NULL, NULL, NULL);
INSERT INTO `sys_operation_records` VALUES (212, '2021-12-21 16:39:02', '127.0.0.1:8888', 'POST', '/authority/getAuthorityList', 1, 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.93 Safari/537.36', '', '{\"page\":1,\"pageSize\":999}', 1, NULL, NULL, NULL, NULL);
INSERT INTO `sys_operation_records` VALUES (213, '2021-12-21 16:39:08', '127.0.0.1:8888', 'POST', '/user/setUserAuthorities', 1, 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.93 Safari/537.36', '', '{\"ID\":1,\"authorityIds\":[\"888\"]}', 1, NULL, NULL, NULL, NULL);
INSERT INTO `sys_operation_records` VALUES (214, '2021-12-21 16:39:21', '127.0.0.1:8888', 'POST', '/user/setUserAuthorities', 1, 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.93 Safari/537.36', '', '{\"ID\":1,\"authorityIds\":[]}', 1, NULL, NULL, NULL, NULL);
INSERT INTO `sys_operation_records` VALUES (215, '2021-12-21 16:39:24', '127.0.0.1:8888', 'POST', '/user/getUserList', 1, 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.93 Safari/537.36', '', '{\"page\":1,\"pageSize\":10,\"path\":\"\",\"method\":\"\"}', 1, NULL, NULL, NULL, NULL);
INSERT INTO `sys_operation_records` VALUES (216, '2021-12-21 16:39:24', '127.0.0.1:8888', 'POST', '/authority/getAuthorityList', 1, 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.93 Safari/537.36', '', '{\"page\":1,\"pageSize\":999}', 1, NULL, NULL, NULL, NULL);

-- ----------------------------
-- Table structure for sys_user_authority
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_authority`;
CREATE TABLE `sys_user_authority`  (
  `sys_user_id` bigint(20) UNSIGNED NOT NULL,
  `sys_authority_authority_id` varchar(90) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_user_id`, `sys_authority_authority_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user_authority
-- ----------------------------
INSERT INTO `sys_user_authority` VALUES (1, '888');

-- ----------------------------
-- Table structure for sys_users
-- ----------------------------
DROP TABLE IF EXISTS `sys_users`;
CREATE TABLE `sys_users`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `uuid` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户UUID',
  `username` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户登录名',
  `password` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户登录密码',
  `nick_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '系统用户' COMMENT '用户昵称',
  `header_img` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'http://qmplusimg.henrongyi.top/head.png' COMMENT '用户头像',
  `authority_id` varchar(90) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '888' COMMENT '用户角色ID',
  `side_mode` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'dark' COMMENT '用户角色ID',
  `active_color` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '#1890ff' COMMENT '用户角色ID',
  `base_color` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '#fff' COMMENT '用户角色ID',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_users_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_users
-- ----------------------------
INSERT INTO `sys_users` VALUES (1, '2021-12-20 16:49:10', '2021-12-21 15:24:41', NULL, 'c42d6ec2-3dd1-486f-99bb-4c0f730023c1', 'admin', 'e10adc3949ba59abbe56e057f20f883e', '超级管理员', 'http://qmplusimg.henrongyi.top/gva_header.jpg', '888', 'dark', '#1890ff', '#fff');

-- ----------------------------
-- View structure for authority_menu
-- ----------------------------
DROP VIEW IF EXISTS `authority_menu`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `authority_menu` AS select `sys_base_menus`.`id` AS `id`,`sys_base_menus`.`created_at` AS `created_at`,`sys_base_menus`.`updated_at` AS `updated_at`,`sys_base_menus`.`deleted_at` AS `deleted_at`,`sys_base_menus`.`menu_level` AS `menu_level`,`sys_base_menus`.`parent_id` AS `parent_id`,`sys_base_menus`.`path` AS `path`,`sys_base_menus`.`name` AS `name`,`sys_base_menus`.`hidden` AS `hidden`,`sys_base_menus`.`component` AS `component`,`sys_base_menus`.`title` AS `title`,`sys_base_menus`.`icon` AS `icon`,`sys_base_menus`.`sort` AS `sort`,`sys_authority_menus`.`sys_authority_authority_id` AS `authority_id`,`sys_authority_menus`.`sys_base_menu_id` AS `menu_id`,`sys_base_menus`.`keep_alive` AS `keep_alive`,`sys_base_menus`.`close_tab` AS `close_tab`,`sys_base_menus`.`default_menu` AS `default_menu` from (`sys_authority_menus` join `sys_base_menus` on((`sys_authority_menus`.`sys_base_menu_id` = `sys_base_menus`.`id`)));

SET FOREIGN_KEY_CHECKS = 1;
