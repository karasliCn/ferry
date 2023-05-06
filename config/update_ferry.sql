ALTER TABLE  p_work_order_circulation_history add column suspend_time timestamp null default null;
ALTER TABLE  p_work_order_circulation_history add column resume_time timestamp null default null;
