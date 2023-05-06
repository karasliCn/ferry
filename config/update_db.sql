INSERT INTO `casbin_rule`(`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'admin', '/api/v1/work-order/suspend', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule`(`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'common', '/api/v1/work-order/suspend', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule`(`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'admin', '/api/v1/work-order/circulation-list', 'GET', NULL, NULL, NULL);
COMMIT;

INSERT INTO sys_menu (menu_id, menu_name, title, icon, path, paths, menu_type, action, permission, parent_id, no_cache, breadcrumb, component, sort, visible, create_by, update_by, is_frame, create_time, update_time, delete_time) VALUES (372, '', '挂起工单', 'bug', '/api/v1/work-order/suspend', '/0/63/281/333/372', 'A', 'POST', '', 333, '0', '', '', 0, '1', '1', '', 1, '2021-03-02 22:46:46', '2021-03-02 22:46:46', null);
INSERT INTO sys_menu (menu_id, menu_name, title, icon, path, paths, menu_type, action, permission, parent_id, no_cache, breadcrumb, component, sort, visible, create_by, update_by, is_frame, create_time, update_time, delete_time) VALUES (372, '', '导出工单', 'bug', '/api/v1/work-order/circulation-list', '/0/63/281/333/373', 'A', 'GET', '', 333, '0', '', '', 0, '1', '1', '', 1, '2021-03-02 22:46:46', '2021-03-02 22:46:46', null);
COMMIT;

INSERT INTO `sys_role_menu`(`role_id`, `menu_id`, `role_name`, `create_by`, `update_by`) VALUES (1, 372, 'admin', NULL, NULL);
INSERT INTO `sys_role_menu`(`role_id`, `menu_id`, `role_name`, `create_by`, `update_by`) VALUES (2, 372, 'common', NULL, NULL);
INSERT INTO `sys_role_menu`(`role_id`, `menu_id`, `role_name`, `create_by`, `update_by`) VALUES (1, 373, 'admin', NULL, NULL);
COMMIT;