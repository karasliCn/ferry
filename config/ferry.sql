/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80023
 Source Host           : localhost
 Source Database       : ferry

 Target Server Type    : MySQL
 Target Server Version : 80023
 File Encoding         : utf-8

 Date: 08/05/2021 15:07:42 PM
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `casbin_rule`
--  接口权限配置（不入仓）
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule` (
  `p_type` varchar(100) DEFAULT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL,
  `id` int NOT NULL AUTO_INCREMENT,
  `create_time` timestamp NULL DEFAULT NULL,
  `update_time` timestamp NULL DEFAULT NULL,
  `delete_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_casbin_rule_delete_time` (`delete_time`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
--  Table structure for `p_process_classify`
--  流程分类
-- ----------------------------
DROP TABLE IF EXISTS `p_process_classify`;
CREATE TABLE `p_process_classify` (
  `id` int NOT NULL AUTO_INCREMENT,
  `create_time` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `delete_time` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `name` varchar(128) DEFAULT NULL COMMENT '分类名称',
  `creator` int DEFAULT NULL COMMENT '创建人',
  PRIMARY KEY (`id`),
  KEY `idx_p_process_classify_delete_time` (`delete_time`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
--  Table structure for `p_process_info`
--  流程信息
-- ----------------------------
DROP TABLE IF EXISTS `p_process_info`;
CREATE TABLE `p_process_info` (
  `id` int NOT NULL AUTO_INCREMENT,
  `create_time` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `delete_time` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `name` varchar(128) DEFAULT NULL COMMENT '流程名称',
  `icon` varchar(128) DEFAULT NULL COMMENT '流程图标',
  `structure` json DEFAULT NULL COMMENT '流程结构',
  `classify` int DEFAULT NULL COMMENT '流程分类',
  `tpls` json DEFAULT NULL COMMENT '模版列表',
  `task` json DEFAULT NULL COMMENT '异步任务',
  `submit_count` int DEFAULT '0' COMMENT '已创建工单数量',
  `creator` int DEFAULT NULL COMMENT '创建人',
  `notice` json DEFAULT NULL COMMENT '通知列表',
  `remarks` varchar(1024) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `idx_p_process_info_delete_time` (`delete_time`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
--  Table structure for `p_task_history`
--  异步任务执行历史
-- ----------------------------
DROP TABLE IF EXISTS `p_task_history`;
CREATE TABLE `p_task_history` (
  `id` int NOT NULL AUTO_INCREMENT,
  `create_time` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `delete_time` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `task` int DEFAULT NULL COMMENT '任务id',
  `name` varchar(256) DEFAULT NULL COMMENT '任务名称',
  `task_type` int DEFAULT NULL COMMENT '任务类型',
  `execution_time` varchar(128) DEFAULT NULL COMMENT '执行时间',
  `result` longtext COMMENT '执行结果',
  PRIMARY KEY (`id`),
  KEY `idx_p_task_history_delete_time` (`delete_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
--  Table structure for `p_task_info`
--  异步任务信息
-- ----------------------------
DROP TABLE IF EXISTS `p_task_info`;
CREATE TABLE `p_task_info` (
  `id` int NOT NULL AUTO_INCREMENT,
  `create_time` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `delete_time` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `name` varchar(256) DEFAULT NULL COMMENT '任务名称',
  `task_type` varchar(45) DEFAULT NULL COMMENT '任务类型',
  `content` longtext COMMENT '任务内容',
  `creator` int DEFAULT NULL COMMENT '创建人',
  `remarks` longtext COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `idx_p_task_info_delete_time` (`delete_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
--  Table structure for `p_tpl_info`
--  模版信息
-- ----------------------------
DROP TABLE IF EXISTS `p_tpl_info`;
CREATE TABLE `p_tpl_info` (
  `id` int NOT NULL AUTO_INCREMENT,
  `create_time` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `delete_time` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `name` varchar(128) DEFAULT NULL COMMENT '模版名称',
  `form_structure` json DEFAULT NULL COMMENT '模版结构',
  `creator` int DEFAULT NULL COMMENT '创建人',
  `remarks` longtext COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `idx_p_tpl_info_delete_time` (`delete_time`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
--  Table structure for `p_work_order_circulation_history`
--  工单流转历史
-- ----------------------------
DROP TABLE IF EXISTS `p_work_order_circulation_history`;
CREATE TABLE `p_work_order_circulation_history` (
  `id` int NOT NULL AUTO_INCREMENT,
  `create_time` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `delete_time` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `suspend_time` timestamp NULL DEFAULT NULL COMMENT '挂起时间',
  `resume_time` timestamp NULL DEFAULT NULL COMMENT '恢复时间',
  `title` varchar(128) DEFAULT NULL COMMENT '标题',
  `work_order` int DEFAULT NULL COMMENT '工单号',
  `state` varchar(128) DEFAULT NULL COMMENT '工单状态',
  `source` varchar(128) DEFAULT NULL COMMENT '流转起始节点',
  `target` varchar(128) DEFAULT NULL COMMENT '流转目标节点',
  `circulation` varchar(128) DEFAULT NULL COMMENT '流转路径',
  `status` int DEFAULT NULL COMMENT '流转状态 - 0：拒绝 1：同意',
  `processor` varchar(45) DEFAULT NULL COMMENT '处理人',
  `processor_id` int DEFAULT NULL COMMENT '处理人id',
  `cost_duration` int DEFAULT NULL COMMENT '节点处理时间',
  `remarks` longtext COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `idx_p_work_order_circulation_history_delete_time` (`delete_time`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
--  Table structure for `p_work_order_info`
--  工单信息
-- ----------------------------
DROP TABLE IF EXISTS `p_work_order_info`;
CREATE TABLE `p_work_order_info` (
  `id` int NOT NULL AUTO_INCREMENT,
  `create_time` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `delete_time` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `title` varchar(128) DEFAULT NULL COMMENT '工单标题',
  `priority` int DEFAULT NULL COMMENT '优先级',
  `process` int DEFAULT NULL COMMENT '流程id',
  `classify` int DEFAULT NULL COMMENT '流程分类',
  `is_end` int DEFAULT '0' COMMENT '是否结束',
  `is_denied` int DEFAULT '0' COMMENT '是否拒绝',
  `state` json DEFAULT NULL COMMENT '当前状态',
  `related_person` json DEFAULT NULL COMMENT '相关人员id',
  `creator` int DEFAULT NULL COMMENT '创建人',
  `urge_count` int DEFAULT '0' COMMENT '催办次数',
  `urge_last_time` int DEFAULT '0' COMMENT '上一次催办时间',
  PRIMARY KEY (`id`),
  KEY `idx_p_work_order_info_delete_time` (`delete_time`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
--  Table structure for `p_work_order_tpl_data`
--  工单表单数据
-- ----------------------------
DROP TABLE IF EXISTS `p_work_order_tpl_data`;
CREATE TABLE `p_work_order_tpl_data` (
  `id` int NOT NULL AUTO_INCREMENT,
  `create_time` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `delete_time` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `work_order` int DEFAULT NULL COMMENT '工单号',
  `form_structure` json DEFAULT NULL COMMENT '表单结构',
  `form_data` json DEFAULT NULL COMMENT '表单数据',
  PRIMARY KEY (`id`),
  KEY `idx_p_work_order_tpl_data_delete_time` (`delete_time`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
--  Table structure for `sys_dept`
--  部门表（不入仓？）
-- ----------------------------
DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept` (
  `dept_id` int NOT NULL AUTO_INCREMENT COMMENT '部门id',
  `parent_id` int DEFAULT NULL COMMENT '上级id',
  `dept_path` varchar(255) DEFAULT NULL COMMENT '部门层级路径',
  `dept_name` varchar(128) DEFAULT NULL COMMENT '部门名称',
  `sort` int DEFAULT NULL COMMENT '排序号',
  `leader` int DEFAULT NULL COMMENT '部门管理员',
  `phone` varchar(11) DEFAULT NULL COMMENT '手机号',
  `email` varchar(64) DEFAULT NULL COMMENT '邮箱',
  `status` int DEFAULT NULL COMMENT '部门状态',
  `create_by` varchar(64) DEFAULT NULL COMMENT '创建人',
  `update_by` varchar(64) DEFAULT NULL COMMENT '更新人',
  `create_time` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `delete_time` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`dept_id`),
  KEY `idx_sys_dept_delete_time` (`delete_time`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
--  Table structure for `sys_loginlog`
--  登陆日志（不入仓）
-- ----------------------------
DROP TABLE IF EXISTS `sys_loginlog`;
CREATE TABLE `sys_loginlog` (
  `info_id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(128) DEFAULT NULL,
  `status` int DEFAULT NULL,
  `ipaddr` varchar(255) DEFAULT NULL,
  `login_location` varchar(255) DEFAULT NULL,
  `browser` varchar(255) DEFAULT NULL,
  `os` varchar(255) DEFAULT NULL,
  `platform` varchar(255) DEFAULT NULL,
  `login_time` timestamp NULL DEFAULT NULL,
  `create_by` varchar(128) DEFAULT NULL,
  `update_by` varchar(128) DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `msg` varchar(255) DEFAULT NULL,
  `create_time` timestamp NULL DEFAULT NULL,
  `update_time` timestamp NULL DEFAULT NULL,
  `delete_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`info_id`),
  KEY `idx_sys_loginlog_delete_time` (`delete_time`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
--  Table structure for `sys_menu`
--  菜单配置（不入仓）
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu` (
  `menu_id` int NOT NULL AUTO_INCREMENT,
  `menu_name` varchar(128) DEFAULT NULL,
  `title` varchar(64) DEFAULT NULL,
  `icon` varchar(128) DEFAULT NULL,
  `path` varchar(128) DEFAULT NULL,
  `paths` varchar(128) DEFAULT NULL,
  `menu_type` varchar(1) DEFAULT NULL,
  `action` varchar(16) DEFAULT NULL,
  `permission` varchar(32) DEFAULT NULL,
  `parent_id` int DEFAULT NULL,
  `no_cache` char(1) DEFAULT NULL,
  `breadcrumb` varchar(255) DEFAULT NULL,
  `component` varchar(255) DEFAULT NULL,
  `sort` int DEFAULT NULL,
  `visible` char(1) DEFAULT NULL,
  `create_by` varchar(128) DEFAULT NULL,
  `update_by` varchar(128) DEFAULT NULL,
  `is_frame` int DEFAULT '0',
  `create_time` timestamp NULL DEFAULT NULL,
  `update_time` timestamp NULL DEFAULT NULL,
  `delete_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`menu_id`),
  KEY `idx_sys_menu_delete_time` (`delete_time`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
--  Table structure for `sys_post`
--  岗位（不入仓）
-- ----------------------------
DROP TABLE IF EXISTS `sys_post`;
CREATE TABLE `sys_post` (
  `post_id` int NOT NULL AUTO_INCREMENT,
  `post_name` varchar(128) DEFAULT NULL,
  `post_code` varchar(128) DEFAULT NULL,
  `sort` int DEFAULT NULL,
  `status` int DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `create_by` varchar(128) DEFAULT NULL,
  `update_by` varchar(128) DEFAULT NULL,
  `create_time` timestamp NULL DEFAULT NULL,
  `update_time` timestamp NULL DEFAULT NULL,
  `delete_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`post_id`),
  KEY `idx_sys_post_delete_time` (`delete_time`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
--  Table structure for `sys_role`
--  角色配置（不入仓）
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
  `role_id` int NOT NULL AUTO_INCREMENT,
  `role_name` varchar(128) DEFAULT NULL,
  `status` int DEFAULT NULL,
  `role_key` varchar(128) DEFAULT NULL,
  `role_sort` int DEFAULT NULL,
  `flag` varchar(128) DEFAULT NULL,
  `create_by` varchar(128) DEFAULT NULL,
  `update_by` varchar(128) DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `admin` char(1) DEFAULT NULL,
  `create_time` timestamp NULL DEFAULT NULL,
  `update_time` timestamp NULL DEFAULT NULL,
  `delete_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`role_id`),
  KEY `idx_sys_role_delete_time` (`delete_time`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
--  Table structure for `sys_role_dept`
--  角色部门关联表（不入仓）
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_dept`;
CREATE TABLE `sys_role_dept` (
  `role_id` int DEFAULT NULL,
  `dept_id` int DEFAULT NULL,
  `id` int NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
--  Table structure for `sys_role_menu`
--  角色菜单权限关联表（不入仓）
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu` (
  `role_id` int DEFAULT NULL,
  `menu_id` int DEFAULT NULL,
  `role_name` varchar(128) DEFAULT NULL,
  `create_by` varchar(128) DEFAULT NULL,
  `update_by` varchar(128) DEFAULT NULL,
  `id` int NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
--  Table structure for `sys_settings`
--  系统配置（不入仓）
-- ----------------------------
DROP TABLE IF EXISTS `sys_settings`;
CREATE TABLE `sys_settings` (
  `id` int NOT NULL AUTO_INCREMENT,
  `create_time` timestamp NULL DEFAULT NULL,
  `update_time` timestamp NULL DEFAULT NULL,
  `delete_time` timestamp NULL DEFAULT NULL,
  `classify` int DEFAULT NULL,
  `content` json DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_sys_settings_delete_time` (`delete_time`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
--  Table structure for `sys_user`
--  用户表
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
  `user_id` int NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `nick_name` varchar(128) DEFAULT NULL COMMENT '昵称',
  `phone` varchar(11) DEFAULT NULL COMMENT '手机号',
  `role_id` int DEFAULT NULL COMMENT '角色id',
  `salt` varchar(255) DEFAULT NULL COMMENT '加密串',
  `avatar` varchar(255) DEFAULT NULL COMMENT '头像',
  `sex` varchar(255) DEFAULT NULL COMMENT '性别',
  `email` varchar(128) DEFAULT NULL COMMENT '邮箱',
  `dept_id` int DEFAULT NULL COMMENT '部门id',
  `post_id` int DEFAULT NULL COMMENT '岗位id',
  `create_by` varchar(128) DEFAULT NULL COMMENT '创建人',
  `update_by` varchar(128) DEFAULT NULL COMMENT '更新人',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `status` int DEFAULT NULL COMMENT '状态',
  `create_time` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `delete_time` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `username` varchar(64) DEFAULT NULL COMMENT '用户名',
  `password` varchar(128) DEFAULT NULL '密码',
  PRIMARY KEY (`user_id`),
  KEY `idx_sys_user_delete_time` (`delete_time`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
